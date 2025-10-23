package core

type EmiFieldMatch struct {

	// The dto name from Emi which will be matched. Might be also work with any other go struct but check the generated code.
	Dto *string `yaml:"dto,omitempty" json:"dto,omitempty" jsonschema:"description=The dto name from Emi which will be matched. Might be also work with any other go struct but check the generated code."`
}

func (x *EmiFieldMatch) PublicName() string {
	if x.Dto == nil {
		return ""
	}

	return ToUpper(*x.Dto) + "Dto"
}
