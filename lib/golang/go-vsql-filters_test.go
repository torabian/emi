package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestGoVsqlCompileNoFiltersOmitsFilterCode(t *testing.T) {
	vsql := core.EmiVsql{
		Name:  "search",
		Query: "select 1",
	}

	chunk, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo")
	if err != nil {
		t.Fatalf("GoVsqlCompile: %v", err)
	}
	src := string(chunk.ActualScript)
	if strings.Contains(src, "VsqlFilters") || strings.Contains(src, "FilterFieldAllowed") {
		t.Fatalf("expected no filter-related code without Filters, got:\n%s", src)
	}
}

func TestGoVsqlCompileFiltersGeneratesAllowListAndValidator(t *testing.T) {
	vsql := core.EmiVsql{
		Name:  "search",
		Query: "select 1",
		Filters: []*core.EmiField{
			{Name: "email", Type: core.FieldTypeString},
			{Name: "role", Type: core.FieldTypeString},
		},
	}

	chunk, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo")
	if err != nil {
		t.Fatalf("GoVsqlCompile: %v", err)
	}
	src := string(chunk.ActualScript)

	if !strings.Contains(src, "type SearchVsqlFilters struct") {
		t.Fatalf("expected a SearchVsqlFilters struct, got:\n%s", src)
	}
	if !strings.Contains(src, `var SearchVsqlFilterFields = []string{`) ||
		!strings.Contains(src, `"email",`) || !strings.Contains(src, `"role",`) {
		t.Fatalf("expected an allow-list of field names, got:\n%s", src)
	}
	if !strings.Contains(src, "func SearchVsqlFilterFieldAllowed(name string) bool {") {
		t.Fatalf("expected a membership-check function, got:\n%s", src)
	}
}

func TestGoVsqlCompileFiltersIndependentOfColumns(t *testing.T) {
	// Filters can be declared explicitly even when the vsql has no Columns
	// at all (nothing for it to default from).
	vsql := core.EmiVsql{
		Name:  "insertUser",
		Query: "insert into users default values",
		Filters: []*core.EmiField{
			{Name: "email", Type: core.FieldTypeString},
		},
	}

	chunk, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo")
	if err != nil {
		t.Fatalf("GoVsqlCompile: %v", err)
	}
	src := string(chunk.ActualScript)
	if !strings.Contains(src, "type InsertUserVsqlFilters struct") {
		t.Fatalf("expected filters to compile independently of columns, got:\n%s", src)
	}
	if strings.Contains(src, "VsqlColumns struct") {
		t.Fatalf("did not expect a columns picker without Columns declared, got:\n%s", src)
	}
}
