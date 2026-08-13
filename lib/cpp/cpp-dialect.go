// Package cpp is the C++ target: one compiler, two dialects (see Dialect below),
// sharing every piece of logic that doesn't actually depend on which C++ world the
// output lands in (field-tree traversal, path/query param extraction, action
// realms, module wiring). Only type spelling, class declaration syntax, and the
// (de)serialization strategy differ per dialect - see cpp-type-resolver.go and the
// per-dialect field-plan files.
//
// Unlike every sibling generator (which is server-and-client, or client-only but
// still assumes a full OS/libc), this target is deliberately client-only: no HTTP
// server handlers, no CLI, no GORM/entity persistence are ever generated - a game
// client or a microcontroller is never the thing implementing the API.
package cpp

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

// Dialect selects which C++ world the generated code targets. See the package
// doc comment and README.md ("Why one compiler, two dialects") for the reasoning
// behind not splitting this into lib/cpp-unreal + lib/cpp-arduino instead.
type Dialect string

const (
	// DialectGeneric emits portable ISO C++17: std:: containers, std::optional
	// nullability, manual cJSON (de)serialization (vendored, same library the C
	// target uses) and a pluggable IEmiHttpTransport/IEmiByteStream seam with
	// ready-made backends for desktop/POSIX, ESP-IDF, and the Arduino framework.
	// This is the default.
	DialectGeneric Dialect = "generic"

	// DialectUnreal emits Unreal Engine idiomatic code, targeting UE4 and UE5
	// alike: USTRUCT(BlueprintType) classes with UPROPERTY fields (so every
	// generated dto is usable from Blueprint and inspectable in the editor for
	// free), UENUM for enums, FString/TArray/TMap for strings/arrays/collections/
	// maps, and (de)serialization via Unreal's own reflection
	// (FJsonObjectConverter) rather than hand-generated code - the same reason
	// lib/csharp never hand-generates per-field ToJson/FromJson either. A
	// nullable field is a paired value + `bXIsSet` bool (see
	// ueFieldMaybeNullable, cpp-field-plan-unreal.go) rather than a reflected
	// TOptional<T>, specifically so every generated USTRUCT compiles unmodified
	// on UE4 as well as every UE5 version, not just recent ones.
	DialectUnreal Dialect = "unreal"
)

// TagUnreal (--tags unreal) is a secondary way to select the Unreal dialect,
// alongside the primary --dialect flag - convenient when a caller is already
// threading --tags through for other purposes (e.g. --tags unreal,no-sdk).
const TagUnreal core.CTag = "unreal"

// ResolveDialect reads the `dialect` flag (--dialect unreal|generic), falling
// back to the `unreal` compiler tag, and defaults to DialectGeneric otherwise.
func ResolveDialect(ctx core.MicroGenContext) Dialect {
	if v, ok := ctx.Flags["dialect"]; ok {
		switch v {
		case "unreal", "ue", "ue5", "ue4":
			return DialectUnreal
		case "generic", "arduino", "esp-idf", "espidf", "embedded":
			return DialectGeneric
		}
	}
	if ctx.HasTag(TagUnreal) {
		return DialectUnreal
	}
	return DialectGeneric
}

// RecognizedComplex mirrors the same concept in the sibling generators: a
// custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) resolves to a real type instead of
// falling back to a raw/dynamic value. ImportLocation is a #include spelling for
// the generic dialect (e.g. `"vector3.hpp"` or `<mymath/vector3.hpp>`), or an
// Unreal module/header name for the unreal dialect.
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

func findComplexLocation(complexName string, complexes []RecognizedComplex) string {
	for _, item := range complexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}
	return ""
}

func resolveComplexSymbol(field *core.EmiField, complexes []RecognizedComplex) (symbol string, resolved bool) {
	symbol = strings.ReplaceAll(field.Complex, "+", "")
	if !strings.Contains(field.Complex, "+") {
		return "", false
	}
	for _, item := range complexes {
		if item.Symbol == symbol {
			return symbol, item.ImportLocation != ""
		}
	}
	return "", false
}

func collectComplexSymbols(fields []*core.EmiField) []string {
	var result []string
	var walk func([]*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field == nil {
				continue
			}
			if strings.Contains(field.Complex, "+") {
				result = append(result, strings.ReplaceAll(field.Complex, "+", ""))
			}
			if len(field.Fields) > 0 {
				walk(field.Fields)
			}
		}
	}
	walk(fields)
	return result
}

var TOKEN_ROOT_CLASS = "root.class"
var TOKEN_ORIGINAL_NAME = core.TOKEN_ORIGINAL_NAME
