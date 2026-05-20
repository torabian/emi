package openapi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/torabian/emi/lib/core"
	"gopkg.in/yaml.v3"
)

// OpenAPIModuleDescriber reads an Emi module and produces an OpenAPI 3
// document, returned as two VirtualFiles: one YAML, one JSON. The structure
// mirrors lib/md/md-describer.go — small helpers do the per-section work so
// each piece can be reused or replaced independently.
func OpenAPIModuleDescriber(
	ctx core.MicroGenContext,
	m core.Emi,
) ([]core.VirtualFile, error) {

	doc, err := BuildOpenAPIDoc(m)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal openapi to json: %w", err)
	}

	yamlBytes, err := jsonToYAML(jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("marshal openapi to yaml: %w", err)
	}

	name := m.Name
	if name == "" {
		name = "openapi"
	}

	files := []core.VirtualFile{
		{
			Name:         name,
			MimeType:     "application/yaml",
			Extension:    ".openapi.yaml",
			ActualScript: string(yamlBytes),
		},
	}
	fmt.Println(1, ctx.Flags["json"])

	if v, ok := ctx.Flags["json"]; ok && v != "" && v != "false" {
		files = append(files, core.VirtualFile{
			Name:         name,
			MimeType:     "application/json",
			Extension:    ".openapi.json",
			ActualScript: string(jsonBytes),
		})
	}

	return files, nil
}

// BuildOpenAPIDoc converts an Emi module into a kin-openapi document.
func BuildOpenAPIDoc(m core.Emi) (*openapi3.T, error) {

	components := &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}

	registerEnumSchemas(m.Enums, components)
	registerComplexSchemas(m.Complexes, components)
	registerDtoSchemas(m.Dto, components)

	paths := openapi3.NewPaths()

	for _, action := range m.Actions {
		if action == nil || action.Url == "" {
			continue
		}
		addActionToPaths(action, paths, components)
	}

	for _, remote := range m.Remotes {
		if remote == nil || remote.Url == "" {
			continue
		}
		addRemoteToPaths(remote, paths, components)
	}

	doc := &openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       firstNonEmpty(m.Name, "Emi Module"),
			Version:     firstNonEmpty(m.Version, "1.0.0"),
			Description: m.Description,
		},
		Paths:      paths,
		Components: components,
	}

	return doc, nil
}

// ----------------------------------------------------------------------
// Component schema registration
// ----------------------------------------------------------------------

func registerEnumSchemas(items []core.EmiEnum, components *openapi3.Components) {
	for _, e := range items {
		values := make([]any, 0, len(e.Fields))
		for _, f := range e.Fields {
			values = append(values, f.GetKey())
		}
		schema := &openapi3.Schema{
			Type: &openapi3.Types{openapi3.TypeString},
			Enum: values,
		}
		components.Schemas[e.GetName()] = openapi3.NewSchemaRef("", schema)
	}
}

// registerComplexSchemas stubs out complex types with an opaque object
// schema. Complex types' field structures aren't known to Emi, but
// publishing a placeholder lets references in other schemas resolve.
func registerComplexSchemas(items []core.EmiComplex, components *openapi3.Components) {
	for _, c := range items {
		schema := &openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeObject},
			Description: fmt.Sprintf("Complex type %q (defined in %s)", c.Name, c.Location),
		}
		components.Schemas[c.Name] = openapi3.NewSchemaRef("", schema)
	}
}

func registerDtoSchemas(items []core.EmiDto, components *openapi3.Components) {
	for _, dto := range items {
		schema := buildObjectSchema(dto.Fields, dto.Description)
		components.Schemas[dto.GetClassName()] = openapi3.NewSchemaRef("", schema)
	}
}

// ----------------------------------------------------------------------
// Actions and remotes → paths
// ----------------------------------------------------------------------

func addActionToPaths(a *core.EmiAction, paths *openapi3.Paths, _ *openapi3.Components) {
	url, pathParams := normalizeURL(a.Url)
	op := buildOperation(a.Description, a.Description, a.Query, pathParams, a.In, a.Out)
	attachOperation(paths, url, methodOf(a.Method, a.Url), op)
}

