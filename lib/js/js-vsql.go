// Compiles a vsql's request/response shapes for the front end: the params
// (and In, already merged into Params by preprocessing) it accepts, and,
// when Columns is declared, the row shape a caller can expect back. There is
// no database connection in a browser, so this deliberately stops there -
// the picker/Prepare/actual-SQL machinery lib/golang generates stays
// Go-only; a front end only ever needs the two shapes to populate a request
// and read a response.
package js

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// jsVsqlColumnFields extracts the embedded EmiField out of each EmiColumn -
// the JS-side twin of lib/golang's vsqlColumnFields (unexported there, so
// duplicated rather than reached across packages) - so a vsql's declared
// columns can be run through JsCommonObjectGenerator exactly like params or
// an EmiDto's fields.
func jsVsqlColumnFields(columns []*core.EmiColumn) []*core.EmiField {
	fields := make([]*core.EmiField, 0, len(columns))
	for _, col := range columns {
		if col == nil {
			continue
		}
		fields = append(fields, &col.EmiField)
	}
	return fields
}

// jsVsqlFilterTypeAppendix builds the TypeScript-only appendix for a vsql's
// Filters chunk: a literal string-union of allowed top-level field names,
// plus one minimal JSON-Logic (https://jsonlogic.com) node type constrained
// to that union. Only a `{ var: ... }` leaf is typed against real fields -
// JsonLogic operators (and/or/==/...) are too polymorphic to model further
// generically, so every other key keeps whatever shape a runtime
// (jsonlogic2sql, or any other) already expects. This is the JS/TS half of
// the same allow-list Go gets as a plain runtime []string + membership check
// (see lib/golang/go-vsql-filters.go) - TypeScript's string-literal unions
// buy a compile-time check Go's generics can't.
//
// The generic index signature `{ [operator: string]: ... }` alone can't
// single out the "var" key from any other operator key - TypeScript's
// structural typing would let `{ var: "notAField" }` slip in as just another
// operator node, silently defeating the whole constraint (verified against a
// real tsc run before landing this). `& { var?: never }` on the operator-node
// arm is what actually forces an object carrying a "var" key to be checked
// against the FieldValue arm - the field-name union - instead.
//
// Requires the consuming tsconfig to have strictNullChecks on (part of
// --strict, or set alone) - also verified against a real tsc run: without
// it, a value literal's `var` property gets inferred as the widened `string`
// instead of the specific field-name literal, and the constraint silently
// stops rejecting anything. examples/vsql-columns/ts has a working
// tsconfig.json plus a filter-typesafety.ts demonstrating both a rejected
// unknown field and a real jsonlogic2sql-shaped filter typechecking.
//
// Returns "" for an empty fields list (nothing to type).
func jsVsqlFilterTypeAppendix(vsql core.EmiVsql, fields []*core.EmiField) string {
	names := make([]string, 0, len(fields))
	for _, f := range fields {
		if f == nil || f.Name == "" {
			continue
		}
		names = append(names, strconv.Quote(f.Name))
	}
	if len(names) == 0 {
		return ""
	}

	fieldType := core.ToUpper(vsql.Name) + "VsqlFilterField"
	filterType := core.ToUpper(vsql.Name) + "VsqlFilter"
	valueType := filterType + "Value"
	varNodeType := filterType + "VarNode"
	opNodeType := filterType + "OperatorNode"

	var out strings.Builder
	fmt.Fprintf(&out, "\n/** Field names a JSON-Logic filter against the %s vsql query may legally reference. */\n", vsql.Name)
	fmt.Fprintf(&out, "export type %s = %s;\n\n", fieldType, strings.Join(names, " | "))
	fmt.Fprintf(&out, "/**\n * Minimal typed JSON-Logic node (https://jsonlogic.com): only `{ var: ... }` is\n")
	fmt.Fprintf(&out, " * constrained to a real field - every other operator stays loosely typed.\n */\n")
	fmt.Fprintf(&out, "export type %s = { var: %s };\n", varNodeType, fieldType)
	fmt.Fprintf(&out, "export type %s = { [operator: string]: %s } & { var?: never };\n", opNodeType, valueType)
	fmt.Fprintf(&out, "export type %s = %s | string | number | boolean | null | %s[];\n", valueType, filterType, valueType)
	fmt.Fprintf(&out, "export type %s = %s | %s;\n", filterType, varNodeType, opNodeType)

	return out.String()
}

// JsVsqlGenerate compiles one vsql into up to three JS/TS classes:
// {Name}VsqlParams (always), {Name}VsqlFilters (when Filters is set - either
// explicitly, or defaulted from Columns during preprocessing), and
// {Name}VsqlRow (when Columns is declared). All three go through
// JsCommonObjectGenerator - the exact same struct generator dtos and actions
// use - so nested objects, arrays, enums and complex fields all work
// identically to anywhere else in the module. Under --tags typescript, the
// Filters chunk also gets the field-name union + minimal JSON-Logic node
// type (see jsVsqlFilterTypeAppendix).
func JsVsqlGenerate(vsql core.EmiVsql, ctx core.MicroGenContext, complexes []RecognizedComplex) ([]*core.CodeChunkCompiled, error) {
	var chunks []*core.CodeChunkCompiled

	paramsChunk, err := JsCommonObjectGenerator(vsql.Params, ctx, JsCommonObjectContext{
		RootClassName:       vsql.GetParamsClassName(),
		RecognizedComplexes: complexes,
		Description:         vsql.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("vsql %s: params: %w", vsql.Name, err)
	}
	chunks = append(chunks, paramsChunk)

	if len(vsql.Filters) > 0 {
		filtersChunk, err := JsCommonObjectGenerator(vsql.Filters, ctx, JsCommonObjectContext{
			RootClassName:       vsql.GetFiltersClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, fmt.Errorf("vsql %s: filters: %w", vsql.Name, err)
		}
		if ctx.HasTag(Typescript) {
			filtersChunk.ActualScript = append(filtersChunk.ActualScript, []byte(jsVsqlFilterTypeAppendix(vsql, vsql.Filters))...)
		}
		chunks = append(chunks, filtersChunk)
	}

	if len(vsql.Columns) > 0 {
		rowChunk, err := JsCommonObjectGenerator(jsVsqlColumnFields(vsql.Columns), ctx, JsCommonObjectContext{
			RootClassName:       vsql.GetRowClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, fmt.Errorf("vsql %s: columns: %w", vsql.Name, err)
		}
		chunks = append(chunks, rowChunk)
	}

	return chunks, nil
}
