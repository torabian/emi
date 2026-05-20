package md

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// MdModuleDescriber reads an Emi module and produces a single markdown
// VirtualFile describing every part of it: header, enums, complexes,
// config, dtos, actions, remotes and manifests. Each section is rendered
// by its own renderer below so they can be reused or replaced individually.
func MdModuleDescriber(
	ctx core.MicroGenContext,
	m core.Emi,
) ([]core.VirtualFile, error) {

	var b strings.Builder

	b.WriteString(renderModuleHeader(m))
	b.WriteString(renderTableOfContents(m))

	sections := []func() string{
		func() string { return renderEnums(m.Enums) },
		func() string { return renderComplexes(m.Complexes) },
		func() string { return renderConfig(m.Config) },
		func() string { return renderDtos(m.Dto) },
		func() string { return renderActions(m.Actions) },
		func() string { return renderRemotes(m.Remotes) },
		func() string { return renderManifests(m.Manifests) },
	}

	for _, section := range sections {
		if s := section(); s != "" {
			b.WriteString(s)
		}
	}

	name := m.Name
	if name == "" {
		name = "module"
	}

	return []core.VirtualFile{
		{
			Name:         name,
			MimeType:     "text/markdown",
			Extension:    ".md",
			ActualScript: b.String(),
		},
	}, nil
}

// ----------------------------------------------------------------------
// Module header
// ----------------------------------------------------------------------

func renderModuleHeader(m core.Emi) string {
	var b strings.Builder

	title := m.Name
	if title == "" {
		title = "Module"
	}
	fmt.Fprintf(&b, "# %s\n\n", core.ToUpper(title))

	if m.Description != "" {
		fmt.Fprintf(&b, "> %s\n\n", m.Description)
	}

	meta := [][2]string{}
	if m.Namespace != "" {
		meta = append(meta, [2]string{"Namespace", "`" + m.Namespace + "`"})
	}
	if m.Version != "" {
		meta = append(meta, [2]string{"Version", "`" + m.Version + "`"})
	}
	if m.Name != "" {
		meta = append(meta, [2]string{"Name", "`" + m.Name + "`"})
	}

	if len(meta) > 0 {
		b.WriteString("| | |\n|---|---|\n")
		for _, row := range meta {
			fmt.Fprintf(&b, "| **%s** | %s |\n", row[0], row[1])
		}
		b.WriteString("\n")
	}

	return b.String()
}

// ----------------------------------------------------------------------
// Table of contents
// ----------------------------------------------------------------------

func renderTableOfContents(m core.Emi) string {
	var b strings.Builder

	entries := []struct {
		label string
		count int
	}{
		{"Enums", len(m.Enums)},
		{"Complex types", len(m.Complexes)},
		{"Configuration", len(m.Config)},
		{"DTOs", len(m.Dto)},
		{"Actions", len(m.Actions)},
		{"Remotes", len(m.Remotes)},
		{"Manifests", len(m.Manifests)},
	}

	anchorOf := map[string]string{
		"Enums":         "#enums",
		"Complex types": "#complex-types",
		"Configuration": "#configuration",
		"DTOs":          "#dtos",
		"Actions":       "#actions",
		"Remotes":       "#remotes",
		"Manifests":     "#manifests",
	}

	present := 0
	for _, e := range entries {
		if e.count > 0 {
			present++
		}
	}
	if present == 0 {
		return ""
	}

	b.WriteString("## Contents\n\n")
	for _, e := range entries {
		if e.count == 0 {
			continue
		}
		fmt.Fprintf(&b, "- [%s](%s) — %d\n", e.label, anchorOf[e.label], e.count)
	}
	b.WriteString("\n")

	return b.String()
}

// ----------------------------------------------------------------------
// Enums
// ----------------------------------------------------------------------

func renderEnums(items []core.EmiEnum) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Enums\n\n")

	for _, enum := range items {
		fmt.Fprintf(&b, "### `%s`\n\n", core.ToUpper(enum.Name))

		if len(enum.Fields) == 0 {
			b.WriteString("_No values._\n\n")
			continue
		}

		b.WriteString("| Key | Description |\n|---|---|\n")
		for _, f := range enum.Fields {
			fmt.Fprintf(&b, "| `%s` | %s |\n", f.GetKey(), escapeCell(f.Description))
		}
		b.WriteString("\n")
	}

	return b.String()
}

// ----------------------------------------------------------------------
// Complex types
// ----------------------------------------------------------------------

func renderComplexes(items []core.EmiComplex) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Complex types\n\n")
	b.WriteString("| Name | Namespace | Location | Compiler |\n|---|---|---|---|\n")

	for _, c := range items {
		fmt.Fprintf(
			&b,
			"| `%s` | %s | %s | %s |\n",
			c.Name,
			optionalCode(c.Namespace),
			optionalCode(c.Location),
			optionalCode(c.Compiler),
		)
	}
	b.WriteString("\n")

	return b.String()
}

