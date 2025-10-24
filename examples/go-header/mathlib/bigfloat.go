package mathlib

import (
	"fmt"
	"math/big"
)

// BigFloat wraps *big.Float for headers / string conversion
type BigFloat struct {
	*big.Float
}

// NewBigFloat creates a new BigFloat
func NewBigFloat(f float64) BigFloat {
	return BigFloat{big.NewFloat(f)}
}

// ToString converts the BigFloat to a string
func (b BigFloat) ToString() string {
	if b.Float == nil {
		return "0"
	}
	return b.Text('f', -1) // -1 = exact precision
}

// FromString parses a string into BigFloat
func (b *BigFloat) FromString(s string) error {
	if b.Float == nil {
		b.Float = big.NewFloat(0)
	}
	_, ok := b.Float.SetString(s)
	if !ok {
		return fmt.Errorf("invalid BigFloat string: %s", s)
	}
	return nil
}
