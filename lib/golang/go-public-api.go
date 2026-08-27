package golang

import (
	"errors"
	"fmt"
	"go/format"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/torabian/emi/lib/core"
)

var registerEntityPreprocessHookOnce sync.Once

func GetGolangPublicActions() core.PublicAPIActions {
	// Entities' optional-dto/plain-dto/EmiAction synthesis (core.PreprocessEntity*)
	// lives in lib/core, but core never registers any of it itself - only the golang
	// backend actually consumes entities today, so it's the one that opts in here.
	// Because RegisterPreprocessHook is global, this makes all of it available to every
	// other backend too (openapi, postman, ...) as soon as the golang module is enabled
	// in a given binary, with zero code changes required in any of them.
	// PreprocessEntityActions runs last (registration order is execution order - see
	// runPreprocessHooks) since its Create/Update/Browse actions reference the
	// plain/optional dtos' class names. Guarded since GetGolangPublicActions is called
	// more than once per process (see gorunner.go, cmd/emi-wasm/main.go).
	registerEntityPreprocessHookOnce.Do(func() {
		core.RegisterPreprocessHook(core.PreprocessEntityOptionalDtos)
		core.RegisterPreprocessHook(core.PreprocessEntityDtos)
		core.RegisterPreprocessHook(core.PreprocessEntityActions)
	})

	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "go:dto",
				Description:      "Generates dto, for golang, both as client and server",
				WasmFunctionName: "goGenObject",
				Flags: []core.FlagDef{
					{
						Name:     "pkg",
						Type:     core.FlagString,
						Usage:    "Package name of the golang",
						Required: true,
					},
					{
						Name:  "emi-runtime",
						Type:  core.FlagString,
						Usage: "Location of the emi runtime",
					},
				},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}

				emiLocation := ""
				if val, ok := ctx.Flags["emi-runtime"]; ok && val != "" && val != "<nil>" {
					emiLocation = val
				}

				res, err := GoCommonStructGenerator(emiDto.Fields, ctx, GoCommonStructContext{RootClassName: emiDto.Name, EmiLocation: emiLocation})
				if err != nil {
					return "", err
				}

				return AsFullDocument(res.MainClass, ctx.Flags["pkg"]), nil
			},
		},
	}

	fileActions := []core.ActionFile{
		GoPrimaryAction,
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}

var GoPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "go",
		Description:      "Compiles golang from .emi catalog spec file",
		WasmFunctionName: "goGen",
		Flags: []core.FlagDef{
			{
				Name:     "emigo",
				Usage:    "Add location to emigo path folder, can be also github.com/torabian/emi/emigo if you wanted to",
				Required: false,
				Type:     core.FlagString,
				Default:  "github.com/torabian/emi/emigo",
			},
			{
				Name:  "pkg",
				Type:  core.FlagString,
				Usage: "Package name of the golang",
			},
		},
	},
	Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
		type_, err := core.DetectEmiStringContentType(ctx.Content)
		if err != nil {
			return nil, err
		}

		if type_ == "module" {
			emiModule, err := core.StringToEmi(ctx.Content)
			if err != nil {
				return nil, err
			}

			files, err := GoModuleFull(&emiModule, ctx)
			if err != nil {
				return nil, err
			}

			return files, err
		}

		return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
	},
}

// Finds the ts/js compatible types.
// Make sure this function is public on later versions
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {

		// only pick general or js/ts specific complexes for js-modules
		if complex.Compiler == "go" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
				Namespace:      complex.Namespace,
			})
		}
	}

	return items
}

type GoModuleGenerationFlags struct {
	Dtos *string `json:"dtos"`
}

func (x GoModuleGenerationFlags) GetDtos() []string {
	return strings.Split(*x.Dtos, ",")
}

