package external
import "encoding/json"
	import emigo "github.com/torabian/emi/examples/fullstack/emigo"
func GetCommonVectorResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "output-vector",
			Type: "slice",
		},
		{
			Name: prefix + "request",
			Type: "one",
		},
	}
}
func CastCommonVectorResponseDtoFromCli(c emigo.CliCastable) CommonVectorResponseDto {
	data := CommonVectorResponseDto{}
			if c.IsSet("output-vector") { 
 emigo.InflatePossibleSlice(c.String("output-vector"), &data.OutputVector) 
}
	return data
}
  // The base class definition for commonVectorResponseDto
type CommonVectorResponseDto struct {
		OutputVector []int `json:"outputVector" yaml:"outputVector"`
		Request CommonVectorComputeDto `json:"request" yaml:"request"`
}
func (x *CommonVectorResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}