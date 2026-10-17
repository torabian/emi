package c

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// cQueryField is one scalar leaf of the query definition. Path is the C member
// path inside the struct ("filter.status") and Key the query string name
// ("filter[status]"); they are equal for a top level field.
type cQueryField struct {
	Path     string
	Key      string
	Type     string
	IsString bool
	Format   string
}

// cQueryNode is a member of the generated struct: a scalar leaf, or an object
// holding more members.
type cQueryNode struct {
	Name     string
	Leaf     *cQueryField
	Children []*cQueryNode
}

// cQueryTree builds the member tree of a query definition and returns it with
// its scalar leaves flattened, which is all the parse and apply code needs.
func cQueryTree(query []*core.EmiQueryField) (nodes []*cQueryNode, leaves []*cQueryField) {
	var build func(name string, typ core.FieldType, kids []*core.EmiField, path, key string) *cQueryNode
	build = func(name string, typ core.FieldType, kids []*core.EmiField, path, key string) *cQueryNode {
		if path != "" {
			path, key = path+"."+name, key+"["+name+"]"
		} else {
			path, key = name, name
		}
		n := &cQueryNode{Name: name}
		if typ == core.FieldTypeObject && len(kids) > 0 {
			for _, k := range kids {
				n.Children = append(n.Children, build(k.GetName(), k.GetType(), k.GetFields(), path, key))
			}
			return n
		}
		t, isString, format := cQueryFieldType(typ)
		n.Leaf = &cQueryField{Path: path, Key: key, Type: t, IsString: isString, Format: format}
		leaves = append(leaves, n.Leaf)
		return n
	}
	for _, q := range query {
		nodes = append(nodes, build(q.Name, q.Type, q.Fields, "", ""))
	}
	return nodes, leaves
}

// cQueryMembers renders the struct members of nodes.
func cQueryMembers(nodes []*cQueryNode, indent string) string {
	var b strings.Builder
	for _, n := range nodes {
		switch {
		case n.Leaf == nil:
			fmt.Fprintf(&b, "%sstruct {\n%s%s} %s;\n", indent, cQueryMembers(n.Children, indent+"    "), indent, n.Name)
		case n.Leaf.IsString:
			fmt.Fprintf(&b, "%schar* %s; /* NULL if unset */\n", indent, n.Name)
		default:
			fmt.Fprintf(&b, "%sbool %s_set;\n%s%s %s;\n", indent, n.Name, indent, n.Leaf.Type, n.Name)
		}
	}
	return b.String()
}

func cQueryFieldType(typ core.FieldType) (ctype string, isString bool, format string) {
	switch typ {
	case core.FieldTypeArray, core.FieldTypeSlice:
		// Arrays in a query string are a niche, non-scalar shape (bracket-
		// or repeat-encoding conventions vary by server) - out of scope for
		// this runtime's simple key=value appender; represented as a raw
		// (already-encoded-by-the-caller) string the same way an
		// unresolved type falls back below.
		return "char*", true, ""
	default:
		t := extractPrimitive(string(typ))
		if t == "" {
			return "char*", true, ""
		}
		f := `"%lld"`
		switch t {
		case "float", "double":
			f = `"%f"`
		case "bool":
			f = `"%d"`
		}
		return t, false, f
	}
}

// Arg is the snprintf argument for this field, see cPrintfArg.
func (f cQueryField) Arg() string { return cPrintfArg(f.Format, "self->"+f.Path) }

var cQueryTmpl = template.Must(template.New("cqs").Parse(`
typedef struct {{ .className }} {
{{ .members }}} {{ .className }};

static inline void {{ .className }}_init({{ .className }}* self) {
    if (self) memset(self, 0, sizeof(*self));
}

static inline void {{ .className }}_free_fields({{ .className }}* self) {
    if (!self) return;
{{- range .fields }}
{{- if .IsString }}
    free(self->{{ .Path }});
    self->{{ .Path }} = NULL;
{{- end }}
{{- end }}
}

/* Returns a newly allocated url with every set field appended as a query
 * parameter - the caller must free() it (url itself is never modified or
 * freed). */
static inline char* {{ .className }}_apply(const {{ .className }}* self, const char* url) {
    char* out = strdup(url);
    if (!self) return out;
{{- range .fields }}
{{- if .IsString }}
    if (self->{{ .Path }}) {
        char* next = fetchx_url_append_query(out, "{{ .Key }}", self->{{ .Path }});
        free(out);
        out = next;
    }
{{- else }}
    if (self->{{ .Path }}_set) {
        char value_buf[64];
        snprintf(value_buf, sizeof(value_buf), {{ .Format }}, {{ .Arg }});
        char* next = fetchx_url_append_query(out, "{{ .Key }}", value_buf);
        free(out);
        out = next;
    }
{{- end }}
{{- end }}
    return out;
}
`))

// CActionQueryStruct renders a struct out of an action's typed `qs:`
// definition, plus an `_apply` function that appends every set field onto a
// url as a query string.
func CActionQueryStruct(action core.EmiRpcAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	query := action.GetQuery()
	if len(query) == 0 {
		return nil, nil
	}

	nodes, fields := cQueryTree(query)

	className := fmt.Sprintf("%vQueryParams", core.ToUpper(core.NormaliseKey(action.GetName())))

	var buf bytes.Buffer
	if err := cQueryTmpl.Execute(&buf, core.H{"className": className, "fields": fields, "members": cQueryMembers(nodes, "    ")}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  core.ToSnakeCase(className),
		SuggestedExtension: ".h",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: className},
		},
	}, nil
}
