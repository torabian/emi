package core

import "encoding/json"

// Macros is a pre-compile mechanism in Emi, and it will modify the module definition
// before it's given to the compiler. The idea is for example, you can add extra entities
// on some modules with it.
// Until version 1.2.1, there is a single macro for EAV database model, which would create
// All of the necessary tables and fields.
// Custom macros can be indefintely useful, but need to be very well defined and documented since
// the parameters are interface{}
type EmiMacro struct {

	// The macro name which you are using. Emi developers need to add the macros name here as reference.
	Using string `yaml:"using,omitempty" json:"using,omitempty" jsonschema:"enum=eav,description=The macro name which you are using. Emi developers need to add the macros name here as reference."`

	// Params are the macro configuration which are dynamically set based on each macro itself.
	// They will be passed as interface{} to macro and function itself will decide what to do next.
	Params interface{} `yaml:"params,omitempty" json:"params,omitempty" jsonschema:"description=Params are the macro configuration which are dynamically set based on each macro itself. They will be passed as interface{} to macro and function itself will decide what to do next."`
}

func ConvertParams(params interface{}) interface{} {
	if params == nil {
		return nil
	}

	// Handle map[interface{}]interface{} case
	if rawMap, ok := params.(map[interface{}]interface{}); ok {
		converted := make(map[string]interface{})
		for k, v := range rawMap {
			if key, isString := k.(string); isString {
				converted[key] = v
			}
		}
		return converted
	}
	return params
}

// Useful for calling when writing a custom macro.
func EmiMacroCastParams[T any](m *EmiMacro) (*T, error) {
	m.Params = ConvertParams(m.Params)

	data, err := json.Marshal(m.Params)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
