package golang

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// goVsqlFiltersBody renders a vsql's Filters as Go source: the
// {Name}VsqlFilters struct itself (via GoCommonStructGenerator - the exact
// same struct generator Params/an EmiDto's own fields use, per Filters'
// own doc comment on EmiVsql), plus an allowed-field-names slice and a
// membership check - the actual runtime "is this JSON-Logic filter only
// touching fields it's allowed to" validation Emi can offer in Go, given Go
// generics can't express a string-literal union the way TypeScript can (see
// the JS compiler's equivalent).
//
// Returns ("", nil, nil) when the vsql declares no Filters.
func goVsqlFiltersBody(vsql core.EmiVsql, ctx core.MicroGenContext, complexes []RecognizedComplex, emiLocation string) (body string, deps []core.CodeChunkDependency, err error) {
	if len(vsql.Filters) == 0 {
		return "", nil, nil
	}

	filtersClass := vsql.GetFiltersClassName()

	filtersChunk, err := GoCommonStructGenerator(vsql.Filters, ctx, GoCommonStructContext{
		RootClassName:       filtersClass,
		RecognizedComplexes: complexes,
		EmiLocation:         emiLocation,
	})
	if err != nil {
		return "", nil, fmt.Errorf("vsql %s: filters: %w", vsql.Name, err)
	}

	fieldNamesVar := vsql.GetClassName() + "FilterFields"
	allowedFunc := vsql.GetClassName() + "FilterFieldAllowed"

	var names strings.Builder
	for _, f := range vsql.Filters {
		if f == nil || f.Name == "" {
			continue
		}
		fmt.Fprintf(&names, "\t%q,\n", f.Name)
	}

	var out strings.Builder
	fmt.Fprintf(&out, "// %s is the allow-list DTO for the %s vsql query's filter -\n", filtersClass, vsql.Name)
	fmt.Fprintf(&out, "// see EmiVsql.Filters. Emi never executes or transpiles a filter itself\n")
	fmt.Fprintf(&out, "// (e.g. via jsonlogic2sql); this only exists to describe, and let you\n")
	fmt.Fprintf(&out, "// validate against, the fields a caller-supplied JSON-Logic expression is\n")
	fmt.Fprintf(&out, "// allowed to reference.\n")
	out.WriteString(string(filtersChunk.MainClass.ActualScript))
	out.WriteString("\n\n")

	fmt.Fprintf(&out, "// %s lists the top-level field names a JSON-Logic filter against the\n", fieldNamesVar)
	fmt.Fprintf(&out, "// %s vsql query may legally reference (each {\"var\": ...} leaf). Does not\n", vsql.Name)
	fmt.Fprintf(&out, "// enumerate a nested object field's own sub-fields - only %s's own\n", filtersClass)
	fmt.Fprintf(&out, "// top-level entries.\n")
	fmt.Fprintf(&out, "var %s = []string{\n%s}\n\n", fieldNamesVar, names.String())

	fmt.Fprintf(&out, "// %s reports whether name is one of %s's allowed\n", allowedFunc, fieldNamesVar)
	fmt.Fprintf(&out, "// top-level filter fields - a cheap check worth running against every\n")
	fmt.Fprintf(&out, "// {\"var\": ...} leaf found while walking a caller-supplied filter, before\n")
	fmt.Fprintf(&out, "// it ever reaches SQL.\n")
	fmt.Fprintf(&out, "func %s(name string) bool {\n\tfor _, f := range %s {\n\t\tif f == name {\n\t\t\treturn true\n\t\t}\n\t}\n\treturn false\n}\n",
		allowedFunc, fieldNamesVar)

	return out.String(), filtersChunk.MainClass.CodeChunkDependensies, nil
}
