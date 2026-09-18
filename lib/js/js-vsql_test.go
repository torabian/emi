package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestJsVsqlGenerateParamsOnly(t *testing.T) {
	vsql := core.EmiVsql{
		Name: "insertUser",
		Params: []*core.EmiField{
			{Name: "email", Type: core.FieldTypeString},
		},
	}

	chunks, err := JsVsqlGenerate(vsql, core.MicroGenContext{}, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}
	if len(chunks) != 1 {
		t.Fatalf("expected only a Params chunk when Columns is empty, got %d", len(chunks))
	}

	src := string(chunks[0].ActualScript)
	if !strings.Contains(src, "class InsertUserVsqlParams") {
		t.Fatalf("expected a InsertUserVsqlParams class, got:\n%s", src)
	}
	if !strings.Contains(src, "email") {
		t.Fatalf("expected the email field to be present, got:\n%s", src)
	}
}

func TestJsVsqlGenerateParamsAndRow(t *testing.T) {
	vsql := core.EmiVsql{
		Name: "selectUserById",
		Params: []*core.EmiField{
			{Name: "id", Type: core.FieldTypeInt64},
		},
		Columns: []*core.EmiColumn{
			{EmiField: core.EmiField{Name: "id", Type: core.FieldTypeInt64}, Selected: true},
			{EmiField: core.EmiField{Name: "email", Type: core.FieldTypeString}, Selected: false},
		},
	}

	chunks, err := JsVsqlGenerate(vsql, core.MicroGenContext{}, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}
	if len(chunks) != 2 {
		t.Fatalf("expected a Params chunk and a Row chunk, got %d", len(chunks))
	}

	paramsSrc := string(chunks[0].ActualScript)
	if !strings.Contains(paramsSrc, "class SelectUserByIdVsqlParams") {
		t.Fatalf("expected SelectUserByIdVsqlParams, got:\n%s", paramsSrc)
	}

	rowSrc := string(chunks[1].ActualScript)
	if !strings.Contains(rowSrc, "class SelectUserByIdVsqlRow") {
		t.Fatalf("expected SelectUserByIdVsqlRow, got:\n%s", rowSrc)
	}
	if !strings.Contains(rowSrc, "email") {
		t.Fatalf("expected the row to carry every declared column regardless of Selected, got:\n%s", rowSrc)
	}
}

func TestJsVsqlGenerateNestedObjectColumn(t *testing.T) {
	vsql := core.EmiVsql{
		Name: "selectUserById",
		Columns: []*core.EmiColumn{
			{
				EmiField: core.EmiField{
					Name: "preferences",
					Type: core.FieldTypeObject,
					Fields: []*core.EmiField{
						{Name: "theme", Type: core.FieldTypeString},
					},
				},
				Selected: true,
			},
		},
	}

	chunks, err := JsVsqlGenerate(vsql, core.MicroGenContext{}, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}
	rowSrc := string(chunks[1].ActualScript)
	if !strings.Contains(rowSrc, "static Preferences") {
		t.Fatalf("expected a nested Preferences class, got:\n%s", rowSrc)
	}
	if !strings.Contains(rowSrc, "theme") {
		t.Fatalf("expected the nested field to carry through, got:\n%s", rowSrc)
	}
}

func TestJsVsqlGenerateFiltersChunk(t *testing.T) {
	vsql := core.EmiVsql{
		Name: "selectUserById",
		Filters: []*core.EmiField{
			{Name: "email", Type: core.FieldTypeString},
			{Name: "role", Type: core.FieldTypeString},
		},
	}

	chunks, err := JsVsqlGenerate(vsql, core.MicroGenContext{}, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}
	if len(chunks) != 2 {
		t.Fatalf("expected a Params chunk and a Filters chunk, got %d", len(chunks))
	}

	filtersSrc := string(chunks[1].ActualScript)
	if !strings.Contains(filtersSrc, "class SelectUserByIdVsqlFilters") {
		t.Fatalf("expected a SelectUserByIdVsqlFilters class, got:\n%s", filtersSrc)
	}
}

func TestJsVsqlGenerateFiltersTypeScriptAppendix(t *testing.T) {
	vsql := core.EmiVsql{
		Name: "selectUserById",
		Filters: []*core.EmiField{
			{Name: "email", Type: core.FieldTypeString},
			{Name: "role", Type: core.FieldTypeString},
		},
	}

	ctx := core.MicroGenContext{Tags: "typescript"}
	chunks, err := JsVsqlGenerate(vsql, ctx, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}

	src := string(chunks[1].ActualScript)
	for _, want := range []string{
		`export type SelectUserByIdVsqlFilterField = "email" | "role";`,
		"export type SelectUserByIdVsqlFilterVarNode = {",
		"var?: never",
	} {
		if !strings.Contains(src, want) {
			t.Fatalf("expected to find %q, got:\n%s", want, src)
		}
	}
}

func TestJsVsqlGenerateFiltersOmitsTypeScriptAppendixForPlainJs(t *testing.T) {
	vsql := core.EmiVsql{
		Name: "selectUserById",
		Filters: []*core.EmiField{
			{Name: "email", Type: core.FieldTypeString},
		},
	}

	chunks, err := JsVsqlGenerate(vsql, core.MicroGenContext{}, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}

	src := string(chunks[1].ActualScript)
	if strings.Contains(src, "VsqlFilterField") {
		t.Fatalf("did not expect the TS-only field-union appendix without --tags typescript, got:\n%s", src)
	}
}

func TestJsVsqlGenerateNoFiltersChunkWhenEmpty(t *testing.T) {
	vsql := core.EmiVsql{
		Name:   "insertUser",
		Params: []*core.EmiField{{Name: "email", Type: core.FieldTypeString}},
	}

	chunks, err := JsVsqlGenerate(vsql, core.MicroGenContext{}, nil)
	if err != nil {
		t.Fatalf("JsVsqlGenerate: %v", err)
	}
	if len(chunks) != 1 {
		t.Fatalf("expected only a Params chunk with no Filters, got %d", len(chunks))
	}
}