// ----------------------------------------------------------------------
// Config
// ----------------------------------------------------------------------

func renderConfig(items []core.EmiConfig) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Configuration\n\n")
	b.WriteString("| Name | Type | Env | Default | Description |\n|---|---|---|---|---|\n")

	for _, c := range items {
		fmt.Fprintf(
			&b,
			"| `%s` | %s | %s | %s | %s |\n",
			c.Name,
			optionalCode(c.Type),
			optionalCode(c.Env),
			optionalCode(c.Default),
			escapeCell(c.Description),
		)
	}
	b.WriteString("\n")

	return b.String()
}

// ----------------------------------------------------------------------
// DTOs
// ----------------------------------------------------------------------

func renderDtos(items []core.EmiDto) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## DTOs\n\n")

	for _, dto := range items {
		fmt.Fprintf(&b, "### `%s`\n\n", dto.GetClassName())
		if dto.Description != "" {
			fmt.Fprintf(&b, "> %s\n\n", dto.Description)
		}
		b.WriteString(renderFieldTable(dto.Fields))
	}

	return b.String()
}

// ----------------------------------------------------------------------
// Actions
// ----------------------------------------------------------------------

func renderActions(items []*core.EmiAction) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Actions\n\n")

	for _, a := range items {
		if a == nil {
			continue
		}
		fmt.Fprintf(&b, "### `%s`\n\n", a.GetName())

		if a.Description != "" {
			fmt.Fprintf(&b, "> %s\n\n", a.Description)
		}

		b.WriteString(renderEndpointMeta(a.Method, a.Url, a.CliName, a.CliShort))
		b.WriteString(renderQueryTable(a.Query))

		if a.HasRequest() {
			b.WriteString("#### Request\n\n")
			b.WriteString(renderActionBody(a.In))
		}
		if a.HasResponse() {
			b.WriteString("#### Response\n\n")
			b.WriteString(renderActionBody(a.Out))
		}
	}

	return b.String()
}

// ----------------------------------------------------------------------
// Remotes
// ----------------------------------------------------------------------

func renderRemotes(items []*core.EmiRemote) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Remotes\n\n")

	for _, r := range items {
		if r == nil {
			continue
		}
		fmt.Fprintf(&b, "### `%s`\n\n", r.GetName())

		b.WriteString(renderEndpointMeta(r.Method, r.Url, r.CliName, r.CliShort))
		b.WriteString(renderQueryTable(r.Query))

		if r.HasRequest() {
			b.WriteString("#### Request\n\n")
			b.WriteString(renderActionBody(r.In))
		}
		if r.HasResponse() {
			b.WriteString("#### Response\n\n")
			b.WriteString(renderActionBody(r.Out))
		}
	}

	return b.String()
}

// ----------------------------------------------------------------------
// Manifests
// ----------------------------------------------------------------------

func renderManifests(items []core.EmiManifest) string {
	if len(items) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Manifests\n\n")

	for _, m := range items {
		fmt.Fprintf(&b, "### `%s`\n\n", m.Name)

		rows := [][2]string{}
		if m.Package != "" {
			rows = append(rows, [2]string{"Package", "`" + m.Package + "`"})
		}
		if m.ModPackageName != "" {
			rows = append(rows, [2]string{"Module", "`" + m.ModPackageName + "`"})
		}
		if m.Dist != "" {
			rows = append(rows, [2]string{"Dist", "`" + m.Dist + "`"})
		}
		if len(m.Types) > 0 {
			rows = append(rows, [2]string{"Types", "`" + strings.Join(m.Types, "`, `") + "`"})
		}
		if len(m.Includes) > 0 {
			rows = append(rows, [2]string{"Includes", "`" + strings.Join(m.Includes, "`, `") + "`"})
		}
		if len(m.Excludes) > 0 {
			rows = append(rows, [2]string{"Excludes", "`" + strings.Join(m.Excludes, "`, `") + "`"})
		}

		if len(rows) > 0 {
			b.WriteString("| | |\n|---|---|\n")
			for _, row := range rows {
				fmt.Fprintf(&b, "| **%s** | %s |\n", row[0], row[1])
			}
			b.WriteString("\n")
		}
	}

	return b.String()
}

// ----------------------------------------------------------------------
// Shared building blocks
// ----------------------------------------------------------------------