func addRemoteToPaths(r *core.EmiRemote, paths *openapi3.Paths, _ *openapi3.Components) {
	url, pathParams := normalizeURL(r.Url)
	op := buildOperation(r.Name, r.Name, r.Query, pathParams, r.In, r.Out)
	attachOperation(paths, url, methodOf(r.Method, r.Url), op)
}

func attachOperation(paths *openapi3.Paths, url, method string, op *openapi3.Operation) {
	item := paths.Value(url)
	if item == nil {
		item = &openapi3.PathItem{}
	}
	switch method {
	case "GET":
		item.Get = op
	case "POST":
		item.Post = op
	case "PUT":
		item.Put = op
	case "PATCH":
		item.Patch = op
	case "DELETE":
		item.Delete = op
	case "OPTIONS":
		item.Options = op
	case "HEAD":
		item.Head = op
	default:
		// Reactive / WebRTC / WebSocket — surface as POST to keep
		// the spec valid; richer modelling would need x-extensions.
		item.Post = op
	}
	paths.Set(url, item)
}

func buildOperation(
	summary, description string,
	query []*core.EmiQueryField,
	pathParams []pathParam,
	in *core.EmiActionBody,
	out *core.EmiActionBody,
) *openapi3.Operation {

	op := &openapi3.Operation{
		Summary:     summary,
		Description: description,
	}

	for _, p := range pathParams {
		op.Parameters = append(op.Parameters, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name:     p.Name,
				In:       openapi3.ParameterInPath,
				Required: true,
				Schema:   openapi3.NewSchemaRef("", primitiveSchema(p.Type)),
			},
		})
	}

	for _, q := range query {
		if q == nil {
			continue
		}
		op.Parameters = append(op.Parameters, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name:        q.Name,
				In:          openapi3.ParameterInQuery,
				Description: q.Description,
				Schema:      openapi3.NewSchemaRef("", queryFieldSchema(q)),
			},
		})
	}

	if in != nil {
		for _, h := range in.Headers {
			op.Parameters = append(op.Parameters, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					Name:        h.Name,
					In:          openapi3.ParameterInHeader,
					Description: h.Description,
					Schema:      openapi3.NewSchemaRef("", headerSchema(h)),
				},
			})
		}

		if bodyRef := bodySchemaRef(in); bodyRef != nil {
			op.RequestBody = &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{Schema: bodyRef},
					},
				},
			}
		}
	}

	op.Responses = openapi3.NewResponses(openapi3.WithStatus(200, buildResponse(out)))

	return op
}

func buildResponse(out *core.EmiActionBody) *openapi3.ResponseRef {
	desc := "OK"
	resp := &openapi3.Response{
		Description: &desc,
	}

	if out != nil {
		if ref := bodySchemaRef(out); ref != nil {
			resp.Content = openapi3.Content{
				"application/json": &openapi3.MediaType{Schema: ref},
			}
		}
		if len(out.Headers) > 0 {
			resp.Headers = openapi3.Headers{}
			for _, h := range out.Headers {
				resp.Headers[h.Name] = &openapi3.HeaderRef{
					Value: &openapi3.Header{
						Parameter: openapi3.Parameter{
							Description: h.Description,
							Schema:      openapi3.NewSchemaRef("", headerSchema(h)),
						},
					},
				}
			}
		}
	}

	return &openapi3.ResponseRef{Value: resp}
}

// bodySchemaRef picks the right OpenAPI schema for an action body. Priority:
// referenced DTO → declared fields → primitive. Envelope is intentionally
// ignored for now — wrapping logic differs per generator and adding it here
// would require modelling each envelope by name.
func bodySchemaRef(body *core.EmiActionBody) *openapi3.SchemaRef {
	if body == nil {
		return nil
	}
	if body.Dto != "" {
		return openapi3.NewSchemaRef("#/components/schemas/"+body.Dto, nil)
	}
	if len(body.Fields) > 0 {
		return openapi3.NewSchemaRef("", buildObjectSchema(body.Fields, ""))
	}
	if body.Primitive != "" {
		return openapi3.NewSchemaRef("", primitiveSchema(body.Primitive))
	}
	return nil
}

// ----------------------------------------------------------------------
// Schemas from Emi fields
// ----------------------------------------------------------------------

