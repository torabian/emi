package golang

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// goVsqlColumnsBody renders the Go source for a vsql's column picker: one
// <Name>VsqlColumns struct (an emigo.ColumnState field per EmiColumn, named
// after its PascalCased field Name), a New<Name>VsqlColumns() constructor
// seeding each field's default Selected state, and a Cols() method that
// joins the currently-selected columns' SQL text (EmiColumn.GetColumn())
// into a projection - so a query template does {{ .Columns.Cols }} instead
// of repeating the column list.
//
// Returns "" (and no dependency) when the vsql declares no columns.
func goVsqlColumnsBody(vsql core.EmiVsql) (body string, usesStrings bool) {
	if len(vsql.Columns) == 0 {
		return "", false
	}

	className := vsql.GetColumnsClassName()

	var fields, ctorFields, colsBody strings.Builder
	for _, col := range vsql.Columns {
		if col == nil || col.Name == "" {
			continue
		}
		fieldName := core.ToUpper(col.Name)

		if col.Description != "" {
			fmt.Fprintf(&fields, "\t// %s\n", col.Description)
		}
		fmt.Fprintf(&fields, "\t%s emigo.ColumnState `json:\"%s\" yaml:\"%s\"`\n", fieldName, col.Name, col.Name)

		fmt.Fprintf(&ctorFields, "\t\t%s: emigo.ColumnState{Selected: %t},\n", fieldName, col.Selected)

		fmt.Fprintf(&colsBody, "\tif c.%s.Selected {\n\t\tparts = append(parts, %q)\n\t}\n", fieldName, col.GetColumn())
	}

	var out strings.Builder
	fmt.Fprintf(&out, "// %s is the column picker for the %s vsql query. Toggle Selected on\n", className, vsql.Name)
	fmt.Fprintf(&out, "// each entry (or start from %s()), then pass it to Prepare%s -\n", "New"+className, vsql.GetClassName())
	fmt.Fprintf(&out, "// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.\n")
	fmt.Fprintf(&out, "type %s struct {\n%s}\n\n", className, fields.String())

	fmt.Fprintf(&out, "// New%s returns the column picker seeded with each column's\n", className)
	fmt.Fprintf(&out, "// default Selected state, as declared on the %s vsql query.\n", vsql.Name)
	fmt.Fprintf(&out, "func New%s() %s {\n\treturn %s{\n%s\t}\n}\n\n", className, className, className, ctorFields.String())

	fmt.Fprintf(&out, "// Cols renders the currently-selected columns as a comma-separated SQL\n")
	fmt.Fprintf(&out, "// projection, in declaration order, so the query template doesn't have to\n")
	fmt.Fprintf(&out, "// repeat the column list.\n")
	fmt.Fprintf(&out, "func (c %s) Cols() string {\n\tparts := []string{}\n%s\treturn strings.Join(parts, \", \")\n}\n", className, colsBody.String())

	return out.String(), true
}

// goVsqlPrepareFunc renders Prepare<Name>Vsql itself. Built as a plain Go
// string (not through the outer text/template pass) so its doc comment can
// mention literal {{ }} template syntax without needing the print-a-string-
// literal trick the surrounding template uses for the same reason.
//
// Four shapes, depending on whether the vsql declares Columns and whether it
// reads its query from an fs.FS (QueryName) instead of a compiled-in
// constant (Query):
//   - neither: today's original shape - (query, params).
//   - columns only: wraps params+columns into <Name>VsqlData.
//   - queryName only: takes fsys fs.FS first, returns an extra error for the
//     file read.
//   - both: takes fsys, wraps into <Name>VsqlData, returns the error too.
func goVsqlPrepareFunc(vsql core.EmiVsql, className, paramsClass, columnsClass, rowClass, dataClass, sqlConstName, queryFileConstName string, hasColumns, useQueryFile bool) string {
	args := fmt.Sprintf("params %s", paramsClass)
	if hasColumns {
		args += fmt.Sprintf(", columns %s", columnsClass)
	}

	returnArgs := "params"
	if hasColumns {
		returnArgs = fmt.Sprintf("%s{ %s: params, Columns: columns }", dataClass, paramsClass)
	}

	var doc strings.Builder
	fmt.Fprintf(&doc, "// Prepare%s returns the query string and params for the %s vsql query,\n", className, vsql.Name)
	fmt.Fprintf(&doc, "// ready to be passed to a SQL driver of your choice.\n")
	if hasColumns {
		fmt.Fprintf(&doc, "// columns controls which optional columns end up in the query - use\n")
		fmt.Fprintf(&doc, "// New%s() to start from the defaults declared in the vsql definition.\n", columnsClass)
		fmt.Fprintf(&doc, "// Scan results into %s.\n", rowClass)
	}

	if !useQueryFile {
		return fmt.Sprintf("%sfunc Prepare%s(%s) (query string, args interface{}) {\n\treturn %s, %s\n}\n",
			doc.String(), className, args, sqlConstName, returnArgs)
	}

	fmt.Fprintf(&doc, "// fsys is read for %s (see %s) instead of a compiled-in\n", vsql.QueryName, queryFileConstName)
	fmt.Fprintf(&doc, "// constant - typically an embed.FS the caller built with go:embed over its\n")
	fmt.Fprintf(&doc, "// own .sql files.\n")
	return fmt.Sprintf(
		"%sfunc Prepare%s(fsys fs.FS, %s) (query string, args interface{}, err error) {\n"+
			"\tb, err := fs.ReadFile(fsys, %s)\n"+
			"\tif err != nil {\n"+
			"\t\treturn \"\", nil, err\n"+
			"\t}\n"+
			"\treturn string(b), %s, nil\n}\n",
		doc.String(), className, args, queryFileConstName, returnArgs)
}

