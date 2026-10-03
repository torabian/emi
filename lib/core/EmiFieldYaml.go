package core

import (
	"fmt"
	"sort"
)

// A field's `description:` and `label:` may each be a plain string or a map of locale -> text. The Go side
// keeps two fields (see EmiField.Description / EmiField.Descriptions, and Label / Labels) so every code
// generator that only knows a string keeps working; the yaml side is handled here.
//
// EmiField's yaml tag on Description and Label is "-" - the (un)marshalers below read and
// write those keys by hand. EmiColumn and EmiFieldTemplate embed EmiField, which would promote its
// (un)marshalers over their own extra keys, so each of them defines its own.

// emiFieldPlain has EmiField's shape without its methods.
type emiFieldPlain EmiField

// emiFieldYAML is what an EmiField is written as: its plain keys plus the two keys that
// may be a string or a locale map, `description` and `label`.
type emiFieldYAML struct {
	emiFieldPlain `yaml:",inline"`
	Description   interface{} `yaml:"description,omitempty"`
	Label         interface{} `yaml:"label,omitempty"`
}

func (x EmiField) yamlForm() emiFieldYAML {
	form := emiFieldYAML{emiFieldPlain: emiFieldPlain(x)}
	form.Description = localizedTextForYAML(x.Description, x.Descriptions)
	form.Label = localizedTextForYAML(x.Label, x.Labels)
	return form
}

func (x EmiField) MarshalYAML() (interface{}, error) {
	return x.yamlForm(), nil
}

func (x *EmiField) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var plain emiFieldPlain
	if err := unmarshal(&plain); err != nil {
		return err
	}
	var raw struct {
		Description interface{} `yaml:"description"`
		Label       interface{} `yaml:"label"`
	}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	*x = EmiField(plain)
	var err error
	if x.Description, x.Descriptions, err = parseLocalizedText(raw.Description); err != nil {
		return fmt.Errorf("field %q: description: %w", plain.Name, err)
	}
	if x.Label, x.Labels, err = parseLocalizedText(raw.Label); err != nil {
		return fmt.Errorf("field %q: label: %w", plain.Name, err)
	}
	return nil
}

func (x EmiColumn) MarshalYAML() (interface{}, error) {
	return struct {
		emiFieldYAML `yaml:",inline"`
		Column       string `yaml:"column,omitempty"`
		Selected     bool   `yaml:"selected,omitempty"`
	}{x.EmiField.yamlForm(), x.Column, x.Selected}, nil
}

func (x *EmiColumn) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var field EmiField
	if err := unmarshal(&field); err != nil {
		return err
	}
	var own struct {
		Column   string `yaml:"column"`
		Selected bool   `yaml:"selected"`
	}
	if err := unmarshal(&own); err != nil {
		return err
	}
	*x = EmiColumn{EmiField: field, Column: own.Column, Selected: own.Selected}
	return nil
}

func (x EmiFieldTemplate) MarshalYAML() (interface{}, error) {
	return struct {
		emiFieldYAML `yaml:",inline"`
		Context      string `yaml:"context,omitempty"`
	}{x.EmiField.yamlForm(), x.Context}, nil
}

func (x *EmiFieldTemplate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var field EmiField
	if err := unmarshal(&field); err != nil {
		return err
	}
	var own struct {
		Context string `yaml:"context"`
	}
	if err := unmarshal(&own); err != nil {
		return err
	}
	*x = EmiFieldTemplate{EmiField: field, Context: own.Context}
	return nil
}

// localizedTextForYAML picks what to write for a string-or-map key: the map when there is
// one, else the plain string, else nothing (omitempty drops the nil).
func localizedTextForYAML(text string, texts map[string]string) interface{} {
	if len(texts) > 0 {
		return texts
	}
	if text != "" {
		return text
	}
	return nil
}

// parseLocalizedText turns a yaml scalar or mapping into the (string, map) pair
// EmiField carries: a scalar gives just the string; a mapping gives the locale map plus
// its default text (the "en" entry, else the first locale alphabetically).
func parseLocalizedText(raw interface{}) (string, map[string]string, error) {
	switch v := raw.(type) {
	case nil:
		return "", nil, nil
	case string:
		return v, nil, nil
	case map[interface{}]interface{}:
		texts := make(map[string]string, len(v))
		for k, val := range v {
			text, ok := val.(string)
			if !ok {
				return "", nil, fmt.Errorf("locale %v must map to a string, got %T", k, val)
			}
			texts[fmt.Sprint(k)] = text
		}
		if len(texts) == 0 {
			return "", nil, nil
		}
		return defaultLocalizedText(texts), texts, nil
	case map[string]interface{}:
		converted := make(map[interface{}]interface{}, len(v))
		for k, val := range v {
			converted[k] = val
		}
		return parseLocalizedText(converted)
	default:
		// Numbers, booleans... - same as a plain string field did before.
		return fmt.Sprint(v), nil, nil
	}
}

func defaultLocalizedText(texts map[string]string) string {
	if en, ok := texts["en"]; ok {
		return en
	}
	locales := make([]string, 0, len(texts))
	for l := range texts {
		locales = append(locales, l)
	}
	sort.Strings(locales)
	return texts[locales[0]]
}
