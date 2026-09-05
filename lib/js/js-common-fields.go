// Renders common JS/TS objects such as entities and DTOs.

package js

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type jsRenderedField struct {
	Name                    string
	Type                    string
	Children                []jsRenderedField
	PrivateField            string
	GetterFunc              string
	SetterCallInConstructor string
	StaticFieldItem         string
	LateInitStatement       string
	SetterFunc              string
}

func getNullableDefaultValue(field *core.EmiField) string {
	if field.Default == nil {
		return "undefined"
	}
	switch v := field.Default.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case int, int64, float64, bool:
		return fmt.Sprintf("%v", v)
	default:
		b, _ := json.Marshal(v)
		return "'" + string(b) + "'"
	}
}

func jsGetSafeFieldValue(field *core.EmiField) string {
	if field == nil {
		return "null"
	}
	if field.Default != nil {
		switch v := field.Default.(type) {
		case string:
			return fmt.Sprintf("%q", v)
		case int, int64, float64, bool:
			return fmt.Sprintf("%v", v)
		default:
			b, _ := json.Marshal(v)
			return string(b)
		}
	}

	switch field.Type {
	case core.FieldTypeArray:
		return "MArray.of([])"
	case core.FieldTypeCollection:
		return "MCollection.of([])"
	case core.FieldTypeSlice:
		return "[]"
	case "json", "object?", "embed", "computed", "any":
		return "null"
	case "string", "text":
		return `""`
	case "float", "float32", "float64", "float?", "float32?", "float64?":
		return "0.0"
	case "int", "int32", "int64", "int?", "int32?", "int64?":
		return "0"
	default:
		return ""
	}
}

type jsFieldVariable struct {
	Name                 string
	NullableDefaultValue string
	SafeDefaultValue     string
	Type                 string
	ConstructorClass     string
	ComplexClass         string
	ComputedType         string
	IsNullable           bool
	IsNumeric            bool
	JsDoc                string
	Modifier             string
}

func (x jsFieldVariable) Upper() string {
	return core.ToUpper(x.Name)
}

// generates a field variable section, such as public #fieldName: string = null;
func (x jsFieldVariable) Compile(isTypeScript bool) string {
	sequence := []string{}

	if x.JsDoc != "" {
		sequence = append(sequence, x.JsDoc)
	}

	if x.Modifier != "" && x.Modifier != "private" {
		sequence = append(sequence, x.Modifier)
	}

	varName := x.Name
	if x.Modifier == "private" {
		varName = "#" + varName
	}

	sequence = append(sequence, varName)
	if isTypeScript {

		// For enums, classes, still not sure if we have to initiate them empty regardless
		// therefor we need to make
		isLateInit := !x.IsNullable && x.SafeDefaultValue == ""

		// When field is nullable, we just put the question mark
		if x.IsNullable {
			sequence = append(sequence, "?")
		} else if isLateInit {
			sequence = append(sequence, "!")
		}

		sequence = append(sequence, ": "+x.ComputedType)

		// Bug fix: this used to unconditionally append " | null" for every
		// nullable field, even when ComputedType already spells out its own
		// "| null | undefined" (FieldTypeArrayNullable/ObjectNullable/
		// OneNullable all do, in js-common-object-class.go) - producing a
		// harmless-but-wrong doubled union like
		// "MOne<X> | null | undefined | null". Skip the generic append
		// whenever ComputedType already mentions "null" itself.
		//
		// " | null | undefined", not just " | null": a field's `?` optional
		// modifier only implies "| undefined" for *reading* it - under
		// `exactOptionalPropertyTypes` (a real, not uncommon strict-mode
		// flag) it does not extend to assignment, so a field declared only
		// "T | null" rejects the literal `undefined` its own setter
		// legitimately assigns on the "explicitly cleared" branch every
		// nullable one/collection field has (see js-setter-function.go) -
		// "Type 'undefined' is not assignable to type 'MCollection<X> |
		// null'." A field whose ComputedType is missing "null" was always
		// missing "undefined" too, for the exact same reason; this one
		// generic append was the only place that still needed telling.
		if x.IsNullable && !strings.Contains(x.ComputedType, "null") {
			sequence = append(sequence, " | null | undefined")
		}
	}

	if x.IsNullable {
		if x.NullableDefaultValue != "" {
			sequence = append(sequence, " = ")
			sequence = append(sequence, x.NullableDefaultValue)
		}
	} else if !x.IsNullable {
		if x.SafeDefaultValue != "" {
			sequence = append(sequence, " = ")
			sequence = append(sequence, x.SafeDefaultValue)
		}
	}

	return strings.Join(sequence, " ")
}

