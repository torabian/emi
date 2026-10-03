package core

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// resolveInterfaces prepares every interface and includes it into the dtos and entities
// implementing it:
//
//  1. An interface's own fields are resolved first, in the interface's context (its
//     `use:` templates become concrete fields) and checked to be shareable: no inline
//     object/array children.
//  2. Each implementer gets the interface's fields in front of its own. A field the
//     implementer already declares is kept as its own when it has the same type, and is
//     an error when it doesn't - the interface is a contract.
//
// Runs after resolveFieldUses, so an implementer's own `use:` fields are concrete when
// compared. Safe to run more than once: a second pass finds every interface field
// already present with the same type and adds nothing.
func (m *Emi) resolveInterfaces() error {
	if len(m.Interfaces) == 0 {
		for _, body := range m.interfaceBodies() {
			if len(body.body.Implements) > 0 {
				return fmt.Errorf("%s implements %v, but the module declares no interfaces", body.owner, body.body.Implements)
			}
		}
		for _, dto := range m.Dto {
			if len(dto.Implements) > 0 {
				return fmt.Errorf("dto %q implements %v, but the module declares no interfaces", dto.Name, dto.Implements)
			}
		}
		for _, entity := range m.Entities {
			if entity != nil && len(entity.Implements) > 0 {
				return fmt.Errorf("entity %q implements %v, but the module declares no interfaces", entity.Name, entity.Implements)
			}
		}
		return nil
	}

	var templates []EmiFieldTemplate
	if m.Templates != nil {
		templates = m.Templates.Fields
	}

	seen := map[string]bool{}
	for i := range m.Interfaces {
		iface := &m.Interfaces[i]
		if iface.Name == "" {
			return fmt.Errorf("interface #%d has no name", i+1)
		}
		if seen[iface.Name] {
			return fmt.Errorf("interface %q is declared twice", iface.Name)
		}
		seen[iface.Name] = true
		if err := resolveInterfaceFields(iface, templates); err != nil {
			return err
		}
	}

	for i := range m.Dto {
		dto := &m.Dto[i] // m.Dto holds values: index it, or Fields would be edited on a copy
		if err := m.includeInterfaces("dto", dto.Name, fieldContextDto, dto.Implements, &dto.Fields); err != nil {
			return err
		}
	}
	for _, item := range m.interfaceBodies() {
		if len(item.body.Implements) > 0 && len(item.body.Fields) == 0 && item.body.Dto != "" {
			return fmt.Errorf("%s implements %v, but references dto %q: implements needs inline fields", item.owner, item.body.Implements, item.body.Dto)
		}
		if err := m.includeInterfaces(item.kind, item.name, fieldContextDto, item.body.Implements, &item.body.Fields); err != nil {
			return err
		}
	}
	for _, entity := range m.Entities {
		if entity == nil {
			continue
		}
		if err := m.includeInterfaces("entity", entity.Name, fieldContextEntity, entity.Implements, &entity.Fields); err != nil {
			return err
		}
	}
	return nil
}

func resolveInterfaceFields(iface *EmiInterface, templates []EmiFieldTemplate) error {
	switch iface.Context {
	case "", fieldContextDto, fieldContextEntity:
	default:
		return fmt.Errorf("interface %q: context must be dto or entity, got %q", iface.Name, iface.Context)
	}
	ctx := iface.Context
	if ctx == "" {
		ctx = fieldContextDto
	}

	if iface.templated == nil {
		iface.templated = map[string]bool{}
	}
	names := map[string]bool{}
	for _, f := range iface.Fields {
		if f == nil {
			return fmt.Errorf("interface %q has an empty field entry", iface.Name)
		}
		if f.Name == "" {
			return fmt.Errorf("interface %q has a field without a name", iface.Name)
		}
		if names[f.Name] {
			return fmt.Errorf("interface %q declares field %q twice", iface.Name, f.Name)
		}
		names[f.Name] = true

		if f.Use != "" {
			iface.templated[f.Name] = true
		}
		if err := resolveInterfaceField(iface, f, f.Name, ctx, templates); err != nil {
			return err
		}
	}
	return nil
}

