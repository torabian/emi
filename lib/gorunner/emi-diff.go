package gorunner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v3"
)

// DiffEmiYamlCommand is a CLI-only helper (no compiler involved, not exposed
// to wasm - see cmd/emi-wasm/main.go, which never imports gorunner) that
// finds every *.emi.yml file changed between two git branches and prints a
// clear file list plus a line-by-line unified diff for each file. It's meant
// to be piped straight into an AI prompt so it can figure out what migration
// is needed for the module/entity changes.
var DiffEmiYamlCommand = cli.Command{
	Name:        "diff:emi",
	Description: "Lists every *.emi.yml file changed between --branch and --with, with a clear per-file line diff. Intended as input for an AI to draft migrations from.",
	Usage:       "Show *.emi.yml changes between two git branches (default: current branch vs main)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "branch",
			Usage: "Branch containing the changes. Defaults to the current checked out branch.",
		},
		&cli.StringFlag{
			Name:  "with",
			Usage: "Base branch to compare against.",
			Value: "main",
		},
		&cli.StringFlag{
			Name:  "include",
			Usage: "Restrict results to a directory (e.g. modules/billing) or a glob pattern (e.g. modules/**/inventory*.emi.yml). Matched against each changed file's repo-relative path. Defaults to every *.emi.yml file in the repo.",
		},
		&cli.StringFlag{
			Name:  "output",
			Usage: "Optional file path to write the report to. Defaults to stdout.",
		},
	},
	Action: func(_ context.Context, c *cli.Command) error {
		branch := c.String("branch")
		with := c.String("with")

		if branch == "" {
			b, err := currentGitBranch()
			if err != nil {
				return fmt.Errorf("could not determine current branch, pass --branch explicitly: %w", err)
			}
			branch = b
		}

		report, err := buildEmiYamlDiffReport(branch, with, c.String("include"))
		if err != nil {
			return err
		}

		if c.String("output") == "" {
			fmt.Println(report)
			return nil
		}

		if err := os.WriteFile(c.String("output"), []byte(report), 0644); err != nil {
			return fmt.Errorf("error writing report to disk: %v: %w", c.String("output"), err)
		}

		return nil
	},
}

// changedEmiFile is one row of `git diff --name-status`, with OldPath set
// only for renames/copies (where git reports two paths on the line).
type changedEmiFile struct {
	Status  string
	Path    string
	OldPath string
}

func currentGitBranch() (string, error) {
	out, err := runGit("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}

	branch := strings.TrimSpace(out)
	if branch == "" || branch == "HEAD" {
		return "", fmt.Errorf("not currently on a named branch (detached HEAD)")
	}

	return branch, nil
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}

	return stdout.String(), nil
}

