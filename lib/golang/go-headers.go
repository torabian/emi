package golang

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/torabian/emi/lib/core"
)

type renderedHeader struct {
	PropertyName string
	Type         string
	Description  string
	Complex      string
	GetterFunc   string
	SetterFunc   string
	ActualType   string
	Dependencies []core.CodeChunkDependency
}

type goHeaderStructContext struct {
	ClassName   string
	Columns     []core.EmiHeader
	PackageName string
}

func renderHeaders(headerColumns []core.EmiHeader) ([]renderedHeader, error) {
	headers := []renderedHeader{}
	for _, header := range headerColumns {

		actualType := header.Type
		if header.Complex != "" {
			actualType = header.Complex
			if header.ComplexNamespace != "" {
				actualType = header.ComplexNamespace + "." + actualType
			}
		}

		funcName := core.ToUpper(headerNameNormalize(header.Name))

		getterFunc := fmt.Sprintf("func (h *%v) %v ", "Anonymouse", funcName)
		setterFunc := fmt.Sprintf("func (h *%v) Set%v(v %v)", "Anonymouse", funcName, header.Type)

		if header.Type == string(core.FieldTypeComplex) {
			getterFunc = fmt.Sprintf("func (h *%v) %v ", "Anonymouse", funcName)
			setterFunc = fmt.Sprintf("func (h *%v) Set%v(v %v)", "Anonymouse", funcName, actualType)
		}

		newHeader := renderedHeader{
			PropertyName: header.Name,
			Type:         header.Type,
			Complex:      header.Complex,
			Description:  header.Description,
			GetterFunc:   getterFunc,
			SetterFunc:   setterFunc,
			ActualType:   actualType,
		}

		if header.ComplexLocation != "" {
			newHeader.Dependencies = append(newHeader.Dependencies, core.CodeChunkDependency{
				Location: header.ComplexLocation,
			})
		}

		headers = append(headers, newHeader)
	}

	if len(headers) > 0 {
		commonDeps := []core.CodeChunkDependency{
			{
				Location: "fmt",
			},
			{
				Location: "net/http",
			},
			{
				Location: "strconv",
			},
		}
		headers[0].Dependencies = append(headers[0].Dependencies, commonDeps...)
	}
	return headers, nil
}

var GEN_NEST_JS_COMPATIBILITY string = "nestjs"
var GEN_TYPESCRIPT_COMPATIBILITY string = "typescript"
var GEN_SKIP_ENVELOPES string = "no-envelope"
var GEN_REACT_COMPATIBILITY string = "react"

func GoHeaderStruct(
	headerctx goHeaderStructContext,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	const tmpl = `/**
* {{.className}} struct
* Auto-generated from emi go header module
*/

type {{.className}} struct {
    headers map[string]string
}

func New{{.className}}() *{{.className}} {
    return &{{.className}}{headers: make(map[string]string)}
}

{{- range .headers }}

// {{ .Description }}
{{.SetterFunc}} {
    {{- if eq .Type "complex" }}
    // Complex type implements ToString()
    h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")] = v.ToString()
    {{- else if eq .Type "int" "int32" "int64" }}
    h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")] = strconv.FormatInt(int64(v), 10)
    {{- else if eq .Type "float32" "float64" }}
    h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")] = strconv.FormatFloat(float64(v), 'f', -1, 64)
    {{- else if eq .Type "bool" }}
    if v {
        h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")] = "true"
    } else {
        h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")] = "false"
    }
    {{- else }}
    h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")] = v
    {{- end }}
}
	
// {{ .Description }}
{{- if eq .Type "string" }}
{{.GetterFunc}}() {{.ActualType}} {
    return h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")]
}
{{- else }}
{{.GetterFunc}}() ({{.ActualType}}, error) {
    s := h.headers[http.CanonicalHeaderKey("{{.PropertyName}}")]

    {{- if eq .Type "complex" }}
    var v {{.ActualType}}
    if s != "" {
        if err := v.FromString(s); err != nil {
            return {{.ActualType}}{}, fmt.Errorf("invalid header '{{.PropertyName}}': %w", err)
        }
    }
    return v, nil

    {{- else if eq .Type "int" "int32" "int64" }}
    if s == "" {
        return 0, nil
    }
    i, err := strconv.ParseInt(s, 10, 64)
    return {{.Type}}(i), err

    {{- else if eq .Type "float32" "float64" }}
    if s == "" {
        return 0, nil
    }
    f, err := strconv.ParseFloat(s, 64)
    return {{.Type}}(f), err

    {{- else if eq .Type "bool" }}
    if s == "" {
        return false, nil
    }
    switch strings.ToLower(s) {
    case "true", "1":
        return true, nil
    case "false", "0":
        return false, nil
    default:
        return false, fmt.Errorf("invalid header '{{.PropertyName}}': %s", s)
    }

 
    {{- else }}
    return s, nil
    {{- end }}
}
{{ end }}
	
{{ end }}
 
func (h *{{.className}}) Set(key string, val string) {
    h.headers[http.CanonicalHeaderKey(key)] = val
}
func (h *{{.className}}) Get(key string) string {
    return h.headers[http.CanonicalHeaderKey(key)]
}


func (h *{{.className}}) PopulateFromHTTP(httpHeaders http.Header) {
	for k, values := range httpHeaders {
		if len(values) > 0 {
			// Canonicalize header key and take first value
			h.headers[http.CanonicalHeaderKey(k)] = values[0]
		}
	}
}

func (h *{{.className}}) WriteToHTTP(w http.ResponseWriter) {
	// Optionally, you can set default or computed values here

	for k, v := range h.headers {
		fmt.Println("Writing to response header:", k, v)
		w.Header().Set(k, v)
	}
}

`
	renderedHeaders, err := renderHeaders(headerctx.Columns)
	if err != nil {
		return nil, err
	}
	res := &core.CodeChunkCompiled{}

	for _, item := range renderedHeaders {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, item.Dependencies...)
	}

	t := template.Must(template.New("headerclass").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"headers":   renderedHeaders,
		"className": headerctx.ClassName,
		"headerctx": headerctx,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.Tokens = []core.GeneratedScriptToken{
		{
			Name:  TOKEN_ROOT_CLASS,
			Value: headerctx.ClassName,
		},
	}

	return res, nil
}

var camelCaseRe = regexp.MustCompile(`^[a-z][a-zA-Z0-9]*$`)

func headerNameNormalize(s string) string {
	if camelCaseRe.MatchString(s) {
		return s
	}
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = re.ReplaceAllString(s, " ")
	parts := strings.Fields(s)
	if len(parts) == 0 {
		return ""
	}
	result := strings.ToLower(parts[0])
	for _, p := range parts[1:] {
		if p == "" {
			continue
		}
		runes := []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		result += string(runes)
	}
	return result
}
