package golang

import "github.com/torabian/emi/lib/core"

type Flags struct {
	Emigo       string `json:"emigo,omitempty"`
	PackageName string `json:"pkg,omitempty"`
}

func GetCommonFlags(ctx core.MicroGenContext) Flags {
	var f Flags = Flags{
		Emigo:       "github.com/torabian/emi/emigo",
		PackageName: DEFAULT_GO_PACKAGE,
	}

	if val, ok := ctx.Flags["emigo"]; ok && val != "" {
		f.Emigo = val
	}

	if val, ok := ctx.Flags["pkg"]; ok && val != "" {
		f.PackageName = val
	}

	return f
}
