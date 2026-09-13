package core

import "fmt"

// preprocessIntents resolves EmiIntent definitions before any compiler backend runs:
//
//  1. Every EmiAction with Intent: true gets a synthetic EmiIntent{Name: a.Name, From:
//     a.Name} appended to m.Intents, unless a hand-declared intent of that name already
//     exists (hand-declared always wins - same yield convention preprocessEntityActions
//     uses for entity actions vs hand-declared ones).
//  2. Every intent with From set is resolved against the module's actions: Description/
//     In/Out are copied from the referenced action wherever the intent didn't already
//     declare its own (an intent's own fields always take precedence).
//  3. Every intent's final In/Out is checked: if it references a Dto by name (rather
//     than declaring Fields inline), that Dto must actually be declared in this
//     module's top-level Dto list - a compiler backend has no field list to fall back
//     on for a Dto reference, so a typo or a Dto that only exists in Templates (which
//     are never compiled) would otherwise surface as a raw "undefined type" error deep
//     inside generated code instead of a clear message here.
//
// Runs unconditionally as part of Preprocess() - unlike PreprocessEntityActions, this
// needs no backend opt-in, since it only ever adds to m.Intents/fills gaps on entries
// already there, never mutates m.Actions.
func (m *Emi) preprocessIntents() error {
	actionByName := make(map[string]*EmiAction, len(m.Actions))
	for _, a := range m.Actions {
		if a != nil && a.Name != "" {
			actionByName[a.Name] = a
		}
	}

	existing := make(map[string]bool, len(m.Intents))
	for _, it := range m.Intents {
		if it != nil && it.Name != "" {
			existing[it.Name] = true
		}
	}

	for _, a := range m.Actions {
		if a == nil || !a.Intent || a.Name == "" || existing[a.Name] {
			continue
		}
		it := &EmiIntent{Name: a.Name, From: a.Name}
		m.Intents = append(m.Intents, it)
		existing[a.Name] = true
	}

	for _, it := range m.Intents {
		if it == nil || it.From == "" {
			continue
		}

		src, ok := actionByName[it.From]
		if !ok {
			return fmt.Errorf("intent %q: from action %q not found", it.Name, it.From)
		}

		if it.Name == "" {
			it.Name = src.Name
		}
		if it.Description == "" {
			it.Description = src.Description
		}
		if it.In == nil {
			it.In = src.In
		}
		if it.Out == nil {
			it.Out = src.Out
		}
	}

	dtoClassNames := make(map[string]bool, len(m.Dto))
	for i := range m.Dto {
		dtoClassNames[m.Dto[i].GetClassName()] = true
	}

	for _, it := range m.Intents {
		if it == nil {
			continue
		}
		if err := validateIntentBodyDto(it.Name, "in", it.In, dtoClassNames); err != nil {
			return err
		}
		if err := validateIntentBodyDto(it.Name, "out", it.Out, dtoClassNames); err != nil {
			return err
		}
	}

	return nil
}

// validateIntentBodyDto checks that body's Dto reference (if any) names a Dto actually
// declared in this module's top-level Dto list - see preprocessIntents step 3.
func validateIntentBodyDto(intentName, side string, body *EmiActionBody, dtoClassNames map[string]bool) error {
	if body == nil || body.Dto == "" {
		return nil
	}
	if !dtoClassNames[body.Dto] {
		return fmt.Errorf("intent %q: %s dto %q is not declared in this module's dtos", intentName, side, body.Dto)
	}
	return nil
}