func getSelfReferencingField(field *core.EmiField, parentChain string) (bool, string) {
	if !strings.HasPrefix(field.Target, SELF_FIELD) {
		return false, ""
	}

	return true, strings.ReplaceAll(field.Target, SELF_FIELD, strings.Split(parentChain, ".")[0])
}

func jsRenderField(
	field *core.EmiField,
	parentChain string,
	fieldDepth string,
	ctx core.MicroGenContext,
) jsRenderedField {
	jsFieldType := jsFieldTypeOnNestedClasses(field, parentChain)
	tsFieldType := tsFieldTypeOnNestedClasses(field, parentChain)
	isFieldNullable := IsNullable(string(field.Type))
	isTypeScript := ctx.HasTag(Typescript)

	jsdoc := NewJsDoc("  ")
	jsdoc.Add(field.Description)
	jsdoc.Add(fmt.Sprintf("@type {%v}", jsFieldType))

	// In case of target, means the constructor class is the target, not
	// the field name with chain
	constructorClass := core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)
	if isSelf, value := getSelfReferencingField(field, parentChain); isSelf {
		constructorClass = value
		//
	} else if field.Target != "" {
		constructorClass = field.Target
	}

	privateFieldToken := jsFieldVariable{
		Modifier:             "private",
		Name:                 field.PrivateName(),
		JsDoc:                jsdoc.String(),
		NullableDefaultValue: getNullableDefaultValue(field),
		SafeDefaultValue:     jsGetSafeFieldValue(field),
		IsNullable:           isFieldNullable,
		Type:                 string(field.Type),
		ComputedType:         tsFieldType,
		IsNumeric:            IsNumericDataType(string(field.Type)),
		ConstructorClass:     constructorClass,
	}

	if field.Complex != "" {
		// complex? is the one place lib/js treats a complex field differently from
		// every other backend: it keeps the nullable suffix (rather than always
		// forcing "complex", as e.g. lib/golang's own goRenderField does) so the
		// setter template below (see js-setter-function.go) can dispatch to its
		// "complex?" branch and allow undefined/null through, exactly like every
		// other nullable field. A non-nullable "complex" field still always
		// instantiates.
		privateFieldToken.Type = "complex"
		if isFieldNullable {
			privateFieldToken.Type = "complex?"
		}

		// Bug fix: this used to only set ComplexClass (which makes the
		// generated setter - see js-setter-function.go's `.ctx.ComplexClass`
		// branch - actually construct `new X(value)`/keep an existing `X`
		// instance, instead of assigning the raw value straight through) when
		// field.Complex contained a literal "+" - the same now-removed
		// convention CollectComplexClasses used to gate the import on (see
		// its own doc comment). A plain `complex: TString` field (the normal,
		// only way any complex type is actually declared) never got a real
		// instance: `title` off the wire is JSON, not run through
		// `new TString(...)`, so it's a plain `{locale: value}` object with
		// none of TString's own methods (toString() included) - which is
		// exactly why generic display code elsewhere needs its own
		// locale-fallback lookup (getTStringValue) rather than being able to
		// rely on the value stringifying itself. Always constructing the real
		// class now - the same as "+TString" already did - fixes that at the
		// source instead of requiring every caller to special-case it.
		privateFieldToken.ComplexClass = strings.ReplaceAll(field.Complex, "+", "")
	}

	// + needs to be cleaned.
	privateFieldToken.ComputedType = strings.ReplaceAll(privateFieldToken.ComputedType, "+", "")
	privateField := privateFieldToken.Compile(isTypeScript)

	getterjsdoc := NewJsDoc("  ")
	getterjsdoc.Add(field.Description)
	getterjsdoc.Add(fmt.Sprintf("@returns {%v}", jsFieldType))
	getterFunc := getterjsdoc.String() + fmt.Sprintf("get %v () { return this.#%v }", field.Name, field.Name)

	// This is enough when type is premitive, or nullable
	setterCallInConstructor := fmt.Sprintf(
		"if (d.%v !== undefined) { this.%v = d.%v }",
		field.Name, field.Name, field.Name,
	)

	// for non-nullable fields which are late init, we need to make sure instance is being created.
	isLateInit := !privateFieldToken.IsNullable && privateFieldToken.SafeDefaultValue == ""
	lateInitStatement := ""

	if isLateInit && field.Type == core.FieldTypeObject {
		lateInitStatement = fmt.Sprintf(
			"if (!(d.%v instanceof %v)) { this.%v = new %v(d.%v || {}) }",
			field.Name, constructorClass, field.Name, constructorClass, field.Name,
		)
	} else if isLateInit && field.Type == core.FieldTypeOne {
		lateInitStatement = fmt.Sprintf(
			"if (!(d.%v instanceof %v)) { this.%v = MOne.of(new %v(d.%v || {})) }",
			field.Name, constructorClass, field.Name, constructorClass, field.Name,
		)
	}

	staticVariables := []string{}

	if field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeObjectNullable || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeArray || field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeObject || field.Type == core.FieldTypeOne || field.Type == core.FieldTypeCollectionNullable {
		staticVariables = append(
			staticVariables,
			fmt.Sprintf("%v$: '%v',", field.Name, field.Name),
		)

		withArrayIndex := ""

		if field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeArrayNullable || field.Type == core.FieldTypeArray || field.Type == core.FieldTypeCollection {
			withArrayIndex = "[:i]"
		}

		newDepth := fieldDepth + "." + field.Name
		if fieldDepth == "" {
			newDepth = field.Name
		}

		classReference := core.ToUpper(parentChain) + "." + core.ToUpper(field.Name)

		isSelf, selfFinal := getSelfReferencingField(field, parentChain)

		if !isSelf && field.Target != "" && (field.Type != core.FieldTypeEnum && field.Type != core.FieldTypeEnumNullable) {
			classReference = field.Target
		} else if isSelf {
			classReference = selfFinal
		}

		if field.Type != core.FieldTypeSlice && field.Type != core.FieldTypeSliceNullable {
			staticVariables = append(
				staticVariables,
				fmt.Sprintf(`get %v() {
					return withPrefix(
						"%v%v",
						%v.Fields
						);
						},`, field.Name, newDepth, withArrayIndex, classReference),
			)
		} else {
			staticVariables = append(
				staticVariables,
				fmt.Sprintf(`get %v() {
					return "%v%v";
						},`, field.Name, newDepth, withArrayIndex),
			)
		}
	} else {
		// For default types the string name is enough.
		staticVariables = append(
			staticVariables,
			fmt.Sprintf("%v: '%v',", field.Name, field.Name),
		)
	}

	staticFieldItem := strings.Join(staticVariables, "\r\n")
	setterFunc := privateFieldToken.CreateSetterFunction(ctx)

	return jsRenderedField{
		Name:                    field.Name,
		Type:                    string(field.Type),
		PrivateField:            privateField,
		SetterFunc:              setterFunc,
		SetterCallInConstructor: setterCallInConstructor,
		StaticFieldItem:         staticFieldItem,
		GetterFunc:              getterFunc,
		LateInitStatement:       lateInitStatement,
	}
}

func jsRenderFieldsShallow(
	fields []*core.EmiField,
	parentChain string,
	fieldDepth string,
	ctx core.MicroGenContext,
	jsctx JsCommonObjectContext,
) []jsRenderedField {
	out := make([]jsRenderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, jsRenderField(f, parentChain, fieldDepth, ctx))
		}
	}
	return out
}
