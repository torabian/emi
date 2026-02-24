package golang

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type goRenderedField struct {
	Name                string
	Type                string
	Children            []goRenderedField
	PrivateField        string
	CliCaptureStatement string
	CliName             string
}

type goFieldVariable struct {
	Name             string
	Type             string
	ConstructorClass string
	ComplexClass     string
	ComputedType     string
	IsNullable       bool
	IsNumeric        bool
	GoDoc            string
	Tags             map[string]string
	Modifier         string
	CliName          string
}

func (x goFieldVariable) Upper() string {
	return core.ToUpper(x.Name)
}

// The code which generates a statement inside a class, such as Name int `...`
func (x goFieldVariable) Compile() string {
	sequence := []string{}

	if x.GoDoc != "" {
		sequence = append(sequence, x.GoDoc)
	}

	if x.Modifier != "" && x.Modifier != "private" {
		sequence = append(sequence, x.Modifier)
	}

	varName := x.Name
	if x.Modifier == "private" {
		varName = core.ToLower(varName)
	}

	sequence = append(sequence, varName)
	sequence = append(sequence, x.ComputedType)

	// We allow each field to have custom tags, but also there is a json and yaml default
	// tags created since they are very common. Remember, if user has defined those already, we skip adding them.
	if x.Tags == nil {
		x.Tags = make(map[string]string)
	}

	if x.Tags["json"] == "" {
		x.Tags["json"] = core.ToLower(varName)
	}
	if x.Tags["yaml"] == "" {
		x.Tags["yaml"] = core.ToLower(varName)
	}

	parts := []string{}
	for k, v := range x.Tags {
		parts = append(parts, fmt.Sprintf(`%v:"%v"`, k, strings.ReplaceAll(v, `"`, `\"`)))
	}

	sequence = append(sequence, "`"+strings.Join(parts, " ")+"`")

	return strings.Join(sequence, " ")
}

func (x goFieldVariable) CliCaptureStatement() string {
	possibleType := core.FieldType(x.Type)
	statement := ""

	// Nullable fields are automatically captured by compiler, all we do is use ParseNullable from string
	// to the target object field and types are infered via generics in golang itself.
	if core.IsNullable(x.Type) {
		return fmt.Sprintf(
			"if c.IsSet(\"%v\") { \r\n emigo.ParseNullable(c.String(\"%v\"), &data.%v) \r\n}",
			x.CliName,
			x.CliName,
			x.Upper(),
		)
	}

	switch possibleType {
	case core.FieldTypeString:
		statement = fmt.Sprintf("c.String(\"%v\")", x.CliName)
	case core.FieldTypeInt, core.FieldTypeInt32, core.FieldTypeInt64:
		statement = fmt.Sprintf("%v(c.Int64(\"%v\"))", x.ComputedType, x.CliName)
	case core.FieldTypeFloat32, core.FieldTypeFloat64:
		statement = fmt.Sprintf("%v(c.Float64(\"%v\"))", x.ComputedType, x.CliName)
	case core.FieldTypeBool:
		statement = fmt.Sprintf("%v(c.Bool(\"%v\"))", x.ComputedType, x.CliName)

	// On arrays, since there is it's own function we look for that
	case core.FieldTypeArray:
		statement = fmt.Sprintf(
			"emigo.CapturePossibleArray(Cast%vFromCli, \"%v\", c)",
			strings.ReplaceAll(x.ComputedType, "[]", ""),
			x.CliName,
		)
	case core.FieldTypeSlice:
		return fmt.Sprintf(
			"if c.IsSet(\"%v\") { \r\n emigo.InflatePossibleSlice(c.String(\"%v\"), &data.%v) \r\n}",
			x.CliName,
			x.CliName,
			x.Upper(),
		)
	case core.FieldTypeComplex:
		return fmt.Sprintf(

			"if c.IsSet(\"%v\") { \r\n if u, ok := any(&data.%v).(encoding.TextUnmarshaler); ok { u.UnmarshalText([]byte(c.String(\"%v\"))) } \r\n}",
			x.CliName,
			x.Upper(),
			x.CliName,
		)
		// On arrays, since there is it's own function we look for that
	case core.FieldTypeObject:
		statement = fmt.Sprintf(
			"Cast%vFromCli(c)",
			strings.TrimSpace(x.ComputedType),
		)
	}

	if statement != "" {
		return fmt.Sprintf(
			"if c.IsSet(\"%v\") { \r\n data.%v = %v \r\n }",
			x.CliName,
			x.Upper(),
			statement,
		)
	}

	return ""
}

