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

	if core.FieldType(x.Type) == core.FieldTypeOne || core.FieldType(x.Type) == core.FieldTypeOneNullable {
		tsValue += " |  InstanceType<typeof " + x.ConstructorClass + ">"
	}

	if core.FieldType(x.Type) == core.FieldTypeCollection || core.FieldType(x.Type) == core.FieldTypeCollectionNullable {
		tsValue += " |  InstanceType<typeof " + x.ConstructorClass + ">[]"
	}

	if core.FieldType(x.Type) == core.FieldTypeArray || core.FieldType(x.Type) == core.FieldTypeArrayNullable {
		tsValue += " |  InstanceType<typeof " + x.ConstructorClass + ">[]"
	}

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

	{{ else if or (eq .ctx.Type "string?") }}
	 	const correctType = typeof value === 'string' || value === undefined || value === null
		this.#{{.ctx.Name}} = correctType ? value : String(value);

	{{ else if and (eq .ctx.IsNumeric true) (eq .ctx.IsNullable true) }}
	 	const correctType = typeof value === 'number' || value === undefined || value === null
		const parsedValue = correctType ? value : Number(value)

		if (!Number.isNaN(parsedValue)) {
			this.#{{.ctx.Name}} = parsedValue;
		}
	
	{{ else if and (eq .ctx.IsNumeric true) (eq .ctx.IsNullable false) }}
	 	const correctType = typeof value === 'number'
		const parsedValue = correctType ? value : Number(value)

		if (!Number.isNaN(parsedValue)) {
			this.#{{.ctx.Name}} = parsedValue;
		}

	{{ else if and (eq .ctx.Type "bool")}}
		this.#{{.ctx.Name}} = Boolean(value);

 	{{ else if .ctx.ComplexClass }}
	 	if (value instanceof {{ .ctx.ComplexClass }}) {
			this.#{{.ctx.Name}} = value
		} else {
		 	this.#{{.ctx.Name}} = new {{ .ctx.ComplexClass }}(value)
		}
	
	{{ else if and (eq .ctx.Type "bool?")}}
	 	const correctType = value === true || value === false || value === undefined || value === null
		this.#{{.ctx.Name}} = correctType ? value : Boolean(value);
	

	{{ else if or (eq .ctx.Type "array") (eq .ctx.Type "array?") }}

		{{ if or (eq .ctx.Type "array?") }}
		// For nullable array, we allow explicit undefined or null values
		if (value === null || value === undefined) {
			this.#{{.ctx.Name}} = value;

			return
		}
		{{ end }}

		// When the passed value is already an array, we check if we need to
		// cast the inner items into class instance.
		if (Array.isArray(value)) {
			if (value.length > 0 && value[0] instanceof {{.ctx.ConstructorClass}}) {
				this.#{{.ctx.Name}} = MArray.of(value);
			} else {
				this.#{{.ctx.Name}} = MArray.of(
					value.map((item) => new {{.ctx.ConstructorClass}}(item)),
				);
			}

			return;
		}

		// If the instance is already an MArray, we assume it's all good.
		if (value instanceof MArray) {
			this.#{{.ctx.Name}} = value;

			return;
		}

		// If the value is not array, and is not a MArray, we need to be consider,
		// it might be eligible to be casted into MArray.
		{{ if .IsTypeScript }}
		const { ok, value: mcastValue } = MArray.cast<unknown>(value);
		if (ok) {
			this.#{{.ctx.Name}} = mcastValue as any;
			return;
		}
		{{ else }}
		const { ok, value: mcastValue } = MArray.cast(value);
		if (ok) {
			this.#{{.ctx.Name}} = mcastValue;
			return;
		}
		{{end }}

		console.warn(
			"Cannot assing value to {{.ctx.Name}}, because it needs MArray instance or an Array.",
		);

	{{ else if or (eq .ctx.Type "collection?") (eq .ctx.Type "collection")}}

		{{ if or (eq .ctx.Type "collection?") }}
		// For nullable collection, we allow explicit undefined or null values
		if (value === null || value === undefined) {
			this.#{{.ctx.Name}} = value;

			return
		}
		{{ end }}

		// When the passed value is already an array, we check if we need to
		// cast the inner items into class instance.
		if (Array.isArray(value)) {
			if (value.length > 0 && value[0] instanceof {{.ctx.ConstructorClass}}) {
				this.#{{.ctx.Name}} = MCollection.of(value);
			} else {
				this.#{{.ctx.Name}} = MCollection.of(
					value.map((item) => new {{.ctx.ConstructorClass}}(item)),
				);
			}

			return;
		}

		// If the instance is already an MCollection, we assume it's all good.
		if (value instanceof MCollection) {
			this.#{{.ctx.Name}} = value;

			return;
		}

		// If the value is not array, and is not a MCollection, we need to be consider,
		// it might be eligible to be casted into MCollection.

		{{ if .IsTypeScript }}
		const { ok, value: mcastValue } = MCollection.cast<unknown>(value);
		if (ok) {
			this.#{{.ctx.Name}} = mcastValue as any;
			return;
		}
		{{ else }}
		const { ok, value: mcastValue } = MCollection.cast(value);
		if (ok) {
			this.#{{.ctx.Name}} = mcastValue;
			return;
		}
		{{ end }}

		console.warn(
			"Cannot assing value to {{.ctx.Name}}, because it needs MCollection instance or an Array.",
		);

	{{ else if or (eq .ctx.Type "one") (eq .ctx.Type "one?")}}
		// For objects, the sub type needs to always be instance of the sub class.
		if (value instanceof MOne) {
			this.#{{.ctx.Name}} = value
		} else if (value instanceof {{.ctx.ConstructorClass}}) {
			this.#{{.ctx.Name}} = MOne.of(value)
		} else {
			this.#{{.ctx.Name}} = MOne.of(new {{.ctx.ConstructorClass}}(value))
		}
	{{ else if or (eq .ctx.Type "object") (eq .ctx.Type "object?") }}
	 	// For objects, the sub type needs to always be instance of the sub class.
	 	if (value instanceof {{.ctx.ConstructorClass}}) {
			this.#{{.ctx.Name}} = value
		} else {
			this.#{{.ctx.Name}} = new {{.ctx.ConstructorClass}}(value)
		}
	{{ else }}
		this.#{{.ctx.Name}} = value;
	{{ end }}
}

set{{ .ctx.Upper }} (|@arg.value|) {
	this.{{.ctx.Name}} = value

	return this
}
`))
	setterjsdoc := NewJsDoc("  ")
	setterjsdoc.Add(fmt.Sprintf("@param {%v}", x.ComputedType))

	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	data := map[string]any{
		"ctx":          x,
		"IsTypeScript": isTypeScript,
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
