package golang

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type goRenderedField struct {
	Name         string
	Type         string
	Children     []goRenderedField
	PrivateField string
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
	Modifier         string
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
	sequence = append(sequence, fmt.Sprintf("`json:\"%v\" yaml:\"%v\"`", core.ToLower(varName), core.ToLower(varName)))

	return strings.Join(sequence, " ")
}

type fieldLike interface {
	GetType() core.FieldType
	GetModule() string
	PublicName() string
	GetTarget() string
	GetName() string
	GetComplex() string
	GetDescription() string
	GetPrimitive() string
}

func goRenderField(
	field fieldLike,
	parentChain string,
) goRenderedField {
	computedType := goFieldTypeOnNestedClasses(field, parentChain)
	isFieldNullable := core.IsNullable(string(field.GetType()))

	GoDoc := NewGoDoc("  ")
	GoDoc.Add(field.GetDescription())

	privateFieldToken := goFieldVariable{
		Name:         field.PublicName(),
		GoDoc:        GoDoc.String(),
		IsNullable:   isFieldNullable,
		Type:         string(field.GetType()),
		ComputedType: computedType,
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
		Name:         field.GetName(),
		Type:         string(field.GetType()),
		PrivateField: privateField,
	}
}

func goRenderFieldsShallow(
	fields []*core.EmiField,
	parentChain string,

) []goRenderedField {
	out := make([]goRenderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, goRenderField(f, parentChain))
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
