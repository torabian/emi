BEGIN;

{{ $u := sqlFieldsExcept . "id" "tags" -}}
UPDATE users
SET {{ $u.Pairs }}
WHERE id = {{ sql .Id }};

{{- if .Tags.IsSet }}
{{- if eq .Tags.Operation "replace" }}
DELETE FROM user_tags WHERE user_id = {{ sql $.Id }};
{{- end }}
{{- range .Tags.Items }}
INSERT INTO user_tags (user_id, key, value)
VALUES ({{ sql $.Id }}, {{ sql .Key }}, {{ sql .Value }})
ON CONFLICT (user_id, key) DO UPDATE SET value = EXCLUDED.value;
{{- end }}
{{- end }}

COMMIT;
