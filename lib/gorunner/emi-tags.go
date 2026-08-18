package gorunner

import (
	"context"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/torabian/emi/lib/c"
	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/cpp"
	"github.com/torabian/emi/lib/csharp"
	"github.com/torabian/emi/lib/dart"
	"github.com/torabian/emi/lib/golang"
	"github.com/torabian/emi/lib/java"
	"github.com/torabian/emi/lib/js"
	"github.com/torabian/emi/lib/kotlin"
	"github.com/torabian/emi/lib/php"
	"github.com/torabian/emi/lib/python"
	"github.com/urfave/cli/v3"
)

// languageCompilerTags is the registry `emi tags` walks to print every
// target's supported --tags values and what they do. Each entry's Tags
// comes from that language package's own CompilerTags var (see
// lib/<lang>/<lang>-compiler-tags.go) - add a row here whenever a language
// package gains one. Swift is omitted: it doesn't read any compiler tags today.
var languageCompilerTags = []struct {
	Name string
	Tags []core.CompilerTagDoc
}{
	{"golang", golang.CompilerTags},
	{"javascript/typescript", js.CompilerTags},
	{"python", python.CompilerTags},
	{"dart", dart.CompilerTags},
	{"java", java.CompilerTags},
	{"csharp", csharp.CompilerTags},
	{"php", php.CompilerTags},
	{"c", c.CompilerTags},
	{"cpp", cpp.CompilerTags},
	{"kotlin", kotlin.CompilerTags},
}

// renderCompilerTags formats every registered language's compiler tags into
// a grouped, aligned text listing.
func renderCompilerTags() string {
	var b strings.Builder

	b.WriteString("Compiler tags\n")
	b.WriteString("=============\n\n")
	b.WriteString("Tags are per-target-language feature switches for the code generator,\n")
	b.WriteString("passed via --tags as a comma separated list, e.g.:\n\n")
	b.WriteString("  emi golang:action --tags split-cli,split-gin --path ... --output ...\n\n")

	for _, lang := range languageCompilerTags {
		if len(lang.Tags) == 0 {
			continue
		}

		fmt.Fprintf(&b, "%s\n", lang.Name)

		w := tabwriter.NewWriter(&b, 0, 0, 3, ' ', 0)
		for _, t := range lang.Tags {
			fmt.Fprintf(w, "  %s\t%s\n", t.Tag, t.Description)
		}
		w.Flush()

		b.WriteString("\n")
	}

	return b.String()
}

var TagsCommand = cli.Command{
	Name:        "tags",
	Description: "List every --tags value each target language's compiler supports, and what it does",
	Usage:       "List every --tags value each target language's compiler supports, and what it does",
	Action: func(_ context.Context, _ *cli.Command) error {
		fmt.Print(renderCompilerTags())
		return nil
	},
}
