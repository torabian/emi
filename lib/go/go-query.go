package gogen

import "github.com/torabian/emi/lib/core"

func GoGenQueries(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {

	return []core.VirtualFile{
		{
			Name:         "model",
			MimeType:     ".go",
			Location:     "module",
			ActualScript: "asd",
			Extension:    ".go",
		},
	}, nil
}