// resolveInterfaceField resolves `use:` templates on a field and, recursively, on the
// children of an inline object/array field - those children become part of a shared
// type, so they have to be concrete. path is the dotted name, for messages.
func resolveInterfaceField(iface *EmiInterface, f *EmiField, path, ctx string, templates []EmiFieldTemplate) error {
	if f.Use != "" {
		if err := applyFieldTemplate(f, ctx, templates); err != nil {
			return fmt.Errorf("interface %q: %w", iface.Name, err)
		}
	}

	// Inline children are fine on object/array (the interface owns the nested type, and
	// every implementer's own nested type is an alias of it). Other types can't carry
	// them in a way an implementer could share.
	if len(f.Fields) > 0 && !IsInlineNestedType(f.Type) {
		return fmt.Errorf("interface %q field %q: %s fields with inline children can't be shared - use object or array, or declare a dto and reference it with target", iface.Name, path, f.Type)
	}

	names := map[string]bool{}
	for _, child := range f.Fields {
		if child == nil || child.Name == "" {
			return fmt.Errorf("interface %q field %q has a child without a name", iface.Name, path)
		}
		if names[child.Name] {
			return fmt.Errorf("interface %q field %q declares child %q twice", iface.Name, path, child.Name)
		}
		names[child.Name] = true
		if err := resolveInterfaceField(iface, child, path+"."+child.Name, ctx, templates); err != nil {
			return err
		}
	}
	return nil
}

// IsInlineNestedType reports whether a field of this type declares its own nested type
// through inline children (object/array), which generators render as a class owned by
// the parent.
func IsInlineNestedType(t FieldType) bool {
	return t == FieldTypeObject || t == FieldTypeObjectNullable || t == FieldTypeArray || t == FieldTypeArrayNullable
}

func (m *Emi) includeInterfaces(kind, owner, ownerContext string, implements []string, fields *[]*EmiField) error {
	if len(implements) == 0 {
		return nil
	}

	existing := map[string]*EmiField{}
	for _, f := range *fields {
		if f != nil {
			existing[f.Name] = f
		}
	}

	var included []*EmiField
	for _, name := range implements {
		iface := m.findInterface(name)
		if iface == nil {
			return fmt.Errorf("%s %q implements unknown interface %q (declared: %s)", kind, owner, name, m.interfaceNames())
		}
		ifaceContext := iface.Context
		if ifaceContext == "" {
			ifaceContext = fieldContextDto
		}

		for _, f := range iface.Fields {
			if kind == "entity" && IsInlineNestedType(f.Type) {
				return fmt.Errorf("entity %q implements interface %q, but its field %q is an inline %s: an entity stores those in child tables of its own, so they can't be shared through an interface - declare a dto and reference it with target",
					owner, name, f.Name, f.Type)
			}

			if iface.templated[f.Name] && ifaceContext != ownerContext {
				return fmt.Errorf("%s %q implements interface %q, but its field %q is built from a template resolved for %s context - declare the interface with context: %s (or use a template-free field)",
					kind, owner, name, f.Name, ifaceContext, ownerContext)
			}

			if own, ok := existing[f.Name]; ok {
				if fieldShape(own) != fieldShape(f) {
					return fmt.Errorf("%s %q declares field %q as %s, but interface %q requires %s",
						kind, owner, f.Name, fieldShape(own), name, fieldShape(f))
				}
				continue
			}

			copied, err := copyField(f)
			if err != nil {
				return err
			}
			existing[f.Name] = copied
			included = append(included, copied)
		}
	}

	if len(included) > 0 {
		*fields = append(included, *fields...)
	}
	return nil
}

func (m *Emi) findInterface(name string) *EmiInterface {
	for i := range m.Interfaces {
		if m.Interfaces[i].Name == name {
			return &m.Interfaces[i]
		}
	}
	return nil
}

