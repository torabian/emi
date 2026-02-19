package external

func CastCommonVectorResponseDtoFromCli() CommonVectorResponseDto {
	data := CommonVectorResponseDto{}
	return data
}

// The base class definition for commonVectorResponseDto
type CommonVectorResponseDto struct {
	OutputVector []int `json:"outputVector" yaml:"outputVector"`
}