func buildObjectSchema(fields []*core.EmiField, description string) *openapi3.Schema {
	schema := &openapi3.Schema{
		Type:        &openapi3.Types{openapi3.TypeObject},
		Description: description,
		Properties:  openapi3.Schemas{},
	}

	for _, f := range fields {
		if f == nil {
			continue
		}
		schema.Properties[f.Name] = fieldSchemaRef(f)
		if !isNullable(string(f.Type)) {
			schema.Required = append(schema.Required, f.Name)
		}
	}

	if len(schema.Properties) == 0 {
		schema.Properties = nil
	}
	return schema
}

func fieldSchemaRef(f *core.EmiField) *openapi3.SchemaRef {
	t := string(f.Type)
	base := strings.TrimSuffix(t, "?")
	nullable := isNullable(t)

	switch core.FieldType(base) {

	case core.FieldTypeString:
		return wrap(primitiveWithDesc("string", f.Description), nullable)

	case core.FieldTypeBool:
		return wrap(primitiveWithDesc("bool", f.Description), nullable)

	case core.FieldTypeInt, core.FieldTypeInt32, core.FieldTypeInt64:
		return wrap(primitiveWithDesc(base, f.Description), nullable)

	case core.FieldTypeFloat32, core.FieldTypeFloat64:
		return wrap(primitiveWithDesc(base, f.Description), nullable)

	case core.FieldTypeAny:
		s := &openapi3.Schema{Description: f.Description}
		return openapi3.NewSchemaRef("", s)

	case core.FieldTypeOne:
		if f.Target != "" {
			return openapi3.NewSchemaRef("#/components/schemas/"+f.Target, nil)
		}
		return wrap(&openapi3.Schema{Type: &openapi3.Types{openapi3.TypeObject}, Description: f.Description}, nullable)

	case core.FieldTypeCollection:
		items := &openapi3.SchemaRef{}
		if f.Target != "" {
			items = openapi3.NewSchemaRef("#/components/schemas/"+f.Target, nil)
		} else {
			items = openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeObject}})
		}
		return wrap(&openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeArray},
			Description: f.Description,
			Items:       items,
		}, nullable)

	case core.FieldTypeSlice:
		items := openapi3.NewSchemaRef("", primitiveSchema(firstNonEmpty(f.Primitive, "string")))
		return wrap(&openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeArray},
			Description: f.Description,
			Items:       items,
		}, nullable)

	case core.FieldTypeArray:
		// Emi `array` is an array of inline objects whose shape comes from
		// nested Fields. If no fields are given, fall back to an array of
		// generic objects.
		var itemsSchema *openapi3.Schema
		if len(f.Fields) > 0 {
			itemsSchema = buildObjectSchema(f.Fields, "")
		} else {
			itemsSchema = &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeObject}}
		}
		return wrap(&openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeArray},
			Description: f.Description,
			Items:       openapi3.NewSchemaRef("", itemsSchema),
		}, nullable)

	case core.FieldTypeObject:
		obj := buildObjectSchema(f.Fields, f.Description)
		return wrap(obj, nullable)

	case core.FieldTypeMap:
		pair := openapi3.NewSchemaRef("", primitiveSchema(firstNonEmpty(f.MapPairOf, "string")))
		return wrap(&openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeObject},
			Description: f.Description,
			AdditionalProperties: openapi3.AdditionalProperties{
				Schema: pair,
			},
		}, nullable)

	case core.FieldTypeEnum:
		values := make([]any, 0, len(f.OfType))
		for _, e := range f.OfType {
			if e == nil {
				continue
			}
			values = append(values, e.GetKey())
		}
		return wrap(&openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeString},
			Description: f.Description,
			Enum:        values,
		}, nullable)

	case core.FieldTypeComplex:
		if f.Complex != "" {
			return openapi3.NewSchemaRef("#/components/schemas/"+f.Complex, nil)
		}
		return openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeObject}})
	}

	// Unknown / unmapped — fall back to primitive resolution then a free-form object.
	if s := primitiveSchema(base); s != nil {
		return wrap(s, nullable)
	}
	return openapi3.NewSchemaRef("", &openapi3.Schema{Description: f.Description})
}

