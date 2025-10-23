package core

import (
	"encoding/json"
	"strings"

	"github.com/invopop/jsonschema"
)

func AppendEavCustomParams(schema *jsonschema.Schema) {
	// Here you can modify the schema generated, append some stuff to it.
	// Look in the fireback the same function name for an example.
	// In emi, we don't need it, hopefully.
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
