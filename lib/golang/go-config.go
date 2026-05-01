package golang

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoConfigGenerate(
	configs []core.EmiConfig,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {

	if len(configs) == 0 {
		return nil, nil
	}

	res := core.CodeChunkCompiled{}

	tmpl := `/**
* Configuration generator
*/



{{ define "configFields" }}
  {{ $fields := index . 0 }}
  {{ $prefix := index . 1 }}
  {{ range $fields }}
    // {{ .Description }}
    {{ if or (eq .Type "string") (eq .Type "")}}
      {{ upper $prefix }}{{ upper .Name }} string $bt$envconfig:"{{- if .Env -}}{{ .Env }}{{else}}{{ snakeUpper .Name }}{{end}}" description:"{{ escape .Description}}"$bt$
    {{ end }}
    {{ if or (eq .Type "int64") }}
      {{ upper $prefix }}{{ upper .Name }} int64 $bt$envconfig:"{{- if .Env -}}{{ .Env }}{{else}}{{ snakeUpper .Name }}{{end}}" description:"{{ escape .Description}}"$bt$
    {{ end }}
    {{ if or (eq .Type "float64") }}
      {{ upper $prefix }}{{ upper .Name }} float64 $bt$envconfig:"{{- if .Env -}}{{ .Env }}{{else}}{{ snakeUpper .Name }}{{end}}" description:"{{ escape .Description}}"$bt$
    {{ end }}
    {{ if or (eq .Type "int") }}
      {{ upper $prefix }}{{ upper .Name }} int $bt$envconfig:"{{- if .Env -}}{{ .Env }}{{else}}{{ snakeUpper .Name }}{{end}}" description:"{{ escape .Description}}"$bt$
    {{ end }}
    {{ if or (eq .Type "bool") (eq .Type "boolean") }}
      {{ upper $prefix }}{{ upper .Name }} bool $bt$envconfig:"{{- if .Env -}}{{ .Env }}{{else}}{{ snakeUpper .Name }}{{end}}" description:"{{ escape .Description}}"$bt$
    {{ end }}
    {{ if or (eq .Type "int32") }}
      {{ upper $prefix }}{{ upper .Name }} int32 $bt$envconfig:"{{- if .Env -}}{{ .Env }}{{else}}{{ snakeUpper .Name }}{{end}}" description:"{{ escape .Description}}"$bt$
    {{ end }}
  {{ end }}
{{ end }}



{{ if .fields }}
type Config struct {
  {{ template "configFields" (arr .fields "") }}
}


func GetConfigCliFlags() []cli.Flag {
	return []cli.Flag{
    {{ range .fields }} 
      {{ if or (eq .Type "string") (eq .Type "")}}
        cli.StringFlag{
          Name:  "{{ .DashedName }}",
          Usage: "{{ .Description }}",
        },
      {{ end }}
      {{ if or (eq .Type "int64") }}
        cli.Int64Flag{
          Name:  "{{ .DashedName }}",
          Usage: "{{ .Description }}",
        },
      {{ end }}
      {{ if or (eq .Type "float64") }}
        cli.Float64Flag{
          Name:  "{{ .DashedName }}",
          Usage: "{{ .Description }}",
        },
      {{ end }}
      {{ if or (eq .Type "int") }}
        cli.IntFlag{
          Name:  "{{ .DashedName }}",
          Usage: "{{ .Description }}",
        },
      {{ end }}
      {{ if or (eq .Type "bool") (eq .Type "boolean") }}
        cli.BoolFlag{
          Name:  "{{ .DashedName }}",
          Usage: "{{ .Description }}",
        },
      {{ end }}
      {{ if or (eq .Type "int32") }}
        cli.Int32Flag{
          Name:  "{{ .DashedName }}",
          Usage: "{{ .Description }}",
        },
      {{ end }}
    {{ end }}
	}
}


func CastConfigFromCli(config *Config, c emigo.CliCastable) {
  {{ range .fields }} 
    if c.IsSet("{{ .DashedName }}") {
      {{ if or (eq .Type "string") (eq .Type "")}}
        config.{{ upper .Name }} = c.String("{{ .DashedName }}")
      {{ end }}
      {{ if or (eq .Type "int64") }}
        config.{{ upper .Name }} = c.Int64("{{ .DashedName }}")
      {{ end }}
      {{ if or (eq .Type "float64") }}
        config.{{ upper .Name }} = c.Float64("{{ .DashedName }}")
      {{ end }}
      {{ if or (eq .Type "int") }}
        config.{{ upper .Name }} = c.Int("{{ .DashedName }}")
      {{ end }}
      {{ if or (eq .Type "bool") (eq .Type "boolean") }}
        config.{{ upper .Name }} = c.Bool("{{ .DashedName }}")
      {{ end }}
      {{ if or (eq .Type "int32") }}
        config.{{ upper .Name }} = c.Int32("{{ .DashedName }}")
      {{ end }}
    }
  {{ end }}
}


func GetConfigCli() []cli.Command {
	return []cli.Command{
    {{ range .fields }}
		{
			Name:  "{{ .DashedName }}",
			Usage: "{{ .Description }} ({{ if or (eq .Type "string") (eq .Type "")}}string{{else}}{{.Type}}{{end}})",

      Subcommands: []cli.Command{
				{
					Name: "get",
					Action: func(c *cli.Context) error {
						fmt.Println(config.{{ upper .Name }})
						return nil
					},
				},
				{
					Name: "set",
					Action: func(c *cli.Context) error {
            {{ if or (eq .Type "bool") (eq .Type "boolean") }}
              return emigo.ConfigSetBoolean(c, config.{{ upper .Name }}, func(value bool) {
                config.{{ upper .Name }} = value
                config.Save(".env")
              })
            {{ end }}
            {{ if or (eq .Type "string") (eq .Type "")}}
              return emigo.ConfigSetString(c, config.{{ upper .Name }}, func(value string) {
                config.{{ upper .Name }} = value
                config.Save(".env")
              })
            {{ end }}
            {{ if or (eq .Type "int64")}}
              return emigo.ConfigSetInt64(c, config.{{ upper .Name }}, func(value int64) {
                config.{{ upper .Name }} = value
                config.Save(".env")
              })
            {{ end }}
            {{ if or (eq .Type "float64")}}
              return emigo.ConfigSetFloat64(c, config.{{ upper .Name }}, func(value float64) {
                config.{{ upper .Name }} = value
                config.Save(".env")
              })
            {{ end }}
            {{ if or (eq .Type "int")}}
              return emigo.ConfigSetInt(c, config.{{ upper .Name }}, func(value int) {
                config.{{ upper .Name }} = value
                config.Save(".env")
              })
            {{ end }}

            return nil
					},
				},
			},
		},
    {{ end }}
	}

}


// The config is usually populated by env vars on LoadConfiguration
var config Config = Config{
  {{ range .fields}}
    {{ if .Default }}
      {{ if or (eq .Type "string") (eq .Type "") }}
        {{ upper .Name }}: "{{ .Default }}",
      {{ else }}
        {{ upper .Name }}: {{ .Default }},
      {{ end }}
    {{ end }}
  {{ end }}
}

/**
You can call this function on first line of your main function.
This is different from fireback configuration (for now), you can
define config: in module3 file, similar to fields in entities,
and we generate the config struct and this function would read .env.local,
.env.prod, etc - depending on the ENV=xxx env variable.
**/
func LoadConfiguration() Config {
	emigo.HandleEnvVars(&config)
	return config
}

func (x *Config) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func (x *Config) Save(filepath string) error {
	return emigo.SaveEnvFile(x, filepath)
}


{{ end }}
`

	tmpl = strings.ReplaceAll(tmpl, "$bt$", "`")

	t := template.Must(template.New("config_generator").Funcs(core.CommonMap).Parse(tmpl))

	type Flags struct {
		Emigo       string `json:"emigo,omitempty"`
		PackageName string `json:"pkg,omitempty"`
	}
	var f Flags = Flags{
		Emigo:       "github.com/torabian/emi/emigo",
		PackageName: DEFAULT_GO_PACKAGE,
	}

	if val, ok := ctx.Flags["emigo"]; ok && val != "" {
		f.Emigo = val
	}

	if val, ok := ctx.Flags["pkg"]; ok && val != "" {
		f.PackageName = val
	}

	if len(configs) > 0 {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			[]core.CodeChunkDependency{
				{
					Location: "encoding/json",
				},
				{
					Location: "fmt",
				},
				{
					Location: "github.com/urfave/cli",
				},
				{
					Location: f.Emigo,
				},
			}...,
		)

	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"fields": configs,
	}); err != nil {
		return nil, err
	}

	res.SuggestedFileName = "Configuration"
	res.ActualScript = buf.Bytes()
	res.SuggestedExtension = ".go"

	return &res, nil
}
