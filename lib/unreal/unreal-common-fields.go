package unreal

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type renderedField struct {
	Name         string
	Type         string
	Children     []renderedField
	PrivateField string
}

type fieldVariable struct {
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

func (x fieldVariable) Upper() string {
	return core.ToUpper(x.Name)
}

// The code which generates a statement inside a class, such as Name int `...`
func (x fieldVariable) Compile() string {
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

func renderField(
	field *core.EmiField,
	parentChain string,
	fieldDepth string,
	ctx core.MicroGenContext,
) renderedField {
	computedType := goFieldTypeOnNestedClasses(field, parentChain)
	isFieldNullable := core.IsNullable(string(field.Type))

	GoDoc := NewGoDoc("  ")
	GoDoc.Add(field.Description)

	privateFieldToken := fieldVariable{
		Name:         field.PublicName(),
		GoDoc:        GoDoc.String(),
		IsNullable:   isFieldNullable,
		Type:         string(field.Type),
		ComputedType: computedType,
		IsNumeric:    core.IsNumericDataType(string(field.Type)),
	}

	if field.Complex != "" {
		privateFieldToken.Type = "complex"

		// This means type is complex, can be instantiated.
		if strings.Contains(field.Complex, "+") {
			privateFieldToken.ComplexClass = strings.ReplaceAll(field.Complex, "+", "")
		}
	}

	// + needs to be cleaned.
	privateFieldToken.ComputedType = strings.ReplaceAll(privateFieldToken.ComputedType, "+", "")
	privateField := privateFieldToken.Compile()

	return renderedField{
		Name:         field.Name,
		Type:         string(field.Type),
		PrivateField: privateField,
	}
}

func renderFieldsShallow(
	fields []*core.EmiField,
	parentChain string,
	fieldDepth string,
	ctx core.MicroGenContext,
	goctx GoCommonStructContext,
) []renderedField {
	out := make([]renderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, renderField(f, parentChain, fieldDepth, ctx))
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