// buildEmiYamlDiffReport compares `branch` against `with` using git's
// triple-dot (merge-base) range, so it reports only what `branch` actually
// changed since it diverged from `with` - not unrelated commits `with`
// picked up in the meantime. When `include` is non-empty, only files whose
// repo-relative path matches it (as a directory prefix or a glob pattern,
// see matchesInclude) are kept.
func buildEmiYamlDiffReport(branch, with, include string) (string, error) {
	rangeSpec := with + "..." + branch

	nameStatusOut, err := runGit("diff", "--name-status", rangeSpec, "--", "*.emi.yml")
	if err != nil {
		return "", fmt.Errorf("failed diffing %q against %q: %w", branch, with, err)
	}

	files := parseNameStatus(nameStatusOut)
	if include != "" {
		files = filterChangedFiles(files, include)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Comparing branch %q against %q (range: %s)\n", branch, with, rangeSpec)
	if include != "" {
		fmt.Fprintf(&b, "Filtered to: %s\n", include)
	}
	b.WriteString(strings.Repeat("=", 70) + "\n\n")

	if len(files) == 0 {
		b.WriteString("No *.emi.yml changes found.\n")
		return b.String(), nil
	}

	fmt.Fprintf(&b, "Changed files (%d):\n", len(files))
	for _, f := range files {
		if f.OldPath != "" {
			fmt.Fprintf(&b, "  %-6s %s -> %s\n", describeGitStatus(f.Status), f.OldPath, f.Path)
		} else {
			fmt.Fprintf(&b, "  %-6s %s\n", describeGitStatus(f.Status), f.Path)
		}
	}
	b.WriteString("\n")

	for _, f := range files {
		b.WriteString(strings.Repeat("-", 70) + "\n")
		if f.OldPath != "" {
			fmt.Fprintf(&b, "File: %s -> %s (%s)\n", f.OldPath, f.Path, describeGitStatus(f.Status))
		} else {
			fmt.Fprintf(&b, "File: %s (%s)\n", f.Path, describeGitStatus(f.Status))
		}
		b.WriteString(strings.Repeat("-", 70) + "\n")

		paths := []string{f.Path}
		if f.OldPath != "" {
			paths = []string{f.OldPath, f.Path}
		}

		diffArgs := append([]string{"diff", "--unified=3", rangeSpec, "--"}, paths...)
		patch, err := runGit(diffArgs...)
		if err != nil {
			return "", fmt.Errorf("failed diffing file %s: %w", f.Path, err)
		}

		if strings.TrimSpace(patch) == "" {
			b.WriteString("(no textual diff - binary or mode-only change)\n\n")
			continue
		}

		b.WriteString(patch)
		if !strings.HasSuffix(patch, "\n") {
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	return b.String(), nil
}

// parseNameStatus turns `git diff --name-status` output into changedEmiFile
// rows. Rename/copy lines carry a similarity-prefixed status (e.g. "R100")
// followed by both the old and new path, tab separated.
func parseNameStatus(out string) []changedEmiFile {
	var files []changedEmiFile

	for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
		if line == "" {
			continue
		}

		fields := strings.Split(line, "\t")
		if len(fields) < 2 {
			continue
		}

		status := fields[0]

		if (strings.HasPrefix(status, "R") || strings.HasPrefix(status, "C")) && len(fields) >= 3 {
			files = append(files, changedEmiFile{Status: status, OldPath: fields[1], Path: fields[2]})
			continue
		}

		files = append(files, changedEmiFile{Status: status, Path: fields[1]})
	}

	return files
}

// filterChangedFiles keeps only the changedEmiFile entries whose new path
// (or, for a rename/copy, either path) matches include.
func filterChangedFiles(files []changedEmiFile, include string) []changedEmiFile {
	var out []changedEmiFile

	for _, f := range files {
		if matchesInclude(include, f.Path) || (f.OldPath != "" && matchesInclude(include, f.OldPath)) {
			out = append(out, f)
		}
	}

	return out
}

// matchesInclude decides whether a repo-relative file path falls under the
// --include filter. A pattern with no glob metacharacters (*, ?, [) is
// treated as a directory (or exact file) prefix, e.g. "modules/billing"
// matches everything under that directory. A pattern containing any of
// those characters is compiled as a glob against the full path, where "**"
// matches across directory separators and "*"/"?" don't.
func matchesInclude(include, filePath string) bool {
	if include == "" {
		return true
	}

	include = strings.TrimPrefix(filepath.ToSlash(include), "./")
	filePath = filepath.ToSlash(filePath)

	if !strings.ContainsAny(include, "*?[") {
		trimmed := strings.TrimSuffix(include, "/")
		return filePath == trimmed || strings.HasPrefix(filePath, trimmed+"/")
	}

	return globToRegexp(include).MatchString(filePath)
}

// globToRegexp compiles a gitignore/glob-style pattern into an anchored
// regexp: "**/" matches zero or more path segments, "**" matches anything
// (including "/"), "*" matches within a single path segment, "?" matches one
// non-separator character, and everything else is matched literally.
func globToRegexp(pattern string) *regexp.Regexp {
	var b strings.Builder
	b.WriteString("^")

	for i := 0; i < len(pattern); {
		switch {
		case strings.HasPrefix(pattern[i:], "**/"):
			b.WriteString("(?:.*/)?")
			i += 3
		case strings.HasPrefix(pattern[i:], "**"):
			b.WriteString(".*")
			i += 2
		case pattern[i] == '*':
			b.WriteString("[^/]*")
			i++
		case pattern[i] == '?':
			b.WriteString("[^/]")
			i++
		default:
			b.WriteString(regexp.QuoteMeta(string(pattern[i])))
			i++
		}
	}

	b.WriteString("$")

	return regexp.MustCompile(b.String())
}

func describeGitStatus(status string) string {
	switch {
	case status == "A":
		return "added"
	case status == "D":
		return "deleted"
	case strings.HasPrefix(status, "R"):
		return "renamed"
	case strings.HasPrefix(status, "C"):
		return "copied"
	case status == "M":
		return "modified"
	default:
		return status
	}
}
