package core

// addDefaultEntityPermissions appends a standard permission group for every entity with
// DefaultPermissions: true, unless the module already declares a root permission of the
// same name (hand-declared always wins, which also keeps repeated Preprocess calls
// idempotent). Shape: <entity> (<entity>.*) with query/create/update/delete children.
func (m *Emi) addDefaultEntityPermissions() {
	existing := make(map[string]bool, len(m.Permissions))
	for _, p := range m.Permissions {
		if p != nil {
			existing[p.Name] = true
		}
	}

	for _, e := range m.Entities {
		if e == nil || !e.DefaultPermissions || e.Name == "" || existing[e.Name] {
			continue
		}
		existing[e.Name] = true

		child := func(op, verb string) *EmiPermission {
			return &EmiPermission{
				Name:        op,
				Key:         op,
				FullKey:     e.Name + "." + op,
				Title:       map[string]string{"en": verb + " " + e.Name},
				Description: map[string]string{"en": verb + " " + e.Name + " records."},
			}
		}

		m.Permissions = append(m.Permissions, &EmiPermission{
			Name:        e.Name,
			Key:         e.Name,
			FullKey:     e.Name + ".*",
			Title:       map[string]string{"en": "Entire " + e.Name + " actions (*)"},
			Description: map[string]string{"en": "Grants every " + e.Name + "-related permission."},
			Children: []*EmiPermission{
				child("query", "Query"),
				child("create", "Create"),
				child("update", "Update"),
				child("delete", "Delete"),
			},
		})
	}
}
