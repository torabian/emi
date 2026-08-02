package external

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Money is a hand-written complex type used by entity.emi.yml's "complex1" field. It
// lives in the same package as the generated entities, so no import/location is needed
// in the complexes: entry.
//
// It implements both encoding.Text(Un)Marshaler (used by the generated CLI cast code)
// and driver.Valuer/sql.Scanner (so gorm can persist it as a plain integer column - a
// real complex type is expected to implement whatever interfaces its own storage needs,
// the same way fireback's own complex types, e.g. XDate, do).
type Money struct {
	Cents int64
}

func (m Money) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", m.Cents)), nil
}

func (m *Money) UnmarshalText(text []byte) error {
	cents, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	m.Cents = cents
	return nil
}

func (m Money) Value() (driver.Value, error) {
	return m.Cents, nil
}

func (m *Money) Scan(value interface{}) error {
	if value == nil {
		m.Cents = 0
		return nil
	}

	switch v := value.(type) {
	case int64:
		m.Cents = v
	case []byte:
		cents, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		m.Cents = cents
	default:
		return fmt.Errorf("unsupported Scan type for Money: %T", value)
	}

	return nil
}
