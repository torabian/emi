BEGIN;

{{ $u := sqlFieldsExcept . "tags" -}}
INSERT INTO users ({{ $u.Columns }})
VALUES ({{ $u.Values }});

{{ $a := sqlFields .Address -}}
INSERT INTO addresses ({{ $a.Columns }})
VALUES ({{ $a.Values }});

{{- if .Tags }}
INSERT INTO user_tags (user_id, key, value) VALUES
{{- range $i, $t := .Tags.Items }}{{ if $i }},{{ end }}
  ({{ sql $.Id }}, {{ sql $t.Key }}, {{ sql $t.Value }})
{{- end }};
{{- end }}

COMMIT;
