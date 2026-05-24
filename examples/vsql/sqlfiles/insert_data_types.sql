{{/*
  Probe insert. Arrays / collections are excluded from the auto-walked
  column list — in a real template you would expand them with a range
  block or drop them entirely. Everything else flows through the
  renderer's normal dispatch:
    primitive → literal
    enum → quoted string
    object → flattened prefix columns
    slice / map / any → jsonb
    complex → SQLValuer
*/}}
{{ $r := sqlFieldsExcept .
       "arrayField"
       "arrayFieldNullable"
       "collectionRef"
       "collectionRefNullable"
-}}
INSERT INTO data_types ({{ $r.Columns }})
VALUES ({{ $r.Values }});
