package mathlib

import (
	"fmt"
	"strconv"
	"strings"
)

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) ToString() string {
	return fmt.Sprintf("%f,%f,%f", v.X, v.Y, v.Z)
}

func (v *Vector3) FromString(s string) error {
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid vector string: %s", s)
	}

	var err error
	v.X, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return err
	}
	v.Y, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return err
	}
	v.Z, err = strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return err
	}

	return nil
}
