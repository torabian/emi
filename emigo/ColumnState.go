package emigo

// ColumnState is the per-column entry of a generated vsql "column picker"
// struct (see EmiColumn / EmiVsql.Columns in the emi compiler). Each column
// declared on a vsql query becomes a field of this type on the generated
// <Name>VsqlColumns struct, letting a query template check
// {{ if .Columns.<Key>.Selected }} without the column carrying any other
// runtime state.
type ColumnState struct {
	Selected bool `json:"selected" yaml:"selected"`
}
