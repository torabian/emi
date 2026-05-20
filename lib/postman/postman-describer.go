package postman

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// PostmanModuleDescriber reads an Emi module and produces a Postman v2.1
// collection as a single VirtualFile. Mirrors the openapi describer in
// shape: small helpers per section so each piece can be reused.
func PostmanModuleDescriber(
	ctx core.MicroGenContext,
	m core.Emi,
) ([]core.VirtualFile, error) {

	host := firstNonEmpty(ctx.Flags["host"], "localhost")
	port := firstNonEmpty(ctx.Flags["port"], "4500")

	collection := BuildPostmanCollection(m, host, port)

	name := m.Name
	if name == "" {
		name = "postman"
	}

	return []core.VirtualFile{
		{
			Name:         name,
			MimeType:     "application/json",
			Extension:    ".postman_collection.json",
			ActualScript: collection.Json(),
		},
	}, nil
}

// BuildPostmanCollection converts an Emi module into a Postman collection.
func BuildPostmanCollection(m core.Emi, host, port string) *PostmanCollection {

	collection := &PostmanCollection{
		Info: PostmanInfo{
			Name:        firstNonEmpty(m.Name, "Emi Module"),
			Description: m.Description,
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Auth: &PostmanAuth{
			Type: "apikey",
			ApiKey: []PostmanAuthApiKey{
				{Key: "value", Value: "{{AUTH}}", Type: "string"},
				{Key: "key", Value: "authorization", Type: "string"},
			},
		},
		Variable: []PostmanVariable{
			{Key: "HOST", Value: host, Type: "default"},
			{Key: "WID", Value: "root", Type: "default", Description: "Current active workspace id"},
			{Key: "AUTH", Value: "", Type: "default", Description: "Authorization header to authenticate"},
			{Key: "RID", Value: "", Type: "default", Description: "Selected role which user is working on that behalf"},
			{Key: "PORT", Value: port, Type: "default"},
		},
	}

	dtoIndex := indexDtos(m.Dto)

	for _, action := range m.Actions {
		if action == nil || strings.TrimSpace(action.Url) == "" {
			continue
		}
		collection.Item = append(collection.Item, buildActionItem(action, dtoIndex))
	}

	for _, remote := range m.Remotes {
		if remote == nil || strings.TrimSpace(remote.Url) == "" {
			continue
		}
		collection.Item = append(collection.Item, buildRemoteItem(remote, dtoIndex))
	}

	return collection
}

// ----------------------------------------------------------------------
// Item builders
// ----------------------------------------------------------------------

func buildActionItem(a *core.EmiAction, dtos map[string]*core.EmiDto) PostmanItem {
	method := strings.ToUpper(a.Method)
	if method == "" {
		method = "GET"
	}

	url := buildUrl(a.Url, a.Query)

	item := PostmanItem{
		Name: firstNonEmpty(a.Name, a.Url),
		Request: PostmanRequest{
			Method:      method,
			Header:      defaultHeaders(actionHeaders(a)),
			Url:         url,
			Description: a.Description,
		},
	}

	if a.HasRequest() {
		item.Request.Body = buildBody(a.In, dtos)
	}

	return item
}

func buildRemoteItem(r *core.EmiRemote, dtos map[string]*core.EmiDto) PostmanItem {
	method := strings.ToUpper(r.Method)
	if method == "" {
		method = "GET"
	}

	url := buildUrl(r.Url, r.Query)

	item := PostmanItem{
		Name: firstNonEmpty(r.Name, r.Url),
		Request: PostmanRequest{
			Method:      method,
			Header:      defaultHeaders(remoteHeaders(r)),
			Url:         url,
			Description: r.Name,
		},
	}

	if r.In != nil {
		item.Request.Body = buildBody(r.In, dtos)
	}

	return item
}

func actionHeaders(a *core.EmiAction) []PostmanHeader {
	if !a.HasRequestHeaders() {
		return nil
	}
	out := make([]PostmanHeader, 0, len(a.In.Headers))
	for _, h := range a.In.Headers {
		out = append(out, PostmanHeader{Key: h.Name, Value: ""})
	}
	return out
}

func remoteHeaders(r *core.EmiRemote) []PostmanHeader {
	if r.In == nil || len(r.In.Headers) == 0 {
		return nil
	}
	out := make([]PostmanHeader, 0, len(r.In.Headers))
	for _, h := range r.In.Headers {
		out = append(out, PostmanHeader{Key: h.Name, Value: ""})
	}
	return out
}

// defaultHeaders prepends the standard Fireback auth/workspace/role headers
// to any action-specific ones. Action-specific headers come second so a
// user can override them in Postman without losing the standard set.
func defaultHeaders(extra []PostmanHeader) []PostmanHeader {
	base := []PostmanHeader{
		{Key: "authorization", Value: "{{AUTH}}"},
		{Key: "role-id", Value: "{{RID}}"},
		{Key: "workspace-id", Value: "{{WID}}"},
	}
	return append(base, extra...)
}

// ----------------------------------------------------------------------
// URL handling
// ----------------------------------------------------------------------

var emiPathParam = regexp.MustCompile(`/:([A-Za-z0-9_]+)(?:\s+[A-Za-z0-9_]+)?`)

// buildUrl turns an Emi URL like `/stream/:id string` into a Postman URL
// with `:id` placeholders preserved (Postman renders them as variables).
func buildUrl(u string, query []*core.EmiQueryField) PostmanUrl {
	normalized := emiPathParam.ReplaceAllStringFunc(u, func(match string) string {
		m := emiPathParam.FindStringSubmatch(match)
		return "/:" + m[1]
	})

	path := []string{}
	if normalized != "" {
		// Drop leading slash, then split. Empty segments are dropped so
		// trailing slashes don't add empty path elements.
		trimmed := strings.TrimPrefix(normalized, "/")
		for _, seg := range strings.Split(trimmed, "/") {
			if seg != "" {
				path = append(path, seg)
			}
		}
	}

	postmanQuery := []PostmanQuery{}
	queryParts := []string{}
	for _, q := range query {
		if q == nil {
			continue
		}
		postmanQuery = append(postmanQuery, PostmanQuery{
			Key:         q.Name,
			Description: q.Description,
		})
		queryParts = append(queryParts, q.Name+"=")
	}

	raw := "http://{{HOST}}:{{PORT}}" + normalized
	if len(queryParts) > 0 {
		raw = raw + "?" + strings.Join(queryParts, "&")
	}

	url := PostmanUrl{
		Raw:      raw,
		Protocol: "http",
		Host:     []string{"{{HOST}}"},
		Port:     "{{PORT}}",
		Path:     path,
	}
	if len(postmanQuery) > 0 {
		url.Query = postmanQuery
	}
	return url
}

// ----------------------------------------------------------------------
// Body / request example
// ----------------------------------------------------------------------

func buildBody(body *core.EmiActionBody, dtos map[string]*core.EmiDto) PostmanBody {
	example := requestExample(body, dtos)
	return PostmanBody{
		Mode: "raw",
		Raw:  example,
		Options: PostmanBodyOption{
			Raw: PostmanBodyOptionRaw{Language: "json"},
		},
	}
}

// requestExample produces a JSON example body for an action's input. It
// resolves DTOs by name from the module index, walks declared fields, or
// falls back to a primitive placeholder.
func requestExample(body *core.EmiActionBody, dtos map[string]*core.EmiDto) string {
	if body == nil {
		return ""
	}
	value := bodyExampleValue(body, dtos)
	if value == nil {
		return ""
	}
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return ""
	}
	return string(out)
}

