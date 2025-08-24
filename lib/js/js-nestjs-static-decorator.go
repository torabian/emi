// For each action, we produce a meta class to hold the method, default url,
// and such details, and provide a function to mimic the call with type safety.

package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type nestJsStaticFunctionUseCase int

const (
	RequestBody nestJsStaticFunctionUseCase = iota
	ResponseBody
	RequestHeaders
	ResponseHeaders
	QueryString
	QueryParams
)

type NestJsStaticDecoratorContext struct {

	// The class which will be created out of the request.
	ClassInstance string

	// represents location of the static function, to change the request section
	// which will be created
	NestJsStaticFunctionUseCase nestJsStaticFunctionUseCase
}

// If we add a static decorator to some classes, we can used them directly in nest.js
// decorators, and req, res, headers, query strings will become typesafe automatically
func JsNestJsStaticDecorator(ctxstatic NestJsStaticDecoratorContext, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	// How to do it iterte and call Compile?

	const tmpl = `/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@{{.className}}.Nest() headers: {{.className}}): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */

  {{ if .isTypeScript }}
  static Nest = createParamDecorator(
    (_data: unknown, ctx: ExecutionContext): {{ .ctx.ClassInstance }} => {
      const request = ctx.switchToHttp().getRequest<{ 
	  	body: unknown ,
		headers: { [s: string]: string }
		query: Record<string, string>;
	  }>();
      return new {{ .ctx.ClassInstance }}({{ .instanceArguments }});
    },
  );
  {{ else }}
  static Nest = createParamDecorator(
    (_data, ctx) => {
      const request = ctx.switchToHttp().getRequest();
      return new {{ .ctx.ClassInstance }}({{ .instanceArguments }});
    },
  );
  {{ end }}

`

	// For different use cases, the argument might be different
	instanceArguments := "null"
	if ctxstatic.NestJsStaticFunctionUseCase == RequestHeaders {
		instanceArguments = `Object.entries(request.headers)`
	}

	if ctxstatic.NestJsStaticFunctionUseCase == RequestBody {
		if isTypeScript {
			instanceArguments = fmt.Sprintf("request.body as Partial<%v>", ctxstatic.ClassInstance)
		} else {
			instanceArguments = ("request.body")
		}
	}

	if ctxstatic.NestJsStaticFunctionUseCase == QueryParams {
		instanceArguments = ("request.query")
	}

	t := template.Must(template.New("qsclass").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"className":         ctxstatic.ClassInstance,
		"isTypeScript":      isTypeScript,
		"instanceArguments": instanceArguments,
		"ctx":               ctxstatic,
	}); err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(buf.Bytes()),
	}

	nestjsCommon := core.CodeChunkDependency{
		Objects: []string{
			"createParamDecorator",
		},
		Location: "@nestjs/common",
	}

	if isTypeScript {
		nestjsCommon.Objects = append(nestjsCommon.Objects, "ExecutionContext")
	}

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, nestjsCommon)

	return res, nil
}