func (m *Emi) interfaceNames() string {
	names := make([]string, 0, len(m.Interfaces))
	for _, i := range m.Interfaces {
		names = append(names, i.Name)
	}
	return strings.Join(names, ", ")
}

// fieldShape is the part of a field that decides its generated type, including the
// children of an inline object/array.
func fieldShape(f *EmiField) string {
	shape := fmt.Sprintf("%s(target=%q complex=%q primitive=%q map=%q/%q)", f.Type, f.Target, f.Complex, f.Primitive, f.MapKeyOf, f.MapPairOf)
	if len(f.Fields) > 0 {
		children := make([]string, 0, len(f.Fields))
		for _, child := range f.Fields {
			if child != nil {
				children = append(children, child.Name+":"+fieldShape(child))
			}
		}
		shape += "{" + strings.Join(children, ",") + "}"
	}
	return shape
}

// copyField deep-copies a field so implementers never share slices or maps with the
// interface or with each other.
func copyField(f *EmiField) (*EmiField, error) {
	raw, err := yaml.Marshal(f)
	if err != nil {
		return nil, err
	}
	var out EmiField
	if err := yaml.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// interfacesForDerivedDto are the interfaces of an entity that the plain dto derived from
// it (BuildEntityDto) can also implement. That dto is built from the entity's fields, but
// relations are rewritten to their nullable dto counterparts, so an interface that has a
// relation (or any nested/list type) would need different types there - the dto would not
// satisfy it, and the generated assertion would not compile. Such an interface stays on
// the entity alone. The optional dto never qualifies: it makes every field nullable.
func (m *Emi) interfacesForDerivedDto(entity *Module3Entity) []string {
	var names []string
	for _, name := range entity.Implements {
		iface := m.findInterface(name)
		if iface == nil || !interfaceHasOnlyPlainFields(iface) {
			continue
		}
		// template-built fields only fit owners of the interface's own context - the dto
		// would be rejected if preprocessing ran again.
		if iface.Context == fieldContextEntity && len(iface.templated) > 0 {
			continue
		}
		names = append(names, name)
	}
	return names
}

func interfaceHasOnlyPlainFields(iface *EmiInterface) bool {
	for _, f := range iface.Fields {
		if f == nil {
			continue
		}
		switch f.Type {
		case FieldTypeOne, FieldTypeOneNullable, FieldTypeCollection, FieldTypeCollectionNullable,
			FieldTypeArray, FieldTypeArrayNullable, FieldTypeObject, FieldTypeObjectNullable,
			FieldTypeList, FieldTypeListNullable, FieldTypeClass, FieldTypeClassNullable,
			FieldTypeMap, FieldTypeMapNullable, FieldTypeSlice, FieldTypeSliceNullable:
			return false
		}
	}
	return true
}

// interfaceBody is an event/permission body that may declare `implements`.
type interfaceBody struct {
	kind  string // "event" or "permission", for messages
	name  string
	owner string // e.g. `event "postPublished" params`
	body  *EmiActionBody
}

// interfaceBodies collects every event params/payload and permission params body of the
// module, the places besides dtos and entities where `implements` is honored.
func (m *Emi) interfaceBodies() []interfaceBody {
	var items []interfaceBody
	for _, e := range m.Events {
		if e == nil {
			continue
		}
		if e.Params != nil {
			items = append(items, interfaceBody{"event", e.Key + " params", fmt.Sprintf("event %q params", e.Key), e.Params})
		}
		if e.Payload != nil {
			items = append(items, interfaceBody{"event", e.Key + " payload", fmt.Sprintf("event %q payload", e.Key), e.Payload})
		}
	}
	for _, p := range FlattenPermissions(m.Permissions) {
		if p.Params != nil {
			items = append(items, interfaceBody{"permission", p.FullKey + " params", fmt.Sprintf("permission %q params", p.FullKey), p.Params})
		}
	}
	return items
}
