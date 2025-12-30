package external
	import emigo "github.com/torabian/emi/examples/fullstack/emigo"
  // The base class definition for commonVectorComputeDto
type CommonVectorComputeDto struct {
		InitialVector1 []int `json:"initialVector1" yaml:"initialVector1"`
		Value emigo.Nullable[string] `json:"value" yaml:"value"`
		InitialVector2 []int `json:"initialVector2" yaml:"initialVector2"`
}