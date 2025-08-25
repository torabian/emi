package js

import "github.com/torabian/emi/lib/core"

func getAngularServiceImports() []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{}

	deps = append(deps, core.CodeChunkDependency{
		Objects: []string{
			"Injectable",
		},
		Location: "@angular/core",
	})

	deps = append(deps, core.CodeChunkDependency{
		Objects: []string{
			"HttpClient",
			"HttpHeaders",
			"HttpParams",
		},
		Location: "@angular/common/http",
	})

	deps = append(deps, core.CodeChunkDependency{
		Objects: []string{
			"Observable",
		},
		Location: "rxjs",
	})

	return deps
}