// Combines entire features for a module, and creates a virtual map of the files
// which is necessary to run entire modules
func GoModuleFull(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {

	f := GetCommonFlags(ctx)

	complexes := DiscoverComplexes(module)
	files := []core.VirtualFile{}

	config := GoModuleGenerationFlags{}
	if ctx.Flags["dtos"] != "" {
		str := ctx.Flags["dtos"]
		config.Dtos = &str
	}

	// core.Emi.Preprocess synthesizes each entity's update dto directly into
	// module.Dto (see lib/core/preprocess-entities.go) - by the time Go codegen runs,
	// it's an ordinary dto like any other, so it's rendered here through the same
	// generic pipeline as every hand-declared dto, landing in its own
	// {Entity}EntityUpdateDto.go file.
	for _, dto := range module.Dto {
		if dto.Name == "" {
			continue
		}

		if config.Dtos != nil && len(*config.Dtos) > 0 && !slices.Contains(config.GetDtos(), dto.Name) {
			continue
		}

		actionRendered, err := GoCommonStructGenerator(dto.Fields, ctx, GoCommonStructContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
			EmiLocation:         f.Emigo,
		})

		if err != nil {
			return nil, err
		}

		dtoItem := actionRendered.MainClass

		files = append(files, core.VirtualFile{
			Name:         dtoItem.SuggestedFileName,
			Extension:    dtoItem.SuggestedExtension,
			ActualScript: AsFullDocument(dtoItem, f.PackageName),
		})

		if actionRendered.CliHelpers != nil {
			files = append(files, core.VirtualFile{
				Name:         actionRendered.CliHelpers.SuggestedFileName,
				Extension:    actionRendered.CliHelpers.SuggestedExtension,
				ActualScript: AsFullDocument(actionRendered.CliHelpers, f.PackageName),
			})
		}
	}

	for _, entity := range module.Entities {
		if entity == nil || entity.Name == "" {
			continue
		}

		// id/uniqueId go on first, exactly once, so every renderer below sees them as
		// ordinary declared fields with one consistent shape.
		PrependEntityDefaultFields(entity)

		// IMPORTANT: the actions codegen has to run BEFORE GoEntityRender does.
		// GoEntityRender calls ApplyEntityGormTags, which mutates entity.Fields (and
		// the individual *EmiField values) in place - appending Row-sibling fields,
		// injecting LinkerId/id/uniqueId into array children. None of that belongs in
		// the actions codegen, which needs the fields as originally declared. The
		// update dto itself doesn't have this hazard - core.Emi.Preprocess already
		// built it, long before Go-specific mutation ever happens, and it's rendered
		// separately by the generic per-dto loop above.
		actionsRendered, err := GoEntityActionsRender(entity, ctx)
		if err != nil {
			return nil, err
		}

		entityRendered, err := GoEntityRender(entity, ctx, complexes)
		if err != nil {
			return nil, err
		}

		// One file per entity - the struct and its Create/Update actions are facets of
		// the same thing, splitting them across separate files just makes them harder
		// to find. The update dto is a plain, portable dto by this point though (see
		// the generic per-dto loop above), so it gets its own {Entity}EntityUpdateDto.go
		// like any other dto instead of being folded in here.
		//
		// entityRendered.CliHelpers is the one exception: when the "split-cli" tag is
		// set, GoCommonStructGenerator (via GoEntityRender) already rendered it as its
		// own chunk whose script leads with a `//go:build !wasm` line (see
		// GoCommonStructGeneratorCli). Folding that chunk into combined here (behind
		// MainClass's own script) doesn't just leave the tag inert - go/format (called
		// by AsFullDocument's FormatGoCode) hoists *any* `//go:build` comment it finds
		// anywhere in a file to the very top, regardless of where it originally sat
		// (verified empirically with `gofmt` directly). So a merged combined file would
		// end up with `//go:build !wasm` gating the *entire* entity - struct,
		// CreateFn/UpdateFn/GetFn/BrowseFn/AwareDelete, all of it - instead of just the
		// CLI helpers it was meant for. It has to be emitted as its own VirtualFile
		// instead, exactly like the generic per-dto loop above does for
		// actionRendered.CliHelpers.
		combined := &core.CodeChunkCompiled{
			SuggestedFileName:  entityRendered.MainClass.SuggestedFileName,
			SuggestedExtension: entityRendered.MainClass.SuggestedExtension,
		}
		appendChunk := func(c *core.CodeChunkCompiled) {
			if c == nil {
				return
			}
			combined.ActualScript = append(combined.ActualScript, c.ActualScript...)
			combined.CodeChunkDependensies = append(combined.CodeChunkDependensies, c.CodeChunkDependensies...)
		}

		appendChunk(entityRendered.MainClass)
		appendChunk(actionsRendered)

		files = append(files, core.VirtualFile{
			Name:         combined.SuggestedFileName,
			Extension:    combined.SuggestedExtension,
			ActualScript: AsFullDocument(combined, f.PackageName),
		})

		if entityRendered.CliHelpers != nil {
			files = append(files, core.VirtualFile{
				Name:         entityRendered.CliHelpers.SuggestedFileName,
				Extension:    entityRendered.CliHelpers.SuggestedExtension,
				ActualScript: AsFullDocument(entityRendered.CliHelpers, f.PackageName),
			})
		}
	}

	for _, action := range module.Actions {

		outputs, err := GoActionRender(action, ctx, complexes)

		if err != nil {
			return nil, err
		}

		for _, output := range outputs {
			files = append(files, core.VirtualFile{
				Name:         output.SuggestedFileName,
				Extension:    output.SuggestedExtension,
				ActualScript: AsFullDocument(output, f.PackageName),
			})
		}
	}

	vsqlFiles, err := GoVsqlsGenerate(module, ctx, complexes, f.Emigo, f.PackageName)
	if err != nil {
		return nil, err
	}
	files = append(files, vsqlFiles...)

	for _, manifest := range module.Manifests {
		gomanifest, err := GoManifest(manifest, module, ctx)
		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         gomanifest.SuggestedFileName,
			Extension:    gomanifest.SuggestedExtension,
			ActualScript: AsFullDocument(gomanifest, f.PackageName),
		})
	}

	configOutputs, err := GoConfigGenerate(module.Config, ctx)

	if err != nil {
		return nil, err
	}

	for _, output := range configOutputs {
		files = append(files, core.VirtualFile{
			Name:         output.SuggestedFileName,
			Extension:    output.SuggestedExtension,
			ActualScript: AsFullDocument(output, f.PackageName),
		})
	}

	permissionsOutput, err := GoPermissionsGenerate(module.Permissions, ctx, f.Emigo)
	if err != nil {
		return nil, err
	}

	if permissionsOutput != nil {
		files = append(files, core.VirtualFile{
			Name:         permissionsOutput.SuggestedFileName,
			Extension:    permissionsOutput.SuggestedExtension,
			ActualScript: AsFullDocument(permissionsOutput, f.PackageName),
		})
	}

	return files, nil
}