func bodyExampleValue(body *core.EmiActionBody, dtos map[string]*core.EmiDto) any {
	if body == nil {
		return nil
	}
	if body.Dto != "" {
		if dto, ok := dtos[body.Dto]; ok {
			return fieldsExample(dto.Fields, dtos, map[string]bool{body.Dto: true})
		}
		return map[string]any{}
	}
	if len(body.Fields) > 0 {
		return fieldsExample(body.Fields, dtos, map[string]bool{})
	}
	if body.Primitive != "" {
		return primitiveExample(body.Primitive)
	}
	return nil
}

func fieldsExample(fields []*core.EmiField, dtos map[string]*core.EmiDto, seen map[string]bool) map[string]any {
	out := map[string]any{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		out[f.Name] = fieldExample(f, dtos, seen)
	}
	return out
}

func fieldExample(f *core.EmiField, dtos map[string]*core.EmiDto, seen map[string]bool) any {
	base := strings.TrimSuffix(string(f.Type), "?")
	switch core.FieldType(base) {

	case core.FieldTypeString:
		return ""
	case core.FieldTypeBool:
		return false
	case core.FieldTypeInt, core.FieldTypeInt32, core.FieldTypeInt64:
		return 0
	case core.FieldTypeFloat32, core.FieldTypeFloat64:
		return 0
	case core.FieldTypeAny:
		return nil

	case core.FieldTypeOne:
		return dtoExample(f.Target, dtos, seen)

	case core.FieldTypeCollection:
		return []any{dtoExample(f.Target, dtos, seen)}

	case core.FieldTypeSlice:
		return []any{primitiveExample(firstNonEmpty(f.Primitive, "string"))}

	case core.FieldTypeArray:
		if len(f.Fields) > 0 {
			return []any{fieldsExample(f.Fields, dtos, seen)}
		}
		return []any{map[string]any{}}

	case core.FieldTypeObject:
		return fieldsExample(f.Fields, dtos, seen)

	case core.FieldTypeMap:
		return map[string]any{}

	case core.FieldTypeEnum:
		for _, e := range f.OfType {
			if e != nil {
				return e.GetKey()
			}
		}
		return ""

	case core.FieldTypeComplex:
		return map[string]any{}
	}

	if v := primitiveExample(base); v != nil {
		return v
	}
	return nil
}

func dtoExample(name string, dtos map[string]*core.EmiDto, seen map[string]bool) any {
	if name == "" {
		return map[string]any{}
	}
	if seen[name] {
		// Recursive DTO reference — break the cycle with an empty object.
		return map[string]any{}
	}
	dto, ok := dtos[name]
	if !ok {
		return map[string]any{}
	}
	next := map[string]bool{}
	for k, v := range seen {
		next[k] = v
	}
	next[name] = true
	return fieldsExample(dto.Fields, dtos, next)
}

func primitiveExample(t string) any {
	switch strings.TrimSuffix(t, "?") {
	case "string":
		return ""
	case "bool", "boolean":
		return false
	case "int", "int32", "int64":
		return 0
	case "float32", "float64":
		return 0
	case "bytes":
		return ""
	case "any":
		return nil
	}
	return nil
}

// ----------------------------------------------------------------------
// Misc helpers
// ----------------------------------------------------------------------

func indexDtos(items []core.EmiDto) map[string]*core.EmiDto {
	out := map[string]*core.EmiDto{}
	for i := range items {
		out[items[i].Name] = &items[i]
	}
	return out
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