func queryFieldSchema(q *core.EmiQueryField) *openapi3.Schema {
	t := string(q.Type)
	switch core.FieldType(strings.TrimSuffix(t, "?")) {
	case core.FieldTypeSlice, core.FieldTypeArray:
		items := openapi3.NewSchemaRef("", primitiveSchema(firstNonEmpty(q.Primitive, "string")))
		return &openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeArray},
			Description: q.Description,
			Items:       items,
		}
	case core.FieldTypeObject:
		return buildObjectSchema(q.Fields, q.Description)
	}
	s := primitiveSchema(t)
	if s == nil {
		s = primitiveSchema("string")
	}
	s.Description = q.Description
	return s
}

func headerSchema(h core.EmiHeader) *openapi3.Schema {
	if h.Complex != "" {
		return &openapi3.Schema{
			Type:        &openapi3.Types{openapi3.TypeObject},
			Description: h.Description,
		}
	}
	s := primitiveSchema(h.Type)
	if s == nil {
		s = primitiveSchema("string")
	}
	s.Description = h.Description
	return s
}

// primitiveSchema maps Emi primitive type names to OpenAPI primitive schemas.
// Returns nil for unknown types so callers can decide on a fallback.
func primitiveSchema(t string) *openapi3.Schema {
	switch strings.TrimSuffix(t, "?") {
	case "string":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeString}}
	case "bool", "boolean":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeBoolean}}
	case "int":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeInteger}}
	case "int32":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeInteger}, Format: "int32"}
	case "int64":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeInteger}, Format: "int64"}
	case "float32":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeNumber}, Format: "float"}
	case "float64":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeNumber}, Format: "double"}
	case "bytes":
		return &openapi3.Schema{Type: &openapi3.Types{openapi3.TypeString}, Format: "byte"}
	case "any":
		return &openapi3.Schema{}
	}
	return nil
}

func primitiveWithDesc(t, desc string) *openapi3.Schema {
	s := primitiveSchema(t)
	if s == nil {
		s = &openapi3.Schema{}
	}
	s.Description = desc
	return s
}

func wrap(s *openapi3.Schema, nullable bool) *openapi3.SchemaRef {
	if s == nil {
		return openapi3.NewSchemaRef("", &openapi3.Schema{})
	}
	if nullable {
		s.Nullable = true
	}
	return openapi3.NewSchemaRef("", s)
}

// ----------------------------------------------------------------------
// URL handling
// ----------------------------------------------------------------------

type pathParam struct {
	Name string
	Type string
}

var emiPathParam = regexp.MustCompile(`/:([A-Za-z0-9_]+)(?:\s+([A-Za-z0-9_]+))?`)

// normalizeURL turns an Emi-style URL like `/stream/:id/:name string` into
// `/stream/{id}/{name}` and returns the extracted parameter names plus their
// declared types. Anything without a declared type defaults to `string`.
func normalizeURL(u string) (string, []pathParam) {
	var params []pathParam
	normalized := emiPathParam.ReplaceAllStringFunc(u, func(match string) string {
		m := emiPathParam.FindStringSubmatch(match)
		name := m[1]
		typ := m[2]
		if typ == "" {
			typ = "string"
		}
		params = append(params, pathParam{Name: name, Type: typ})
		return "/{" + name + "}"
	})
	return normalized, params
}

// methodOf normalizes the Emi method to an HTTP verb. Empty or non-HTTP
// methods on actions that still have a URL default to GET.
func methodOf(method, url string) string {
	switch strings.ToUpper(method) {
	case "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD":
		return strings.ToUpper(method)
	}
	if url != "" {
		return "GET"
	}
	return ""
}

// ----------------------------------------------------------------------
// Misc helpers
// ----------------------------------------------------------------------

func isNullable(t string) bool {
	return strings.HasSuffix(t, "?")
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// jsonToYAML round-trips a JSON document through a generic map so the YAML
// output mirrors the JSON exactly. kin-openapi's own YAML marshalling can
// produce subtly different shapes for some fields, and round-tripping via
// JSON avoids those divergences.
func jsonToYAML(jsonBytes []byte) ([]byte, error) {
	var v any
	if err := json.Unmarshal(jsonBytes, &v); err != nil {
		return nil, err
	}
	return yaml.Marshal(v)
}