func FormatGoCode(code string) string {
	src := []byte(code)
	formatted, err := format.Source(src)
	if err != nil {
		// if formatting fails, just return original
		return code
	}
	return string(formatted)
}

func AsFullDocument(x *core.CodeChunkCompiled, packageName string) string {

	if item := core.FindTokenByName(x.Tokens, "PACKAGE_NAME"); item != nil {
		packageName = item.Value
	}

	// A "split" chunk (SplitCli/SplitGin/reactive-wasm, ...) emits its `//go:build` line
	// as the first line of its own ActualScript, for readability at the template level.
	// It has to be pulled out here rather than left in place: go/build only recognizes a
	// build constraint that appears before the `package` clause, separated from it by a
	// blank line - anywhere else (e.g. right after `package`, which is where it'd land
	// if left inline below) it's just an inert comment, and the file compiles
	// unconditionally in every GOOS/GOARCH/tag combination, silently defeating the
	// whole point of splitting it out. (Verified empirically: `go build`/`go vet` both
	// ignore a `//go:build` comment placed after the package clause.)
	buildTag, script := extractLeadingBuildTag(string(x.ActualScript))

	importsList := CombineGoImports(*x)
	var finalContent string = "package " + packageName + "\r\n" + importsList + "\r\n" + script

	finalContent = FormatGoCode(string(core.EscapeLines([]byte(finalContent))))

	if buildTag != "" {
		// Prepended after EscapeLines/FormatGoCode, never before: EscapeLines strips
		// every blank line it finds, which would merge the constraint straight into
		// `package` and break the very rule we're fixing.
		finalContent = buildTag + "\n\n" + finalContent
	}

	return finalContent
}

// extractLeadingBuildTag pulls a leading `//go:build ...` line off the front of a
// generated chunk's script (tolerating leading blank lines/whitespace, since templates
// emit it via `{{ if .SplitX }}\n//go:build !wasm\n{{ end }}`), returning it separately
// from the remaining script. Returns ("", script) unchanged if there is none.
func extractLeadingBuildTag(script string) (string, string) {
	trimmed := strings.TrimLeft(script, "\r\n\t ")
	if !strings.HasPrefix(trimmed, "//go:build") {
		return "", script
	}

	nl := strings.IndexAny(trimmed, "\r\n")
	if nl == -1 {
		return trimmed, ""
	}

	return trimmed[:nl], trimmed[nl:]
}
func CombineGoImports(chunk core.CodeChunkCompiled) string {
	statements := map[string]struct{}{}

	// Collect unique import statements
	for _, dep := range chunk.CodeChunkDependensies {
		statement := ""
		if len(dep.Objects) > 0 {
			statement = fmt.Sprintf(`%v "%v"`, dep.Objects[0], dep.Location)
		} else {
			statement = fmt.Sprintf(`"%v"`, dep.Location)
		}
		statements[statement] = struct{}{}
	}

	// Sort statements for deterministic output
	var sorted []string
	for stmt := range statements {
		sorted = append(sorted, stmt)
	}
	sort.Strings(sorted)

	// Combine into final import block
	if len(sorted) == 0 {
		return ""
	} else if len(sorted) == 1 {
		return "import " + sorted[0]
	} else {
		return "import (\n" + strings.Join(sorted, "\n") + "\n)"
	}
}
