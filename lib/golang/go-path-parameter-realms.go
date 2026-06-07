package golang

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

type GoPathParameterRealms struct {
	Params       []GoPathParamer
	ClassName    string
	TypeName     string
	Dependencies []core.CodeChunkDependency
	SplitCli     bool
}

type GoPathParamer struct {
	PlaceHolderValue string
	GolangFieldName  string
	GolangType       string
	CliName          string
}

func GoActionPathParamsRealms(action core.EmiRpcAction, ctx core.MicroGenContext) (*GoPathParameterRealms, error) {

	res := GoPathParameterRealms{
		Dependencies: []core.CodeChunkDependency{},
	}

	res.SplitCli = ctx.HasTag(SplitCli)
	placeholders0 := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders0) == 0 {
		return nil, nil // nothing to generate
	}

	placeholders := []GoPathParamer{}
	for _, item := range placeholders0 {

		golangType := goPrimitiveDetect(item.Type)
		// If the type is not string, str conv is necessary to happen
		if golangType != "string" {
			res.Dependencies = append(res.Dependencies, core.CodeChunkDependency{
				Location: "strconv",
			})
		}

		placeholders = append(placeholders, GoPathParamer{
			PlaceHolderValue: item.Original,
			GolangFieldName:  core.ToUpper(item.Original),
			GolangType:       golangType,
			CliName:          FieldToCliName(item.Original),
		})
	}

	className := action.GetName()
	typeName := fmt.Sprintf("%vPathParameter", className)
	res.ClassName = className
	res.TypeName = typeName
	res.Params = placeholders

	return &res, nil
}
