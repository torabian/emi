package python

import (
	"reflect"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// pyActionRealms bundles every generated piece an action/remote needs, mirroring
// lib/kotlin/kotlin-action-realms.go#actionRealms.
type pyActionRealms struct {
	ActionName   string // PascalCase, e.g. "CreateUserAction"
	FunctionName string // snake_case, e.g. "create_user_action"
	HttpMethod   string
	Url          string // with type annotations stripped (":id" not ":id int")
	HasUrl       bool
	IsReactive   bool

	PathParameter *core.CodeChunkCompiled // nil when the url has no placeholders
	QueryClass    *core.CodeChunkCompiled // nil when the action defines no `qs:`

	RequestClass *core.CodeChunkCompiled
	// RequestClassName is the type-hint to use for the request body argument -
	// it might reference an imported Dto instead of a locally generated class.
	RequestClassName   string
	RequestIsAmbiguous bool

	ResponseClass       *core.CodeChunkCompiled
	ResponseClassName   string
	ResponseIsAmbiguous bool
	HasResponseShape    bool

	RequestHeadersClass  *core.CodeChunkCompiled
	ResponseHeadersClass *core.CodeChunkCompiled
}

func pyGetActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (pyActionRealms, []core.CodeChunkDependency, error) {
	realms := pyActionRealms{}
	deps := []core.CodeChunkDependency{
		{Location: "dataclasses", Objects: []string{"dataclass", "field"}},
		{Location: "typing", Objects: []string{"Optional", "List", "Dict", "Any"}},
	}

	if action == nil || reflect.ValueOf(action).IsNil() {
		return realms, deps, nil
	}

	realms.ActionName = core.ToUpper(core.NormaliseKey(action.GetName()))
	realms.FunctionName = core.ToSnakeCase(realms.ActionName)
	realms.HttpMethod = action.MethodUpper()
	realms.Url = core.RemoveTypeAnnotations(action.GetUrl())
	realms.HasUrl = realms.Url != ""
	realms.IsReactive = realms.HttpMethod == "REACTIVE"

	// Path parameters
	pathParameter, err := PythonActionPathParams(action)
	if err != nil {
		return realms, nil, err
	}
	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	// Query string
	queryClass, err := PythonActionQueryClass(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if queryClass != nil {
		deps = append(deps, queryClass.CodeChunkDependensies...)
		realms.QueryClass = queryClass
	}

	// Request body
	if action.HasRequestFields() {
		reqClassName := realms.ActionName + "Request"
		fields, err := PythonCommonObjectGenerator(action.GetRequestFields(), ctx, PyCommonObjectContext{
			RootClassName:       reqClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, nil, err
		}
		deps = append(deps, fields.CodeChunkDependensies...)
		realms.RequestClass = fields
		realms.RequestClassName = reqClassName
	} else if action.HasRequestDto() {
		realms.RequestClassName, realms.RequestIsAmbiguous = firstDtoNameOrAmbiguous(action.GetRequestDto())
		deps = append(deps, DtoNameToImportDependency(action.GetRequestDto())...)
	}

	// Response body
	if action.HasResponseFields() {
		resClassName := realms.ActionName + "Response"
		fields, err := PythonCommonObjectGenerator(action.GetResponseFields(), ctx, PyCommonObjectContext{
			RootClassName:       resClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, nil, err
		}
		deps = append(deps, fields.CodeChunkDependensies...)
		realms.ResponseClass = fields
		realms.ResponseClassName = resClassName
		realms.HasResponseShape = true
	} else if action.HasResponseDto() {
		realms.ResponseClassName, realms.ResponseIsAmbiguous = firstDtoNameOrAmbiguous(action.GetResponseDto())
		deps = append(deps, DtoNameToImportDependency(action.GetResponseDto())...)
		realms.HasResponseShape = true
	}

	// Headers
	if action.HasRequestHeaders() {
		className := realms.ActionName + "RequestHeaders"
		chunk, err := PythonHeaderClass(className, action.GetRequestHeaders(), ctx)
		if err != nil {
			return realms, nil, err
		}
		realms.RequestHeadersClass = chunk
	}
	if action.HasResponseHeaders() {
		className := realms.ActionName + "ResponseHeaders"
		chunk, err := PythonHeaderClass(className, action.GetResponseHeaders(), ctx)
		if err != nil {
			return realms, nil, err
		}
		realms.ResponseHeadersClass = chunk
	}

	return realms, deps, nil
}

// firstDtoNameOrAmbiguous mirrors JS's "isAmbiguousCreator" handling: a dto
// reference can be `NameA|NameB` when different outcomes are possible, which
// a generated client can't disambiguate on its own. We surface the first
// name as the type-hint but flag it so the render step can fall back to
// returning the raw parsed payload instead of instantiating a class.
func firstDtoNameOrAmbiguous(dtoName string) (string, bool) {
	parts := strings.Split(dtoName, "|")
	names := []string{}
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		_, className := core.ParseDtoPath(p)
		names = append(names, className)
	}

	if len(names) != 1 {
		return "Any", true
	}
	return names[0], false
}
