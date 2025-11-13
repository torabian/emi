package unknownpackage
  // The base class definition for computeDto
type ComputeDto struct {
		  // Minimum number which can be generated
 Min interface{} `json:"min" yaml:"min"`
		  // Maximum number which can be generated
 Max int `json:"max" yaml:"max"`
		  // How many numbers you want to be generated based on maximum and minimum
 Count int `json:"count" yaml:"count"`
}