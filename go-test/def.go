package main

import emi "test/goruntime"

// The base class definition for singleDto

type SingleDto struct {

	// Minimum number which can be generated
	Min int `json:"min" yaml:"min"`

	// Maximum number which can be generated
	Max int `json:"max" yaml:"max"`

	// How many numbers you want to be generated based on maximum and minimum
	Count int `json:"count" yaml:"count"`

	// This object is optional. Can be passed or not. On typ has ? operator
	NullableObject emi.Nullable[SingleDtoNullableObject] `json:"nullableObject" yaml:"nullableObject"`

	StaticArray []SingleDtoStaticArray `json:"staticArray" yaml:"staticArray"`

	NullableArray emi.Nullable[[]SingleDtoNullableArray] `json:"nullableArray" yaml:"nullableArray"`

	// This object is not nullable. On classes, will be always instantiated, also on types, is not having ? operator
	StaticObject SingleDtoStaticObject `json:"staticObject" yaml:"staticObject"`
}

// The base class definition for nullableObject

type SingleDtoNullableObject struct {
	FirstName string `json:"firstName" yaml:"firstName"`
}

// The base class definition for staticArray

type SingleDtoStaticArray struct {
	Item12 string `json:"item12" yaml:"item12"`
}

// The base class definition for nullableArray

type SingleDtoNullableArray struct {
	Item5 string `json:"item5" yaml:"item5"`
}

// The base class definition for staticObject

type SingleDtoStaticObject struct {
	FirstName string `json:"firstName" yaml:"firstName"`
}
