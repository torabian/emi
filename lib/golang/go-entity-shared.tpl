{{ define "entityTableName" }}
{{ if .Table }}
// TableName overrides the default gorm table name for {{ .ClassName }}.
func (x *{{ .ClassName }}) TableName() string {
	return "{{ .Table }}"
}
{{ end }}
{{ end }}
