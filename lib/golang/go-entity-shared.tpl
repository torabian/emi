{{ define "entityTableName" }}
{{ if .Table }}
// TableName overrides the default gorm table name for {{ .ClassName }}.
func (x *{{ .ClassName }}) TableName() string {
	return "{{ .Table }}"
}
{{ end }}
{{ end }}

{{ define "entityBeforeCreateHooks" }}
{{ range .IdentityStructs }}
// BeforeCreate assigns UniqueId a random UUID (v4) if the caller hasn't already set one -
// gorm calls this automatically from every Create()/Save() insert path (including the
// has-many/many-to-many reconcile helpers in emigorm, which persist child rows via
// tx.Save() directly rather than through a generated *CreateFn). This replaces relying
// on a DB-level column default (e.g. Postgres's gen_random_uuid()): sqlite and MySQL
// have no dialect-portable equivalent, so assigning it here instead works identically
// across every gorm dialect, with no SQL default expression at all.
func (x *{{ . }}) BeforeCreate(tx *gorm.DB) error {
	if x.UniqueId == "" {
		x.UniqueId = emigo.NewUUIDv4()
	}
	return nil
}
{{ end }}
{{ end }}
