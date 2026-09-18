package golang

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// goVsqlColumnsQueryBody renders <Name>VsqlColumnsQuery: one bool field per
// column (mirroring the picker 1:1), decoded from a raw query string with
// the exact same machinery GoActionQueryParams uses for an action's `qs:`
// fields (emigo.UnmarshalQs + emigo.NewDecoder) - so a caller can build the
// picker straight from an incoming HTTP request's query string
// (?id=true&email=false) instead of constructing <Name>VsqlColumns by hand.
// Its Columns() method folds the parsed values into a real picker, falling
// back to each column's own declared `selected:` default for any key the
// query string didn't actually set - a caller only needs to mention the
// columns they want to override.
//
// Returns "" (and no dependencies) when the vsql declares no columns.
func goVsqlColumnsQueryBody(vsql core.EmiVsql, emiLocation string) (body string, deps []core.CodeChunkDependency) {
	if len(vsql.Columns) == 0 {
		return "", nil
	}

	columnsClass := vsql.GetColumnsClassName()
	queryClass := columnsClass + "Query"

	var fields, assigns strings.Builder
	for _, col := range vsql.Columns {
		if col == nil || col.Name == "" {
			continue
		}
		fieldName := core.ToUpper(col.Name)
		fmt.Fprintf(&fields, "\t%s bool `json:\"%s\"`\n", fieldName, col.Name)
		fmt.Fprintf(&assigns, "\tif q.values.Has(%q) {\n\t\tcols.%s.Selected = q.%s\n\t}\n", col.Name, fieldName, fieldName)
	}

	var out strings.Builder
	fmt.Fprintf(&out, "// %s lets a caller build %s straight from a raw\n", queryClass, columnsClass)
	fmt.Fprintf(&out, "// query string (e.g. \"id=true&email=false\") or an *http.Request, instead\n")
	fmt.Fprintf(&out, "// of constructing %s by hand - see Columns().\n", columnsClass)
	fmt.Fprintf(&out, "type %s struct {\n\tvalues url.Values\n\tmapped map[string]interface{}\n\n%s}\n\n", queryClass, fields.String())

	fmt.Fprintf(&out, "// %sFromString parses rawQuery (a URL query string, no leading \"?\") into a\n", queryClass)
	fmt.Fprintf(&out, "// %s. Keys it doesn't recognize are ignored; see Columns() for how\n", queryClass)
	fmt.Fprintf(&out, "// recognized ones fold into the actual picker.\n")
	fmt.Fprintf(&out, "func %sFromString(rawQuery string) %s {\n", queryClass, queryClass)
	fmt.Fprintf(&out, "\tq := %s{}\n", queryClass)
	fmt.Fprintf(&out, "\tvalues, _ := url.ParseQuery(rawQuery)\n\n")
	fmt.Fprintf(&out, "\tmapped := map[string]interface{}{}\n")
	fmt.Fprintf(&out, "\tif result, err := emigo.UnmarshalQs(rawQuery); err == nil {\n\t\tmapped = result\n\t}\n\n")
	fmt.Fprintf(&out, "\tdecoder, err := emigo.NewDecoder(&emigo.DecoderConfig{\n")
	fmt.Fprintf(&out, "\t\tTagName:          \"json\",\n\t\tWeaklyTypedInput: true,\n\t\tResult:           &q,\n\t})\n")
	fmt.Fprintf(&out, "\tif err == nil {\n\t\t_ = decoder.Decode(mapped)\n\t}\n\n")
	fmt.Fprintf(&out, "\tq.values = values\n\tq.mapped = mapped\n\treturn q\n}\n\n")

	fmt.Fprintf(&out, "// %sFromHttp is %sFromString applied to r.URL.RawQuery.\n", queryClass, queryClass)
	fmt.Fprintf(&out, "func %sFromHttp(r *http.Request) %s {\n\treturn %sFromString(r.URL.RawQuery)\n}\n\n", queryClass, queryClass, queryClass)

	fmt.Fprintf(&out, "// Columns folds the parsed query string into a %s, seeded from\n", columnsClass)
	fmt.Fprintf(&out, "// New%s() so a column the query string never mentioned keeps its\n", columnsClass)
	fmt.Fprintf(&out, "// own declared default rather than becoming false.\n")
	fmt.Fprintf(&out, "func (q %s) Columns() %s {\n\tcols := New%s()\n%s\treturn cols\n}\n",
		queryClass, columnsClass, columnsClass, assigns.String())

	deps = []core.CodeChunkDependency{
		{Location: "net/url"},
		{Location: "net/http"},
		{Location: emiLocation},
	}
	return out.String(), deps
}