// vsqlColumnFields extracts the embedded EmiField out of each EmiColumn, so
// the declared columns can be run through GoCommonStructGenerator exactly
// like vsql.Params or an EmiDto's fields.
func vsqlColumnFields(columns []*core.EmiColumn) []*core.EmiField {
	fields := make([]*core.EmiField, 0, len(columns))
	for _, col := range columns {
		if col == nil {
			continue
		}
		fields = append(fields, &col.EmiField)
	}
	return fields
}

// GoVsqlCompile generates a single Go file for an EmiVsql definition. The
// file contains:
//
//   - A params DTO struct (produced via GoCommonStructGenerator using the
//     vsql.Params fields, same machinery as regular EmiDto generation).
//   - When vsql.Filters is set (explicitly, or defaulted from Columns during
//     preprocessing): a filters allow-list DTO plus an allowed-field-names
//     slice and membership check (see goVsqlFiltersBody).
//   - When vsql.Columns is set: a row/response DTO (produced the same way,
//     from the columns' embedded EmiField - see vsqlColumnFields), a column
//     picker struct plus a Cols() projection helper (see goVsqlColumnsBody),
//     and a <Name>VsqlData struct combining Params with the picker for
//     template execution.
//   - A const holding the raw SQL query string.
//   - A const holding the vsql query name.
//   - A Prepare<Name>Vsql helper that returns (query string, params interface{}).
//
// The Prepare helper exposes a uniform shape consumers can pass straight to
// db.Prepare/Exec/Query call sites without binding to a specific driver.
func GoVsqlCompile(vsql core.EmiVsql, ctx core.MicroGenContext, complexes []RecognizedComplex, emiLocation string) (*core.CodeChunkCompiled, error) {

	if vsql.Query != "" && vsql.QueryName != "" {
		return nil, fmt.Errorf("vsql %s: query and queryName are mutually exclusive - set exactly one", vsql.Name)
	}
	useQueryFile := vsql.QueryName != ""

	paramsClass := vsql.GetParamsClassName()

	paramsChunk, err := GoCommonStructGenerator(vsql.Params, ctx, GoCommonStructContext{
		RootClassName:       paramsClass,
		RecognizedComplexes: complexes,
		EmiLocation:         emiLocation,
	})
	if err != nil {
		return nil, err
	}

	rowClass := vsql.GetRowClassName()
	var rowBody string
	var rowDeps []core.CodeChunkDependency
	if len(vsql.Columns) > 0 {
		rowChunk, err := GoCommonStructGenerator(vsqlColumnFields(vsql.Columns), ctx, GoCommonStructContext{
			RootClassName:       rowClass,
			RecognizedComplexes: complexes,
			EmiLocation:         emiLocation,
		})
		if err != nil {
			return nil, fmt.Errorf("vsql %s: columns: %w", vsql.Name, err)
		}
		rowBody = string(rowChunk.MainClass.ActualScript)
		rowDeps = rowChunk.MainClass.CodeChunkDependensies
	}

	filtersBody, filtersDeps, err := goVsqlFiltersBody(vsql, ctx, complexes, emiLocation)
	if err != nil {
		return nil, err
	}

	columnsBody, usesStrings := goVsqlColumnsBody(vsql)
	columnsQueryBody, columnsQueryDeps := goVsqlColumnsQueryBody(vsql, emiLocation)
	columnsClass := vsql.GetColumnsClassName()
	dataClass := vsql.GetClassName() + "Data"
	sqlConstName := vsql.GetClassName() + "Query"
	queryFileConstName := vsql.GetClassName() + "QueryFile"

	var querySection string
	if useQueryFile {
		querySection = fmt.Sprintf(
			"// %s is the file name Prepare%s reads this vsql's SQL from, via a\n"+
				"// caller-supplied fs.FS, instead of a compiled-in constant.\n"+
				"const %s = %q\n",
			queryFileConstName, vsql.GetClassName(), queryFileConstName, vsql.QueryName)
	} else {
		querySection = fmt.Sprintf(
			"// %s is the raw SQL string for the %s vsql query.\nconst %s = `%s`\n",
			sqlConstName, vsql.Name, sqlConstName, vsql.Query)
	}

	hasColumns := len(vsql.Columns) > 0
	prepareFunc := goVsqlPrepareFunc(vsql, vsql.GetClassName(), paramsClass, columnsClass, rowClass, dataClass, sqlConstName, queryFileConstName, hasColumns, useQueryFile)

	const tmpl = `
{{ .paramsBody }}

{{ .filtersBody }}

{{ if .hasColumns }}
// {{ .rowClass }} is the response/row DTO for the {{ .name }} vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
{{ .rowBody }}
{{ .columnsBody }}
{{ .columnsQueryBody }}
// {{ .dataClass }} is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ "{{ .Name }}" }} (promoted from {{ .paramsClass }}) alongside {{ "{{ .Columns.Cols }}" }}.
type {{ .dataClass }} struct {
	{{ .paramsClass }}
	Columns {{ .columnsClass }} ` + "`json:\"columns\" yaml:\"columns\"`" + `
}
{{ end }}
// {{ .queryConstName }} is the name of the vsql query, useful for logging or routing.
const {{ .queryConstName }} = "{{ .name }}"

{{ .querySection }}

{{ .prepareFunc }}
`

	t := template.Must(template.New("go_vsql").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"paramsBody":       string(paramsChunk.MainClass.ActualScript),
		"filtersBody":      filtersBody,
		"rowClass":         rowClass,
		"rowBody":          rowBody,
		"columnsBody":      columnsBody,
		"columnsQueryBody": columnsQueryBody,
		"hasColumns":       hasColumns,
		"className":        vsql.GetClassName(),
		"paramsClass":      paramsClass,
		"columnsClass":     columnsClass,
		"dataClass":        dataClass,
		"queryConstName":   vsql.GetClassName() + "Name",
		"querySection":     querySection,
		"prepareFunc":      prepareFunc,
		"name":             vsql.Name,
	}); err != nil {
		return nil, err
	}

	deps := paramsChunk.MainClass.CodeChunkDependensies
	deps = append(deps, filtersDeps...)
	if hasColumns {
		deps = append(deps, rowDeps...)
		deps = append(deps, core.CodeChunkDependency{Location: emiLocation})
		if usesStrings {
			deps = append(deps, core.CodeChunkDependency{Location: "strings"})
		}
		deps = append(deps, columnsQueryDeps...)
	}
	if useQueryFile {
		deps = append(deps, core.CodeChunkDependency{Location: "io/fs"})
	}

	res := &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: deps,
		Tokens:                paramsChunk.MainClass.Tokens,
		SuggestedFileName:     vsql.GetClassName() + ".go",
		SuggestedExtension:    paramsChunk.MainClass.SuggestedExtension,
	}

	return res, nil
}

// GoVsqlsGenerate iterates module.Vsqls and produces a VirtualFile per query.
func GoVsqlsGenerate(module *core.Emi, ctx core.MicroGenContext, complexes []RecognizedComplex, emiLocation, packageName string) ([]core.VirtualFile, error) {
	files := []core.VirtualFile{}
	for _, v := range module.Vsqls {
		if v.Name == "" {
			continue
		}
		chunk, err := GoVsqlCompile(v, ctx, complexes, emiLocation)
		if err != nil {
			return nil, fmt.Errorf("vsql %s: %w", v.Name, err)
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk, packageName),
		})
	}
	return files, nil
}
