package swift

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
	DefaultValue     string
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

	// A defaulted field has to render as `var`, not `let`: Swift drops a `let`
	// property with an initial-value default from Codable's synthesized decoding
	// *and* from the memberwise initializer entirely once it has an initializer
	// expression - confirmed empirically (Xcode 27 / Swift 6.4): `Foo(a: "x")` fails
	// to compile ("argument passed to call that takes no arguments") for
	// `struct Foo: Codable { let a: String = "" }`, and the compiler even warns
	// "immutable property will not be decoded". `var` keeps both working. Note this
	// default does NOT, by itself, make a non-Optional field tolerate a missing JSON
	// key - Swift's synthesis only ever falls back to a property's default via
	// `decodeIfPresent` for a genuinely Optional (`T?`) property; forcing that
	// Optional-ness for a field the server can legitimately omit is
	// alwaysOptionalOnTheWire's job (swift-type-resolver.go), not this default
	// value's.
	defaultStatement := ""
	keyword := "let"
	if x.DefaultValue != "" {
		defaultStatement = " = " + x.DefaultValue
		keyword = "var"
	}

	sequence = append(sequence, fmt.Sprintf(
		`%v %v: %v%v`,
		keyword,
		core.ToLower(x.Name),
		x.ComputedType,
		defaultStatement,
	))

	return strings.Join(sequence, " ")
}

func renderField(
	field *core.EmiField,
	parentChain string,
	fieldDepth string,
	ctx core.MicroGenContext,
	rootClassName string,
) renderedField {
	computedType := goFieldTypeOnNestedClasses(field, parentChain, rootClassName)
	isFieldNullable := core.IsNullable(string(field.Type))

	GoDoc := NewDocC("  ")
	GoDoc.Add(field.Description)

	// alwaysOptionalOnTheWire (see swift-type-resolver.go) can make computedType
	// Optional even when field.Type itself isn't nullable (json:"-", or an
	// array/slice/collection/map) - SwiftSafeDefaultValue only looks at field.Type,
	// so without this override a forced-Optional field would get a bare, non-nil
	// default (e.g. `var password: String? = ""` - harmless, since decode leniency
	// only ever came from the Optional-ness, not the literal default, but "unset"
	// should still read as nil, not an arbitrary empty value).
	defaultValue := SwiftSafeDefaultValue(field)
	if !isFieldNullable && alwaysOptionalOnTheWire(field) {
		defaultValue = "nil"
	}

	privateFieldToken := fieldVariable{
		Name:         field.PublicName(),
		GoDoc:        GoDoc.String(),
		IsNullable:   isFieldNullable,
		Type:         string(field.Type),
		ComputedType: computedType,
		IsNumeric:    core.IsNumericDataType(string(field.Type)),
		DefaultValue: defaultValue,
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
	goctx commonClassContext,
) []renderedField {
	out := make([]renderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, renderField(f, parentChain, fieldDepth, ctx, goctx.RootClassName))
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
