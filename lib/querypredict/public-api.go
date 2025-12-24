package querypredict

import (
	"github.com/torabian/emi/lib/core"
)

func GetQPPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{}

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "qp:sql",
				Description:      "Compiles a single query into golang executable code, and transforms meta data into pure sql",
				WasmFunctionName: "sqlQueryPredict",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {

				return ProcessQueryPredicts(QueryDocument{Queries: []QuerySpec{QuerySpec{
					Name:  "query.sql",
					Query: ctx.Content,
				}}})

			},
		},
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}
