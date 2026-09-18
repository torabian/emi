SELECT {{ .Columns.Cols }}
FROM users
ORDER BY id
LIMIT {{ .Limit }} OFFSET {{ .Offset }};
