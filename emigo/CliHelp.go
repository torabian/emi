package emigo

import (
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli/v3"
)

// Importing emigo (every generated cli does) reshapes the urfave/cli help
// output of the whole application:
//
//   - names of flags and commands are bold, the value type of a flag is blue
//     (only on an interactive terminal; NO_COLOR turns it off)
//   - a flag or command with a long description gets that description on its
//     own line below the name, followed by an empty line, instead of a crowded
//     single line.

// cliHelpLongUsage is the description length above which the description moves
// under the name.
const cliHelpLongUsage = 60

func cliHelpUsesColor() bool { return cliUsesColor(os.Stdout) }

// cliUsesColor is true only when out is an interactive terminal and NO_COLOR is unset.
func cliUsesColor(out *os.File) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	st, err := out.Stat()
	return err == nil && st.Mode()&os.ModeCharDevice != 0
}

var cliYamlKey = regexp.MustCompile(`(?m)^(\s*(?:- )*)([^\s-][^:\n]*?):(\s|$)`)

// cliBoldYamlKeys makes the keys of a YAML document bold when out is a terminal.
func cliBoldYamlKeys(doc string, out *os.File) string {
	if !cliUsesColor(out) {
		return doc
	}
	return cliYamlKey.ReplaceAllString(doc, "${1}"+cliHelpBold+"${2}"+cliHelpReset+":${3}")
}

// ansi 256-color 33 stays readable on dark and light backgrounds.
const (
	cliHelpBold  = "\x1b[1m"
	cliHelpBlue  = "\x1b[38;5;33m"
	cliHelpReset = "\x1b[0m"
)

func cliHelpColorFlagName(name string) string {
	words := strings.Split(name, " ")
	for i, w := range words {
		switch {
		case strings.HasPrefix(w, "-"):
			words[i] = cliHelpBold + w + cliHelpReset
		case w == "[" || w == "]" || w == "":
		default:
			words[i] = cliHelpBlue + w + cliHelpReset
		}
	}
	return strings.Join(words, " ")
}

const (
	cliHelpCommandCategoryOriginal = `{{template "visibleCommandCategoryTemplate" .}}`
	cliHelpCommandFlatOriginal     = `{{template "visibleCommandTemplate" .}}`
)

// cliHelpCommandList lists the visible commands of the current dot.
func cliHelpCommandList(bold, reset string) string {
	return `{{ $cv := offsetCommands .VisibleCommands 5}}{{range .VisibleCommands}}
   {{$s := join .Names ", "}}{{if gt (len .Usage) 60}}` + bold + `{{$s}}` + reset + `
      {{wrap .Usage 6}}
{{else}}` + bold + `{{$s}}` + reset + `{{ $sp := subtract $cv (offset $s 3) }}{{ indent $sp ""}}{{wrap .Usage $cv}}{{end}}{{end}}`
}

// The default listings are private templates of urfave, so the exported help
// templates are rewritten to carry these ones. Named categories keep the
// original layout.
func cliHelpCommandCategoryTemplate(bold, reset string) string {
	return `{{range .VisibleCategories}}{{if .Name}}

   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}` + cliHelpCommandList(bold, reset) + `{{end}}{{end}}`
}

func init() {
	color := cliHelpUsesColor()

	defaultStringer := cli.FlagStringer
	cli.FlagStringer = func(f cli.Flag) string {
		out := defaultStringer(f)
		name, usage, ok := strings.Cut(out, "\t")
		if !ok {
			return out
		}
		if color {
			name = cliHelpColorFlagName(name)
		}
		if len(usage) <= cliHelpLongUsage {
			return name + "\t" + usage
		}
		// No tab: the help tabwriter leaves the line alone. urfave's template
		// indents continuation lines by 6 spaces; the trailing newline leaves an
		// empty line before the next flag.
		return name + "\n" + usage + "\n"
	}

	bold, reset := "", ""
	if color {
		bold, reset = cliHelpBold, cliHelpReset
	}
	category := cliHelpCommandCategoryTemplate(bold, reset)
	flat := cliHelpCommandList(bold, reset)
	for _, t := range []*string{&cli.RootCommandHelpTemplate, &cli.CommandHelpTemplate, &cli.SubcommandHelpTemplate} {
		*t = strings.ReplaceAll(*t, cliHelpCommandCategoryOriginal, category)
		*t = strings.ReplaceAll(*t, cliHelpCommandFlatOriginal, flat)
	}
}
