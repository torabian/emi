package java

import (
	"reflect"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// javaActionRealms bundles every generated piece an action/remote needs,
// mirroring the sibling generators' action realms. Since Java requires one
// public type per file, anything that can expand into more than one type
// (inline request/response shapes, which may themselves contain flattened
// nested objects/enums) is a *slice* of chunks here, one file each.
type javaActionRealms struct {
	ActionName   string // PascalCase, e.g. "CreateWidgetAction"
	FunctionName string // lowerCamelCase method name, e.g. "createWidgetAction"
	HttpMethod   string
	Url          string
	HasUrl       bool
	IsReactive   bool

	PathParameter *core.CodeChunkCompiled
	QueryClass    *core.CodeChunkCompiled

	RequestClasses     []*core.CodeChunkCompiled
	RequestClassName   string
	RequestIsAmbiguous bool

	ResponseClasses      []*core.CodeChunkCompiled
	ResponseClassName    string
	ResponseIsAmbiguous  bool
	HasResponseShape     bool
	RequestHeadersClass  *core.CodeChunkCompiled
	ResponseHeadersClass *core.CodeChunkCompiled
}

func javaGetActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (javaActionRealms, error) {
	realms := javaActionRealms{}

	if action == nil || reflect.ValueOf(action).IsNil() {
		return realms, nil
	}

	realms.ActionName = core.ToUpper(core.NormaliseKey(action.GetName()))
	realms.FunctionName = core.ToLower(realms.ActionName)
	realms.HttpMethod = action.MethodUpper()
	realms.Url = core.RemoveTypeAnnotations(action.GetUrl())
	realms.HasUrl = realms.Url != ""
	realms.IsReactive = realms.HttpMethod == "REACTIVE"

	pathParameter, err := JavaActionPathParams(action)
	if err != nil {
		return realms, err
	}
	if pathParameter != nil {
		realms.PathParameter = pathParameter
	}

	queryClass, err := JavaActionQueryClass(action, ctx)
	if err != nil {
		return realms, err
	}
	if queryClass != nil {
		realms.QueryClass = queryClass
	}

	if action.HasRequestFields() {
		reqClassName := realms.ActionName + "Request"
		chunks, err := JavaCommonObjectGenerator(action.GetRequestFields(), ctx, JavaCommonObjectContext{
			RootClassName:       reqClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, err
		}
		realms.RequestClasses = chunks
		realms.RequestClassName = reqClassName
	} else if action.HasRequestDto() {
		realms.RequestClassName, realms.RequestIsAmbiguous = firstDtoNameOrAmbiguous(action.GetRequestDto())
	}

	if action.HasResponseFields() {
		resClassName := realms.ActionName + "Response"
		chunks, err := JavaCommonObjectGenerator(action.GetResponseFields(), ctx, JavaCommonObjectContext{
			RootClassName:       resClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, err
		}
		realms.ResponseClasses = chunks
		realms.ResponseClassName = resClassName
		realms.HasResponseShape = true
	} else if action.HasResponseDto() {
		realms.ResponseClassName, realms.ResponseIsAmbiguous = firstDtoNameOrAmbiguous(action.GetResponseDto())
		realms.HasResponseShape = true
	}

	if action.HasRequestHeaders() {
		className := realms.ActionName + "RequestHeaders"
		chunk, err := JavaHeaderClass(className, action.GetRequestHeaders(), ctx)
		if err != nil {
			return realms, err
		}
		realms.RequestHeadersClass = chunk
	}
	if action.HasResponseHeaders() {
		className := realms.ActionName + "ResponseHeaders"
		chunk, err := JavaHeaderClass(className, action.GetResponseHeaders(), ctx)
		if err != nil {
			return realms, err
		}
		realms.ResponseHeadersClass = chunk
	}

	return realms, nil
}

// firstDtoNameOrAmbiguous mirrors the sibling generators: a dto reference can
// be `NameA|NameB` when different outcomes are possible, which a generated
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
		return "Object", true
	}
	return names[0], false
}
