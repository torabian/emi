package cpp

import (
	"reflect"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// cppActionRealms bundles every generated piece an action/remote needs,
// mirroring the sibling generators' own action realms. Dialect-dependent pieces
// (PathParameter/QueryClass/*Class/*HeadersClass) are already rendered in the
// right dialect by the time they land here - this struct itself has no dialect
// branching of its own.
type cppActionRealms struct {
	Dialect      Dialect
	ActionName   string // PascalCase, e.g. "CreateWidgetAction"
	FunctionName string
	HttpMethod   string
	Url          string
	HasUrl       bool
	IsReactive   bool

	PathParameter *core.CodeChunkCompiled
	QueryClass    *core.CodeChunkCompiled

	RequestClass       *core.CodeChunkCompiled
	RequestClassName   string
	RequestIsAmbiguous bool

	ResponseClass        *core.CodeChunkCompiled
	ResponseClassName    string
	ResponseIsAmbiguous  bool
	HasResponseShape     bool
	RequestHeadersClass  *core.CodeChunkCompiled
	ResponseHeadersClass *core.CodeChunkCompiled
}

func cppGetActionRealms(
	action core.EmiRpcAction,
	dialect Dialect,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (cppActionRealms, []core.CodeChunkDependency, error) {
	realms := cppActionRealms{Dialect: dialect}
	deps := []core.CodeChunkDependency{}

	if action == nil || reflect.ValueOf(action).IsNil() {
		return realms, deps, nil
	}

	realms.ActionName = core.ToUpper(core.NormaliseKey(action.GetName()))
	realms.FunctionName = realms.ActionName
	realms.HttpMethod = action.MethodUpper()
	realms.Url = core.RemoveTypeAnnotations(action.GetUrl())
	realms.HasUrl = realms.Url != ""
	realms.IsReactive = realms.HttpMethod == "REACTIVE"

	pathParameter, err := CppActionPathParams(action, dialect)
	if err != nil {
		return realms, nil, err
	}
	realms.PathParameter = pathParameter

	queryClass, err := CppActionQueryClass(action, dialect, ctx)
	if err != nil {
		return realms, nil, err
	}
	realms.QueryClass = queryClass

	if action.HasRequestFields() {
		reqClassName := realms.ActionName + "Request"
		fields, err := CppCommonObjectGenerator(action.GetRequestFields(), ctx, CppCommonObjectContext{
			Dialect:             dialect,
			RootClassName:       reqClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, nil, err
		}
		deps = append(deps, fields.CodeChunkDependensies...)
		realms.RequestClass = fields
		realms.RequestClassName = cppRuntimeClassName(dialect, reqClassName)
	} else if action.HasRequestDto() {
		realms.RequestClassName, realms.RequestIsAmbiguous = cppFirstDtoNameOrAmbiguous(action.GetRequestDto(), dialect)
	}

	if action.HasResponseFields() {
		resClassName := realms.ActionName + "Response"
		fields, err := CppCommonObjectGenerator(action.GetResponseFields(), ctx, CppCommonObjectContext{
			Dialect:             dialect,
			RootClassName:       resClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, nil, err
		}
		deps = append(deps, fields.CodeChunkDependensies...)
		realms.ResponseClass = fields
		realms.ResponseClassName = cppRuntimeClassName(dialect, resClassName)
		realms.HasResponseShape = true
	} else if action.HasResponseDto() {
		realms.ResponseClassName, realms.ResponseIsAmbiguous = cppFirstDtoNameOrAmbiguous(action.GetResponseDto(), dialect)
		realms.HasResponseShape = true
	}

	if action.HasRequestHeaders() {
		className := realms.ActionName + "RequestHeaders"
		chunk, err := CppHeaderClass(className, action.GetRequestHeaders(), dialect, ctx)
		if err != nil {
			return realms, nil, err
		}
		realms.RequestHeadersClass = chunk
	}
	if action.HasResponseHeaders() {
		className := realms.ActionName + "ResponseHeaders"
		chunk, err := CppHeaderClass(className, action.GetResponseHeaders(), dialect, ctx)
		if err != nil {
			return realms, nil, err
		}
		realms.ResponseHeadersClass = chunk
	}

	return realms, deps, nil
}

// cppRuntimeClassName applies the unreal dialect's `F`-prefix convention to a
// class name Emi itself generated (a request/response body flattened straight
// out of an action's `in`/`out` fields) - mirrors ueStructName, kept as its own
// tiny wrapper here since realms deal in plain strings, not EmiField trees.
func cppRuntimeClassName(dialect Dialect, name string) string {
	if dialect == DialectUnreal {
		return ueStructName(name)
	}
	return name
}

// cppFirstDtoNameOrAmbiguous mirrors the sibling generators: a dto reference can
// be `NameA|NameB` when different outcomes are possible, which a generated
// client can't disambiguate on its own.
func cppFirstDtoNameOrAmbiguous(dtoName string, dialect Dialect) (string, bool) {
	parts := strings.Split(dtoName, "|")
	names := []string{}
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		_, className := core.ParseDtoPath(p)
		names = append(names, cppRuntimeClassName(dialect, className))
	}
	if len(names) != 1 {
		fallback := "emi::EmiJson"
		if dialect == DialectUnreal {
			fallback = "TSharedPtr<FJsonValue>"
		}
		return fallback, true
	}
	return names[0], false
}