type fieldLike interface {
	GetType() core.FieldType
	GetModule() string
	PublicName() string
	GetTarget() string
	GetName() string
	GetCliName() string
	GetComplex() string
	GetDescription() string
	GetPrimitive() string
	GetTags() map[string]string
}

func goRenderField(
	field fieldLike,
	parentChain string,
	goctx GoCommonStructContext,
) goRenderedField {
	computedType := goFieldTypeOnNestedClasses(field, parentChain, goctx)
	isFieldNullable := core.IsNullable(string(field.GetType()))

	GoDoc := NewGoDoc("  ")
	GoDoc.Add(field.GetDescription())

	privateFieldToken := goFieldVariable{
		Name:         field.PublicName(),
		GoDoc:        GoDoc.String(),
		IsNullable:   isFieldNullable,
		CliName:      field.GetCliName(),
		Type:         string(field.GetType()),
		ComputedType: computedType,
		Tags:         field.GetTags(),
		IsNumeric:    core.IsNumericDataType(string(field.GetType())),
	}

	if field.GetComplex() != "" {
		privateFieldToken.Type = "complex"

		// This means type is complex, can be instantiated.
		if strings.Contains(field.GetComplex(), "+") {
			privateFieldToken.ComplexClass = strings.ReplaceAll(field.GetComplex(), "+", "")
		}
	}

	// + needs to be cleaned.
	privateFieldToken.ComputedType = strings.ReplaceAll(privateFieldToken.ComputedType, "+", "")
	privateField := privateFieldToken.Compile()

	return goRenderedField{
		Name:                field.GetName(),
		Type:                string(field.GetType()),
		PrivateField:        privateField,
		CliCaptureStatement: privateFieldToken.CliCaptureStatement(),
		CliName:             ComputedCliName(field),
	}
}

func ComputedCliName(x fieldLike) string {
	if x.GetCliName() != "" {
		return x.GetCliName()
	}
	return strings.ReplaceAll(core.ToSnakeCase((x.GetName())), "_", "-")
}

func goRenderFieldsShallow(
	fields []*core.EmiField,
	parentChain string,
	goctx GoCommonStructContext,

) []goRenderedField {
	out := make([]goRenderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, goRenderField(f, parentChain, goctx))
		}
	}
	return out
}

func castDtoNameToCodeChunk(dtoName string) *core.CodeChunkCompiled {
	names := strings.Split(dtoName, "|")

	chunk := &core.CodeChunkCompiled{
		ActualScript:          []byte(""),
		Tokens:                []core.GeneratedScriptToken{},
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}

	chunk.Tokens = append(chunk.Tokens,
		core.GeneratedScriptToken{Name: TOKEN_OBJ_CLASS, Value: dtoName},
		core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: dtoName},
	)

	for _, name := range names {
		directory, className := core.ParseDtoPath(strings.TrimSpace(name))

		chunk.CodeChunkDependensies = append(chunk.CodeChunkDependensies,
			core.CodeChunkDependency{
				Objects:  []string{className},
				Location: directory,
			},
		)
	}

	return chunk
}

var TOKEN_ROOT_CLASS = "root.class"
var TOKEN_OBJ_CLASS = "root.request.class"