func renderEndpointMeta(method, url, cliName, cliShort string) string {
	rows := [][2]string{}
	if method != "" {
		rows = append(rows, [2]string{"Method", "`" + strings.ToUpper(method) + "`"})
	}
	if url != "" {
		rows = append(rows, [2]string{"URL", "`" + url + "`"})
	}
	if cliName != "" {
		rows = append(rows, [2]string{"CLI", "`" + cliName + "`"})
	}
	if cliShort != "" {
		rows = append(rows, [2]string{"CLI short", "`" + cliShort + "`"})
	}

	if len(rows) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("| | |\n|---|---|\n")
	for _, row := range rows {
		fmt.Fprintf(&b, "| **%s** | %s |\n", row[0], row[1])
	}
	b.WriteString("\n")
	return b.String()
}

func renderActionBody(body *core.EmiActionBody) string {
	if body == nil {
		return ""
	}

	var b strings.Builder

	rows := [][2]string{}
	if body.Dto != "" {
		rows = append(rows, [2]string{"DTO", "`" + body.Dto + "`"})
	}
	if body.Envelope != "" {
		rows = append(rows, [2]string{"Envelope", "`" + body.Envelope + "`"})
	}
	if body.Primitive != "" {
		rows = append(rows, [2]string{"Primitive", "`" + body.Primitive + "`"})
	}

	if len(rows) > 0 {
		b.WriteString("| | |\n|---|---|\n")
		for _, row := range rows {
			fmt.Fprintf(&b, "| **%s** | %s |\n", row[0], row[1])
		}
		b.WriteString("\n")
	}

	if len(body.Headers) > 0 {
		b.WriteString(renderHeaderTable(body.Headers))
	}
	if len(body.Fields) > 0 {
		b.WriteString(renderFieldTable(body.Fields))
	}

	return b.String()
}

func renderFieldTable(fields []*core.EmiField) string {
	if len(fields) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("| Field | Type | Description |\n|---|---|---|\n")
	for _, f := range fields {
		if f == nil {
			continue
		}
		fmt.Fprintf(
			&b,
			"| `%s` | %s | %s |\n",
			f.Name,
			formatFieldType(f),
			escapeCell(f.Description),
		)
	}
	b.WriteString("\n")
	return b.String()
}

func renderHeaderTable(headers []core.EmiHeader) string {
	if len(headers) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("**Headers**\n\n")
	b.WriteString("| Header | Type | Description |\n|---|---|---|\n")
	for _, h := range headers {
		fmt.Fprintf(
			&b,
			"| `%s` | %s | %s |\n",
			h.Name,
			optionalCode(h.Type),
			escapeCell(h.Description),
		)
	}
	b.WriteString("\n")
	return b.String()
}

func renderQueryTable(query []*core.EmiQueryField) string {
	if len(query) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("**Query parameters**\n\n")
	b.WriteString("| Name | Type | Description |\n|---|---|---|\n")
	for _, q := range query {
		if q == nil {
			continue
		}
		fmt.Fprintf(
			&b,
			"| `%s` | %s | %s |\n",
			q.Name,
			optionalCode(string(q.Type)),
			escapeCell(q.Description),
		)
	}
	b.WriteString("\n")
	return b.String()
}

// formatFieldType renders an EmiField's type in a way that reads naturally
// in a markdown table — surfacing target, primitive, complex, enum and map
// information without dumping the entire struct.
func formatFieldType(f *core.EmiField) string {
	t := string(f.Type)

	switch f.Type {
	case core.FieldTypeOne, core.FieldTypeOneNullable:
		if f.Target != "" {
			return fmt.Sprintf("`%s` → `%s`", t, f.Target)
		}
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if f.Target != "" {
			return fmt.Sprintf("`%s` → `[]%s`", t, f.Target)
		}
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		if f.Primitive != "" {
			return fmt.Sprintf("`%s` → `[]%s`", t, f.Primitive)
		}
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		if f.MapKeyOf != "" || f.MapPairOf != "" {
			return fmt.Sprintf("`%s` → `map[%s]%s`", t, defaultIfEmpty(f.MapKeyOf, "any"), defaultIfEmpty(f.MapPairOf, "any"))
		}
	case core.FieldTypeComplex:
		if f.Complex != "" {
			return fmt.Sprintf("`complex` → `%s`", f.Complex)
		}
	case core.FieldTypeEnum, core.FieldTypeEnumNullable:
		if len(f.OfType) > 0 {
			keys := make([]string, 0, len(f.OfType))
			for _, k := range f.OfType {
				if k == nil {
					continue
				}
				keys = append(keys, k.GetKey())
			}
			return fmt.Sprintf("`enum` → `%s`", strings.Join(keys, "` \\| `"))
		}
	}

	if t == "" {
		return ""
	}
	return "`" + t + "`"
}

// ----------------------------------------------------------------------
// String helpers
// ----------------------------------------------------------------------

func escapeCell(s string) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "|", "\\|")
	return strings.TrimSpace(s)
}

func optionalCode(s string) string {
	if s == "" {
		return ""
	}
	return "`" + s + "`"
}

func defaultIfEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}
