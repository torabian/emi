package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func AppendEavCustomParams(schema *jsonschema.Schema) {

	using := orderedmap.New[string, *jsonschema.Schema]()
	using.Set("using", &jsonschema.Schema{
		Const: "eav",
	})

	for key := range schema.Definitions {
		// if key == "EmiMacro" {

		// 	jsonStr := `{
		// 		"oneOf": [
		// 			{
		// 				"if": {
		// 					"properties": {
		// 						"using": {
		// 							"const": "eav"
		// 						}
		// 					}
		// 				},
		// 				"then": {
		// 					"properties": {
		// 						"params": {
		// 							"$ref": "#/definitions/EavMacroParams"
		// 						}
		// 					}
		// 				}
		// 			}
		// 		]
		// 	}`

		// 	var schemaP jsonschema.Schema
		// 	err := json.Unmarshal([]byte(jsonStr), &schemaP)
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		return
		// 	}

		// 	schema.Definitions[key].OneOf = schemaP.OneOf
		// }
		if key == "EmiConfigField" {

			jsonStr := `{
				"anyOf": [
					{
					"if": {
						"properties": {
						"type": {
							"const": "bool"
						}
						}
					},
					"then": {
						"properties": {
						"default": {
							"type": "boolean"
						}
						}
					}
					},
					{
					"if": {
						"properties": {
						"type": {
							"const": "string"
						}
						}
					},
					"then": {
						"properties": {
						"default": {
							"type": "string"
						}
						}
					}
					}
				]
			}`

			var schemaP jsonschema.Schema
			err := json.Unmarshal([]byte(jsonStr), &schemaP)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			var schemaEmpty jsonschema.Schema
			json.Unmarshal([]byte("{}"), &schemaEmpty)

			schema.Definitions[key].Properties.Set("default", &schemaEmpty)
			schema.Definitions[key].AnyOf = schemaP.AnyOf
		}
	}

	// // add the missing definitions
	// reflector := jsonschema.Reflector{}
	// schema2 := reflector.Reflect(&EavMacroParams{})

	// for key := range schema2.Definitions {
	// 	schema.Definitions[key] = schema2.Definitions[key]
	// }

}

// When the type is module, we create the spec only for that.
func genModuleSpec() string {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&Emi{})

	AppendEavCustomParams(schema)

	// Convert the schema to JSON
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}

	toWrite := string(schemaJSON)
	toWrite = strings.ReplaceAll(toWrite, "$defs", "definitions")
	toWrite = strings.ReplaceAll(toWrite, "https://json-schema.org/draft/2020-12/schema", "http://json-schema.org/draft-07/schema#")

	return toWrite
}

// On vscode settings, there is a redhat yaml plugin, if you provide the schemas
// it would provide the autocomplete.
func vsCodeSchemaMapping() string {
	return `
	{
		"yaml.schemas": {
			"./.vscode/emi-module-spec.json": [
				"*.emi.yml",
				"*.emi.yaml",
			],
			"./.vscode/emi-dto-spec.json": [
				"*.dto.yml",
				"*.dto.yaml",
			]
		},
	}
	
	`
}

func genDtoSpec() string {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&EmiDto{})

	AppendEavCustomParams(schema)

	// Convert the schema to JSON
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}

	toWrite := string(schemaJSON)
	toWrite = strings.ReplaceAll(toWrite, "$defs", "definitions")
	toWrite = strings.ReplaceAll(toWrite, "https://json-schema.org/draft/2020-12/schema", "http://json-schema.org/draft-07/schema#")

	return toWrite
}

func GenerateJsonSpecForEmi() ([]VirtualFile, error) {

	return []VirtualFile{
		{
			Name:         "emi-module-spec",
			MimeType:     "text/json",
			Extension:    ".json",
			ActualScript: genModuleSpec(),
		},
		{
			Name:         "emi-dto-spec",
			MimeType:     "text/json",
			Extension:    ".json",
			ActualScript: genDtoSpec(),
		},
		{
			Name:         "vscode.settings",
			MimeType:     "text",
			Extension:    ".txt",
			ActualScript: vsCodeSchemaMapping(),
		},
	}, nil
}
