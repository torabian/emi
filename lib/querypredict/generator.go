package querypredict

import (
	"bytes"
	"embed"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/torabian/emi/lib/core"
)

//go:embed query.tpl manifest.tpl exec.tpl shared.tpl
var tplFS embed.FS

type Field struct {
	Name string
	Type string
	From string
}

type QueryData struct {
	FuncName      string
	SQL           string
	Fields        []Field
	Params        string // e.g., "active bool"
	GoPackageName string
}

// helper for template: detect last element
func lastFunc(i int, a []Field) bool {
	return i == len(a)-1
}

func MakeValidGoField(col string) string {
	// Remove *, dots, parentheses, quotes, and any non-alphanumeric chars
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	clean := re.ReplaceAllString(col, " ")

	// Split words and capitalize each (CamelCase)
	parts := strings.Fields(clean)
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}

	field := strings.Join(parts, "")

	// Make sure the first letter is uppercase
	if field == "" {
		return "Col"
	}
	runes := []rune(field)
	if !unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func LoadTemplate(name string) (*template.Template, error) {
	// Start with a new template with funcs
	tpl := template.New(name).Funcs(template.FuncMap{
		"last": lastFunc,
	})

	// Parse shared templates first
	if _, err := tpl.ParseFS(tplFS, "shared.tpl"); err != nil {
		return nil, fmt.Errorf("parse shared template: %w", err)
	}

	// Parse the main template
	tplContent, err := tplFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("read template %s: %w", name, err)
	}

	if _, err := tpl.Parse(string(tplContent)); err != nil {
		return nil, fmt.Errorf("parse template %s: %w", name, err)
	}

	return tpl, nil
}

type ManifestItem struct {
	CtxName    string
	Return     string
	Name       string
	ErrorValue string
}

type Manifest struct {
	Items         []ManifestItem
	GoPackageName string
}

// ProcessQueryPredicts generates Go code files in memory
func ProcessQueryPredicts(document QueryDocument) ([]core.VirtualFile, error) {
	files := []core.VirtualFile{}
	manifest := Manifest{
		GoPackageName: "xxx",
	}

	queryTpl, err := LoadTemplate("query.tpl")
	if err != nil {
		return nil, err
	}

	execTpl, err := LoadTemplate("exec.tpl")
	if err != nil {
		return nil, err
	}

	for _, spec := range document.Queries {
		var buf bytes.Buffer
		data := QueryData{
			FuncName:      MakeValidGoField(spec.Name),
			GoPackageName: "xxx",
		}

		// Check if query contains --skip, then we do not parse it at all.
		skipParsing := strings.Contains(spec.Query, "--skip")

		sel, err := GetSelectStatementFromQuery(spec.Query)
		if !skipParsing && (err != nil && err.Error() != "not a SELECT statement") {
			if !spec.Force {
				return files, err
			}
		}

		if err != nil {
			skipParsing = true
		}

		if err == nil && !skipParsing {
			/// It means the statement is not a select - then we render exec.tpl instead
			fields := []Field{}
			fields2, err := ExtractColumnsFromSqlSelect(sel)
			if err != nil {
				return nil, err
			}

			for _, column := range fields2 {
				goName := column.GoName
				if goName == "" {
					goName = MakeValidGoField(column.ActualName)
				}

				typeof := "string"
				if column.Type != "" {
					typeof = column.Type
				}
				fields = append(fields,
					Field{
						Name: goName,
						From: column.ActualName,
						Type: typeof,
					},
				)
			}

			cleanSQL := spec.Query

			re := regexp.MustCompile(`field\(\s*([^,]+),[^\)]*\)`)
			cleanSQL = re.ReplaceAllString(cleanSQL, "$1")

			data.SQL = cleanSQL
			data.Fields = fields

			if err := queryTpl.Execute(&buf, data); err != nil {
				return nil, err
			}

			manifest.Items = append(manifest.Items, ManifestItem{
				Name:       data.FuncName,
				CtxName:    data.FuncName + "Context",
				Return:     "[]" + data.FuncName + "Row",
				ErrorValue: fmt.Sprintf("[]%vRow{}", data.FuncName),
			})
		} else {
			data.SQL = spec.Query
			if err := execTpl.Execute(&buf, data); err != nil {
				return nil, err
			}

			manifest.Items = append(manifest.Items, ManifestItem{
				Name:       data.FuncName,
				CtxName:    data.FuncName + "Context",
				Return:     "sql.Result",
				ErrorValue: "nil",
			})
		}

		file := core.VirtualFile{
			Name:         spec.Name,
			ActualScript: buf.String(),
			Extension:    ".go",
		}

		files = append(files, file)
	}

	// --- load manifest.tpl ---
	manifestTplContent, err := tplFS.ReadFile("manifest.tpl")
	if err != nil {
		return nil, err
	}

	manifestTpl, err := template.New("manifest.tpl").
		Parse(string(manifestTplContent))
	if err != nil {
		return nil, err
	}

	// --- generate manifest.go ---
	var manifestBuf bytes.Buffer
	if err := manifestTpl.Execute(&manifestBuf, manifest); err != nil {
		return nil, err
	}

	files = append(files, core.VirtualFile{
		Name:         "manifest",
		ActualScript: manifestBuf.String(),
		Extension:    ".go",
	})

	return files, nil
}
