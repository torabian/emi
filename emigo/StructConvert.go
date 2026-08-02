package emigo

import "encoding/json"

// ConvertViaJSON converts src into a T by round-tripping through JSON: marshal src,
// then unmarshal the result into a zero-valued T. This is the simplest way to turn an
// entity (with its gorm tags, Row-sibling fields, internal Id, etc.) into a plain
// response dto - or an external system's payload into a dto - without hand-writing a
// field-by-field copy: it only requires the two types' `json` tags to agree on the
// fields that matter. Fields only one side has are silently dropped (src) or left at
// their zero value (T).
//
// This isn't zero-cost - it allocates and reflects twice - so avoid it in a hot loop.
// But for a one-off Create/Update response shape, it's simpler and safer than
// maintaining a second, hand-mapped conversion function per entity, especially since
// the data usually either came from or is headed to an external system as JSON anyway.
func ConvertViaJSON[T any](src any) (T, error) {
	var out T
	if src == nil {
		return out, nil
	}
	buf, err := json.Marshal(src)
	if err != nil {
		return out, err
	}
	if err := json.Unmarshal(buf, &out); err != nil {
		return out, err
	}
	return out, nil
}
