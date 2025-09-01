package core

func (x *Emi) PublicName() string {
	return ToUpper(x.Name)
}

func (x *EmiField) PublicName() string {
	return ToUpper(x.Name)
}

func (x *EmiFieldMatch) PublicName() string {
	if x.Dto == nil {
		return ""
	}

	return ToUpper(*x.Dto) + "Dto"
}

func (x *EmiField) PrivateName() string {
	return x.Name
}

func (x *EmiAction) Upper() string {
	return ToUpper(x.Name)
}
