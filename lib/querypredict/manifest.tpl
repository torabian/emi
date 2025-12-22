package artifacts

import "database/sql"

type Manifest struct {
	DB            *sql.DB
	FilterResolver func(string) (string, error)
}

{{ range .Items }}
func (m *Manifest) {{ .Name }}(ctx {{ .CtxName }}) ({{ .Return }}, error) {
	if m.FilterResolver != nil {
		filter, err := m.FilterResolver(ctx.Filter)
		if err != nil {
			return {{ .ErrorValue }}, err
		}
		ctx.Filter = filter
	}
	return {{ .Name }}(m.DB, ctx)
}
{{ end }}
