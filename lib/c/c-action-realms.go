package c

import (
	"reflect"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// cActionRealms bundles every generated piece an action/remote needs,
// mirroring the sibling generators' action realms.
type cActionRealms struct {
	ActionName   string // PascalCase, e.g. "CreateWidgetAction"
	FunctionName string // snake_case function name, e.g. "create_widget_action"
	HttpMethod   string
	Url          string
	HasUrl       bool
	IsReactive   bool

	PathParameter *core.CodeChunkCompiled
	QueryStruct   *core.CodeChunkCompiled

	RequestStruct      *core.CodeChunkCompiled
	RequestTypeName    string
	RequestIsAmbiguous bool

	ResponseStruct       *core.CodeChunkCompiled
	ResponseTypeName     string
	ResponseIsAmbiguous  bool
	HasResponseShape     bool
	RequestHeadersStruct *core.CodeChunkCompiled
}

func cGetActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (cActionRealms, error) {
	realms := cActionRealms{}

	if action == nil || reflect.ValueOf(action).IsNil() {
		return realms, nil
	}

	realms.ActionName = core.ToUpper(core.NormaliseKey(action.GetName()))
	realms.FunctionName = core.ToSnakeCase(realms.ActionName)
	realms.HttpMethod = action.MethodUpper()
	realms.Url = core.RemoveTypeAnnotations(action.GetUrl())
	realms.HasUrl = realms.Url != ""
	realms.IsReactive = realms.HttpMethod == "REACTIVE"

	pathParameter, err := CActionPathParams(action)
	if err != nil {
		return realms, err
	}
	realms.PathParameter = pathParameter

	queryStruct, err := CActionQueryStruct(action, ctx)
	if err != nil {
		return realms, err
	}
	realms.QueryStruct = queryStruct

	if action.HasRequestFields() {
		reqTypeName := realms.ActionName + "Request"
		chunk, err := CCommonObjectGenerator(action.GetRequestFields(), ctx, CCommonObjectContext{
			RootTypeName:        reqTypeName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, err
		}
		realms.RequestStruct = chunk
		realms.RequestTypeName = reqTypeName
	} else if action.HasRequestDto() {
		realms.RequestTypeName, realms.RequestIsAmbiguous = firstDtoNameOrAmbiguous(action.GetRequestDto())
	}

	if action.HasResponseFields() {
		resTypeName := realms.ActionName + "Response"
		chunk, err := CCommonObjectGenerator(action.GetResponseFields(), ctx, CCommonObjectContext{
			RootTypeName:        resTypeName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, err
		}
		realms.ResponseStruct = chunk
		realms.ResponseTypeName = resTypeName
		realms.HasResponseShape = true
	} else if action.HasResponseDto() {
		realms.ResponseTypeName, realms.ResponseIsAmbiguous = firstDtoNameOrAmbiguous(action.GetResponseDto())
		realms.HasResponseShape = true
	}

	if action.HasRequestHeaders() {
		className := realms.ActionName + "RequestHeaders"
		chunk, err := CHeaderStruct(className, action.GetRequestHeaders(), ctx)
		if err != nil {
			return realms, err
		}
		realms.RequestHeadersStruct = chunk
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
		return "", true
	}
	return names[0], false
}
