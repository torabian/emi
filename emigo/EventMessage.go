package emigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

// A message is a Go text/template (text/template, not html/template: it is a sentence, not a
// page). Its data is the event's payload - every field of it, under the name it has in the
// definition ({{.firstName}}, not {{.FirstName}}) - plus whatever params the firing call adds.
// Params win over payload fields of the same name. A message without "{{" is used as it is.
//
//	message:
//	  en: "{{.inviter}} invited {{.firstName}} to join {{.workspace}}."
//
// Quote such a message in yaml: a plain scalar can not start with "{".
const messageTemplateMarker = "{{"

// messageFuncs are the helpers available inside a message.
var messageFuncs = template.FuncMap{
	// default returns fallback when value is empty: {{default "someone" .inviter}}
	"default": func(fallback string, value any) string {
		if s := strings.TrimSpace(fmt.Sprint(value)); s != "" && s != "<nil>" {
			return s
		}
		return fallback
	},
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
}

// MessageData is what a message template sees: the payload's fields, then the params on top.
// A payload that is not an object (nil, a string, ...) contributes nothing.
func (x Event) MessageData(params map[string]any) map[string]any {
	data := map[string]any{}

	if x.Payload != nil {
		if raw, err := json.Marshal(x.Payload); err == nil {
			var fromPayload map[string]any
			if json.Unmarshal(raw, &fromPayload) == nil {
				for k, v := range fromPayload {
					data[k] = v
				}
			}
		}
	}
	for k, v := range params {
		data[k] = v
	}

	return data
}

// RenderMessage returns the event's Message with every locale's template filled in from the
// payload and params (see MessageData). It never fails the event: a locale whose template can
// not be read or run keeps its text as written, and the problems come back together in the
// error, for the caller to log. A key the data does not have renders as nothing - use the
// default function for a fallback ({{default "someone" .inviter}}).
func (x Event) RenderMessage(params map[string]any) (map[string]string, error) {
	if len(x.Message) == 0 {
		return x.Message, nil
	}

	var data map[string]any
	out := make(map[string]string, len(x.Message))
	var problems []error

	locales := make([]string, 0, len(x.Message))
	for locale := range x.Message {
		locales = append(locales, locale)
	}
	sort.Strings(locales)

	for _, locale := range locales {
		text := x.Message[locale]
		out[locale] = text
		if !strings.Contains(text, messageTemplateMarker) {
			continue
		}

		if data == nil {
			data = x.MessageData(params)
		}

		t, err := template.New(x.Key + "." + locale).Funcs(messageFuncs).Option("missingkey=zero").Parse(text)
		if err != nil {
			problems = append(problems, fmt.Errorf("message %q: %w", locale, err))
			continue
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			problems = append(problems, fmt.Errorf("message %q: %w", locale, err))
			continue
		}

		// A key missing from the data prints as "<no value>" - nothing is the right text.
		out[locale] = strings.ReplaceAll(buf.String(), "<no value>", "")
	}

	return out, errors.Join(problems...)
}
