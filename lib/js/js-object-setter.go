package js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

var setterTemplateNew = template.Must(template.New("setter").Parse(`
{{.Jsdoc}}
/// XXX
set{{.UpperName}}(value{{if .FieldType}}: {{.FieldType}}{{end}}) {
	// Only accept array types
	{{if eq .Type "array"}}
	if (!Array.isArray(value) && value !== null && value !== undefined) {
		return this
	}

	// If the value is array, we need to check first item, if is instance of the class
	if (value.length > 0 && value[0] instanceof {{.ArrayClass}}) {
		this["{{.Name}}"] = value
	} else {
		this["{{.Name}}"] = value.map(item => new {{.ArrayClass}}(item))
	}
	{{else}}
	this["{{.Name}}"] = value
	{{end}}
	return this
}
`))

// For each field, we can generate a setter function
func JsTsFieldSetterGenerator(jsFieldType string, field *core.EmiField, isTypeScript bool) string {
	setterjsdoc := NewJsDoc("  ")
	setterjsdoc.Add(field.Description)
	setterjsdoc.Add(fmt.Sprintf("@param {%v}", jsFieldType))
	// isFieldNullable := strings.Contains(string(field.Type), "?")

	// if isTypeScript {
	// 	nullableMark := ""
	// 	if isFieldNullable {
	// 		nullableMark = " | null"
	// 	}
	// 	setterFunc = setterjsdoc.String() +
	// 		fmt.Sprintf(
	// 			"set%v (value: %v %v) { this[`%v`] = value; return this; }",
	// 			core.ToUpper(field.Name),
	// 			tsFieldType,
	// 			nullableMark,
	// 			field.Name,
	// 		)
	// }

	data := map[string]any{
		"Jsdoc":     setterjsdoc.String(),
		"UpperName": core.ToUpper(field.Name),
		"Name":      field.Name,
		"FieldType": jsFieldType,
		"Type":      field.Type,
	}

	var buf bytes.Buffer
	if err := setterTemplateNew.Execute(&buf, data); err != nil {
		panic(err)
	}

	return setterjsdoc.String() + buf.String()
}
