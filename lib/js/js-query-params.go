// Contains code to generate type-safe query class for javascript, and typescript definitions
// Basically, it would extend the queries class from native javascript, and add the keys there.
// It's a direct drop in for URLSearchParams, and works as expected with all libraries.

// When modifiying this file, test both js and ts definition are in sync.

package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type renderedQsField struct {
	PropertyName string
	Type         string
	Description  string
	GetterFunc   string
	SetterFunc   string
}

func renderJsTsCommonQsInfo(action core.EmiRpcAction) ([]renderedQsField, error) {
	fields := []renderedQsField{}
	for _, query := range action.GetQuery() {
		queryType, err := normalizeJsHeaderType(string(query.Type))
		if err != nil {
			return nil, err
		}
		funcName := core.ToUpper(headerNameNormalize(query.Name))
		fields = append(fields, renderedQsField{
			PropertyName: query.Name,
			Type:         queryType,
			Description:  query.Description,
			GetterFunc:   "get" + funcName,
			SetterFunc:   "set" + funcName,
		})
	}

	return fields, nil
}

// generic renderer
func renderTsJsQsClass(ctx core.MicroGenContext, action core.EmiRpcAction, fields []renderedQsField, tmpl string) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	t := template.Must(template.New("qsclass").Funcs(core.CommonMap).Parse(tmpl))
	nestJsDecorator := strings.Contains(ctx.Tags, GEN_NEST_JS_COMPATIBILITY)
	className := fmt.Sprintf("%vQueryParams", core.ToUpper(action.GetName()))

	var nestJsDecoratorRendered = ""

	if nestJsDecorator {

		nestjsStaticDecorator, err := JsNestJsStaticDecorator(NestJsStaticDecoratorContext{
			ClassInstance:               className,
			NestJsStaticFunctionUseCase: QueryParams,
		}, ctx)

		if err != nil {
			return nil, err
		}

		// Make sure to add dependencies to render tree
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, nestjsStaticDecorator.CodeChunkDependenies...)

		// Add the static function to the class bottom
		nestJsDecoratorRendered = string(nestjsStaticDecorator.ActualScript)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":                  action,
		"fields":                  fields,
		"nestJsDecoratorRendered": nestJsDecoratorRendered,
		"shouldExport":            true,
		"nestjsDecorator":         nestJsDecorator,
		"className":               className,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
		Objects: []string{
			"URLSearchParamsX",
		},
		Location: INTERNAL_SDK_JS_LOCATION,
	})

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_ROOT_CLASS,
		Value: className,
	})

	return res, nil
}

func JsActionQsClass(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	const tmpl = `/**
 * {{.className}} class
 * Auto-generated from EmiAction
 */
{{ if .shouldExport -}} export {{- end }} class {{.className}} extends URLSearchParamsX {

  {{- range .fields }}
  /**
   * {{ .Description }}
   * @returns { {{.Type}} }
   */
  {{.GetterFunc}} () {
    return this.getTyped('{{.PropertyName}}' , '{{.Type}}');
  }
  /**
   * {{ .Description }}
   * @param { {{.Type}} } value
   */
  {{.SetterFunc}} (value) {
    this.set('{{.PropertyName}}', value);
    return this;
  }
  {{- end }}

  {{ if .nestjsDecorator }}
	{{ .nestJsDecoratorRendered }}
  {{ end }}
}
`
	result, err := renderJsTsCommonQsInfo(action)
	if err != nil {
		return nil, err
	}
	return renderTsJsQsClass(ctx, action, result, tmpl)
}
