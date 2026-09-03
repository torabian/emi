package golang

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type goQueryCliField struct {
	CliKey      string // flag key without prefix, e.g. "qs-snap-points"
	Name        string // original field name, used as the url.Values key (matches HTTP's raw query key)
	FieldName   string // exported Go struct field, matches {{ upper .Name }} in GoActionQueryParams
	Type        string
	Description string
}

func queryFieldCliName(name string) string {
	return strings.ReplaceAll(core.ToSnakeCase(name), "_", "-")
}

func buildGoQueryCliFields(qs []*core.EmiQueryField) []goQueryCliField {
	out := make([]goQueryCliField, 0, len(qs))
	for _, q := range qs {
		if q == nil {
			continue
		}
		out = append(out, goQueryCliField{
			CliKey:      "qs-" + queryFieldCliName(q.Name),
			Name:        q.Name,
			FieldName:   core.ToUpper(q.Name),
			Type:        string(q.Type),
			Description: q.Description,
		})
	}
	return out
}

// GoActionQueryParamsCli generates CLI flags plus a {Action}QueryFromCli(c *cli.Command)
// helper, mirroring {Action}QueryFromString for the CLI transport so query parameters
// bind onto a urfave v3 command exactly the way they bind off a real HTTP request.
//
// Nested "object"/"array" query fields have no separate cast function of their own (the
// query struct only ever declares them as anonymous inline structs), so they are captured
// generically as a single JSON-encoded flag unmarshalled straight into the field.
func GoActionQueryParamsCli(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	qs := action.GetQuery()
	if len(qs) == 0 {
		return nil, nil
	}

	actionName := core.ToUpper(core.NormaliseKey(action.GetName()))
	fields := buildGoQueryCliFields(qs)

	f := GetCommonFlags(ctx)

	const tmpl = `
func Get{{ .ActionName }}QueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{{ range .Fields }}
		{
			Name: prefix + "{{ .CliKey }}",
			Type: "{{ .Type }}",
			{{ if .Description }}
			Description: {{ escapeBackTick .Description }},
			{{ end }}
		},
		{{ end }}
	}
}

// {{ .ActionName }}QueryFromCli extracts and casts query parameters the same way
// {{ .ActionName }}QueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func {{ .ActionName }}QueryFromCli(c *cli.Command) {{ .ActionName }}Query {
	data := {{ .ActionName }}Query{}
	values := url.Values{}

	{{ range .Fields }}
	{{ if eq .Type "string" }}
	if c.IsSet("{{ .CliKey }}") {
		data.{{ .FieldName }} = c.String("{{ .CliKey }}")
		values.Set("{{ .Name }}", data.{{ .FieldName }})
	}
	{{ else if eq .Type "bool" }}
	if c.IsSet("{{ .CliKey }}") {
		data.{{ .FieldName }} = c.Bool("{{ .CliKey }}")
		values.Set("{{ .Name }}", strconv.FormatBool(data.{{ .FieldName }}))
	}
	{{ else if or (eq .Type "int") (eq .Type "int8") (eq .Type "int16") (eq .Type "int32") (eq .Type "int64") }}
	if c.IsSet("{{ .CliKey }}") {
		data.{{ .FieldName }} = {{ .Type }}(c.Int64("{{ .CliKey }}"))
		values.Set("{{ .Name }}", strconv.FormatInt(int64(data.{{ .FieldName }}), 10))
	}
	{{ else if or (eq .Type "float32") (eq .Type "float64") }}
	if c.IsSet("{{ .CliKey }}") {
		data.{{ .FieldName }} = {{ .Type }}(c.Float64("{{ .CliKey }}"))
		values.Set("{{ .Name }}", strconv.FormatFloat(float64(data.{{ .FieldName }}), 'f', -1, 64))
	}
	{{ else if eq .Type "slice" }}
	if c.IsSet("{{ .CliKey }}") {
		raw := c.String("{{ .CliKey }}")
		emigo.InflatePossibleSlice(raw, &data.{{ .FieldName }})
		values.Set("{{ .Name }}", raw)
	}
	{{ else if or (eq .Type "string?") (eq .Type "text?") (eq .Type "html?") (eq .Type "enum?") }}
	if c.IsSet("{{ .CliKey }}") {
		raw := c.String("{{ .CliKey }}")
		data.{{ .FieldName }} = emigo.NullableOf(raw)
		values.Set("{{ .Name }}", raw)
	}
	{{ else if eq .Type "bool?" }}
	if c.IsSet("{{ .CliKey }}") {
		v := c.Bool("{{ .CliKey }}")
		data.{{ .FieldName }} = emigo.NullableOf(v)
		values.Set("{{ .Name }}", strconv.FormatBool(v))
	}
	{{ else if or (eq .Type "int?") (eq .Type "int8?") (eq .Type "int16?") (eq .Type "int32?") (eq .Type "int64?") (eq .Type "uint?") (eq .Type "uint8?") (eq .Type "uint16?") (eq .Type "uint32?") (eq .Type "uint64?") }}
	if c.IsSet("{{ .CliKey }}") {
		v := c.Int64("{{ .CliKey }}")
		data.{{ .FieldName }} = emigo.NullableOf({{ trimQuestion .Type }}(v))
		values.Set("{{ .Name }}", strconv.FormatInt(v, 10))
	}
	{{ else if or (eq .Type "float32?") (eq .Type "float64?") }}
	if c.IsSet("{{ .CliKey }}") {
		v := c.Float64("{{ .CliKey }}")
		data.{{ .FieldName }} = emigo.NullableOf({{ trimQuestion .Type }}(v))
		values.Set("{{ .Name }}", strconv.FormatFloat(v, 'f', -1, 64))
	}
	{{ else }}
	if c.IsSet("{{ .CliKey }}") {
		raw := c.String("{{ .CliKey }}")
		json.Unmarshal([]byte(raw), &data.{{ .FieldName }})
		values.Set("{{ .Name }}", raw)
	}
	{{ end }}
	{{ end }}

	data.SetValues(values)
	return data
}
`

	t := template.Must(template.New("queryParamsCli").Funcs(core.CommonMap).Funcs(template.FuncMap{
		"trimQuestion": func(t string) string { return strings.TrimSuffix(t, "?") },
	}).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, map[string]any{
		"ActionName": actionName,
		"Fields":     fields,
	}); err != nil {
		return nil, err
	}

	deps := []core.CodeChunkDependency{
		{Location: "github.com/urfave/cli/v3"},
		{Location: f.Emigo},
		{Location: "net/url"},
	}

	if usesJSONFallback(fields) {
		deps = append(deps, core.CodeChunkDependency{Location: "encoding/json"})
	}
	if usesStrconv(fields) {
		deps = append(deps, core.CodeChunkDependency{Location: "strconv"})
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: deps,
	}, nil
}

// usesJSONFallback reports whether any field falls through to the generic
// json.Unmarshal capture branch (object/array query fields with no dedicated cast path).
func usesJSONFallback(fields []goQueryCliField) bool {
	for _, f := range fields {
		switch f.Type {
		case "string", "bool", "int", "int8", "int16", "int32", "int64",
			"float32", "float64", "slice",
			"string?", "text?", "html?", "enum?", "bool?",
			"int?", "int8?", "int16?", "int32?", "int64?",
			"uint?", "uint8?", "uint16?", "uint32?", "uint64?",
			"float32?", "float64?":
			continue
		default:
			return true
		}
	}
	return false
}

// usesStrconv reports whether any field needs strconv to stringify its value for the
// underlying url.Values.
func usesStrconv(fields []goQueryCliField) bool {
	for _, f := range fields {
		switch f.Type {
		case "bool", "int", "int8", "int16", "int32", "int64", "float32", "float64",
			"bool?", "int?", "int8?", "int16?", "int32?", "int64?",
			"uint?", "uint8?", "uint16?", "uint32?", "uint64?",
			"float32?", "float64?":
			return true
		}
	}
	return false
}
