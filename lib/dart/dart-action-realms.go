package dart

import (
	"reflect"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// dartActionRealms bundles every generated piece an action/remote needs,
// mirroring lib/python/python-action-realms.go#pyActionRealms.
type dartActionRealms struct {
	ActionName   string // PascalCase, e.g. "CreateWidgetAction"
	FunctionName string // lowerCamelCase, e.g. "createWidgetAction"
	HttpMethod   string
	Url          string // with type annotations stripped (":id" not ":id int")
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

func dartGetActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (dartActionRealms, []core.CodeChunkDependency, error) {
	realms := dartActionRealms{}
	deps := []core.CodeChunkDependency{}

	if action == nil || reflect.ValueOf(action).IsNil() {
		return realms, deps, nil
	}

	realms.ActionName = core.ToUpper(core.NormaliseKey(action.GetName()))
	realms.FunctionName = core.ToLower(realms.ActionName)
	realms.HttpMethod = action.MethodUpper()
	realms.Url = core.RemoveTypeAnnotations(action.GetUrl())
	realms.HasUrl = realms.Url != ""
	realms.IsReactive = realms.HttpMethod == "REACTIVE"

	pathParameter, err := DartActionPathParams(action)
	if err != nil {
		return realms, nil, err
	}
	if pathParameter != nil {
		realms.PathParameter = pathParameter
	}

	queryClass, err := DartActionQueryClass(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if queryClass != nil {
		realms.QueryClass = queryClass
	}

	if action.HasRequestFields() {
		reqClassName := realms.ActionName + "Request"
		fields, err := DartCommonObjectGenerator(action.GetRequestFields(), ctx, DartCommonObjectContext{
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

	if action.HasResponseFields() {
		resClassName := realms.ActionName + "Response"
		fields, err := DartCommonObjectGenerator(action.GetResponseFields(), ctx, DartCommonObjectContext{
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

	if action.HasRequestHeaders() {
		className := realms.ActionName + "RequestHeaders"
		chunk, err := DartHeaderClass(className, action.GetRequestHeaders(), ctx)
		if err != nil {
			return realms, nil, err
		}
		realms.RequestHeadersClass = chunk
	}
	if action.HasResponseHeaders() {
		className := realms.ActionName + "ResponseHeaders"
		chunk, err := DartHeaderClass(className, action.GetResponseHeaders(), ctx)
		if err != nil {
			return realms, nil, err
		}
		realms.ResponseHeadersClass = chunk
	}

	return realms, deps, nil
}

// firstDtoNameOrAmbiguous mirrors lib/python's: a dto reference can be
// `NameA|NameB` when different outcomes are possible, which a generated
// client can't disambiguate on its own.
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
		return "dynamic", true
	}
	return names[0], false
}
