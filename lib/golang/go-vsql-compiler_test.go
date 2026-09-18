package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestGoVsqlCompileQueryVsQueryNameMutuallyExclusive(t *testing.T) {
	vsql := core.EmiVsql{
		Name:      "search",
		Query:     "select 1",
		QueryName: "search.sql",
	}

	if _, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo"); err == nil {
		t.Fatalf("expected an error when both Query and QueryName are set")
	}
}

func TestGoVsqlCompileQueryNameReadsFromFsAtRuntime(t *testing.T) {
	vsql := core.EmiVsql{
		Name:      "search",
		QueryName: "search.sql",
		Params: []*core.EmiField{
			{Name: "term", Type: core.FieldTypeString},
		},
	}

	chunk, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo")
	if err != nil {
		t.Fatalf("GoVsqlCompile: %v", err)
	}
	src := string(chunk.ActualScript)

	if !strings.Contains(src, `const SearchVsqlQueryFile = "search.sql"`) {
		t.Fatalf("expected a QueryFile constant naming search.sql, got:\n%s", src)
	}
	if strings.Contains(src, "const SearchVsqlQuery =") {
		t.Fatalf("expected no inline query constant when QueryName is set, got:\n%s", src)
	}
	if !strings.Contains(src, "func PrepareSearchVsql(fsys fs.FS, params SearchVsqlParams) (query string, args interface{}, err error)") {
		t.Fatalf("expected Prepare to take fs.FS and return an error, got:\n%s", src)
	}
	if !strings.Contains(src, "fs.ReadFile(fsys, SearchVsqlQueryFile)") {
		t.Fatalf("expected the body to read the query from fsys, got:\n%s", src)
	}

	foundFsDep := false
	for _, dep := range chunk.CodeChunkDependensies {
		if dep.Location == "io/fs" {
			foundFsDep = true
		}
	}
	if !foundFsDep {
		t.Fatalf("expected an io/fs dependency to be recorded")
	}
}

func TestGoVsqlCompileQueryNameWithColumns(t *testing.T) {
	vsql := core.EmiVsql{
		Name:      "search",
		QueryName: "search.sql",
		Columns: []*core.EmiColumn{
			{EmiField: core.EmiField{Name: "id", Type: core.FieldTypeInt64}, Selected: true},
		},
	}

	chunk, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo")
	if err != nil {
		t.Fatalf("GoVsqlCompile: %v", err)
	}
	src := string(chunk.ActualScript)

	if !strings.Contains(src, "func PrepareSearchVsql(fsys fs.FS, params SearchVsqlParams, columns SearchVsqlColumns) (query string, args interface{}, err error)") {
		t.Fatalf("expected Prepare to combine fsys with the columns picker, got:\n%s", src)
	}
	if !strings.Contains(src, "return string(b), SearchVsqlData{ SearchVsqlParams: params, Columns: columns }, nil") {
		t.Fatalf("expected the read bytes and wrapped data to be returned together, got:\n%s", src)
	}
}

func TestGoVsqlCompileWithoutQueryNameIsUnchanged(t *testing.T) {
	vsql := core.EmiVsql{
		Name:  "search",
		Query: "select 1",
	}

	chunk, err := GoVsqlCompile(vsql, core.MicroGenContext{}, nil, "github.com/torabian/emi/emigo")
	if err != nil {
		t.Fatalf("GoVsqlCompile: %v", err)
	}
	src := string(chunk.ActualScript)

	if !strings.Contains(src, "func PrepareSearchVsql(params SearchVsqlParams) (query string, args interface{}) {") {
		t.Fatalf("expected the original two-return-value Prepare shape, got:\n%s", src)
	}
	for _, dep := range chunk.CodeChunkDependensies {
		if dep.Location == "io/fs" {
			t.Fatalf("did not expect an io/fs dependency without QueryName")
		}
	}
}
