// Renders common JS/TS objects such as entities and DTOs.

package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func (x jsFieldVariable) CreateSetterFunction(ctx core.MicroGenContext) string {

	tsValue := "value: " + x.ComputedType

	if x.IsNullable {
		tsValue += " | null | undefined"
	}
	claims := []core.JsFnArgument{
		{
			Key: "arg.value",
			Ts:  tsValue,
			Js:  "value",
		},
	}
	claimsRendered := core.ClaimRender(claims, ctx)

	var setterTemplate = template.Must(template.New("setter").Parse(`
{{.ctx.JsDoc}}
set {{ .ctx.Name }} (|@arg.value|) {
	{{ if or (eq .ctx.Type "string") }}
		this.#{{.ctx.Name}} = String(value);
	{{ end }}

	{{ if or (eq .ctx.Type "string?") }}
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#{{.ctx.Name}} = correctType ? value : String(value);
	{{ end }}

	{{ if and (eq .ctx.IsNumeric true) (eq .ctx.IsNullable true) }}
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)

		if (!Number.isNaN(parsedValue)) {
			this.#{{.ctx.Name}} = parsedValue;
		}
	{{ end }}
	
	{{ if and (eq .ctx.IsNumeric true) (eq .ctx.IsNullable false) }}
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)

		if (!Number.isNaN(parsedValue)) {
			this.#{{.ctx.Name}} = parsedValue;
		}
	{{ end }}
	
	{{ if and (eq .ctx.Type "bool")}}
		this.#{{.ctx.Name}} = Boolean(value);
	{{ end }}

 	{{ if .ctx.ComplexClass }}
	 	if (value instanceof {{ .ctx.ComplexClass }}) {
			this.#{{.ctx.Name}} = value
		} else {
		 	this.#{{.ctx.Name}} = new {{ .ctx.ComplexClass }}(value)
		}
	{{ end }}
	
	{{ if and (eq .ctx.Type "bool?")}}
	 	const correctType = value === true || value === false || value === undefined || value === null
		this.#{{.ctx.Name}} = correctType ? value : Boolean(value);
	{{ end }}
	
	{{ if and (eq .ctx.Type "any")}}
		this.#{{.ctx.Name}} = value;
	{{ end }}
	
	{{ if or (eq .ctx.Type "array") (eq .ctx.Type "array?") (eq .ctx.Type "collection?") (eq .ctx.Type "collection")}}
	 	// For arrays, you only can pass arrays to the object
	 	if (!Array.isArray(value)) {
			return;
		}
		if (value.length > 0 && value[0] instanceof {{.ctx.ConstructorClass}}) {
			this.#{{.ctx.Name}} = value
		} else {
			this.#{{.ctx.Name}} = value.map(item => new {{.ctx.ConstructorClass}}(item))
		}
 	{{ end }}

	{{ if or (eq .ctx.Type "object") (eq .ctx.Type "object?") (eq .ctx.Type "one") (eq .ctx.Type "one?")}}
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof {{.ctx.ConstructorClass}}) {
			this.#{{.ctx.Name}} = value
		} else {
			this.#{{.ctx.Name}} = new {{.ctx.ConstructorClass}}(value)
		}
 	{{ end }}
}

set{{ .ctx.Upper }} (|@arg.value|) {
	this.{{.ctx.Name}} = value

	return this
}
`))
	setterjsdoc := NewJsDoc("  ")
	setterjsdoc.Add(fmt.Sprintf("@param {%v}", x.ComputedType))

	data := map[string]any{
		"ctx": x,
	}

	var buf bytes.Buffer
	if err := setterTemplate.Execute(&buf, data); err != nil {
		panic(err)
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	return templateResult
}
