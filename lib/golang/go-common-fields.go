package golang

import (
	"fmt"
	"regexp"
	"sort"
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
	Description         string
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

	// collect keys
	keys := make([]string, 0, len(x.Tags))
	for k := range x.Tags {
		keys = append(keys, k)
	}

	// sort keys
	sort.Strings(keys)

	// build parts in order
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		v := x.Tags[k]
		parts = append(parts, fmt.Sprintf(`%v:"%v"`, k, strings.ReplaceAll(v, `"`, `\"`)))
	}

	sequence = append(sequence, "`"+strings.Join(parts, " ")+"`")

	return strings.Join(sequence, " ")
}

func (x goFieldVariable) CliCaptureStatement() string {
	possibleType := core.FieldType(x.Type)

	wrapIfSet := func(statement string) string {
		return fmt.Sprintf(
			"if c.IsSet(\"%v\") { \r\n data.%v = %v \r\n }",
			x.CliName,
			x.Upper(),
			statement,
		)
	}

	genericParam := func() string {
		re := regexp.MustCompile(`^emigo\.\w+\[(.+)\]$`)
		if m := re.FindStringSubmatch(x.ComputedType); len(m) > 1 {
			return m[1]
		}
		return ""
	}

	// Array and ArrayNullable, Collection and CollectionNullable
	// are being also casted with special function
	switch possibleType {
	case core.FieldTypeArray, core.FieldTypeArrayNullable:
		captureFn := "CapturePossibleArray"
		if possibleType == core.FieldTypeArrayNullable {
			captureFn = "CapturePossibleArrayNullable"
		}
		return wrapIfSet(fmt.Sprintf(
			"emigo.%v(Cast%vFromCli, \"%v\", c)",
			captureFn,
			genericParam(),
			x.CliName,
		))

	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		captureFn := "CapturePossibleCollection"
		if possibleType == core.FieldTypeCollectionNullable {
			captureFn = "CapturePossibleCollectionNullable"
		}

		pureFn := genericParam()
		externalPackage, targetItem := "", pureFn
		if strings.Contains(pureFn, ".") {
			mt := strings.Split(pureFn, ".")
			externalPackage = mt[0] + "."
			targetItem = mt[1]
		}

		return wrapIfSet(fmt.Sprintf(
			"emigo. %v(%v Cast%vFromCli, \"%v\", c)",
			captureFn,
			externalPackage,
			targetItem,
			x.CliName,
		))
	case core.FieldTypeOne, core.FieldTypeOneNullable:
		captureFn := "CapturePossibleOne"
		if possibleType == core.FieldTypeOneNullable {
			captureFn = "CapturePossibleOneNullable"
		}

		pureFn := genericParam()
		externalPackage, targetItem := "", pureFn
		if strings.Contains(pureFn, ".") {
			mt := strings.Split(pureFn, ".")
			externalPackage = mt[0] + "."
			targetItem = mt[1]
		}

		return wrapIfSet(fmt.Sprintf(
			"emigo. %v(%v Cast%vFromCli, \"%v\", c)",
			captureFn,
			externalPackage,
			targetItem,
			x.CliName,
		))
	}

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
		return wrapIfSet(fmt.Sprintf("c.String(\"%v\")", x.CliName))
	case core.FieldTypeInt, core.FieldTypeInt32, core.FieldTypeInt64:
		return wrapIfSet(fmt.Sprintf("%v(c.Int64(\"%v\"))", x.ComputedType, x.CliName))
	case core.FieldTypeFloat32, core.FieldTypeFloat64:
		return wrapIfSet(fmt.Sprintf("%v(c.Float64(\"%v\"))", x.ComputedType, x.CliName))
	case core.FieldTypeBool:
		return wrapIfSet(fmt.Sprintf("%v(c.Bool(\"%v\"))", x.ComputedType, x.CliName))
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
	case core.FieldTypeObject:
		return wrapIfSet(fmt.Sprintf("Cast%vFromCli(c)", strings.TrimSpace(x.ComputedType)))
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
	GetCliShort() string
	GetComplex() string
	GetDescription() string
	GetPrimitive() string
	GetMapKeyType() string
	GetMapValueType() string
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
		Description:         field.GetDescription(),
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

		var className = ""
		directory, className2 := core.ParseDtoPath(strings.TrimSpace(name))

		// When it has an empty space, it means that the import class is already mentioned
		if strings.Contains(name, " ") {
			m := strings.Split(name, " ")
			className = m[0]
			directory = m[1]
		} else {
			className = className2
		}

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
