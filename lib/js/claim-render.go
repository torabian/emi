package js

import (
	"github.com/torabian/emi/lib/core"
)

func ClaimRender(claims []core.JsFnArgument, ctx core.MicroGenContext) map[string]string {
	rendered := make(map[string]string)
	for _, c := range claims {
		if ctx.HasTag(Typescript) {
			rendered[c.Key] = c.CompileTs()
		} else {
			rendered[c.Key] = c.CompileJs()
		}
	}
	return rendered
}
