package js

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// In this file we are going to put small functions which can generate code for different languages

type PackageJSON struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Description  string            `json:"description"`
	Main         string            `json:"main"`
	Scripts      map[string]string `json:"scripts"`
	Dependencies map[string]string `json:"dependencies"`
}

func GeneratePackageJSON(packageName string, locations []string) ([]byte, error) {
	// remove duplicates
	locMap := map[string]struct{}{}
	for _, loc := range locations {
		locMap[loc] = struct{}{}
	}

	// Used versions
	var packagesUsedVersions = map[string]string{
		"@tanstack/react-query": "^5.85.5",
		"axios":                 "^1.11.0",
		"qs":                    "^6.14.0",
		"@types/qs":             "^6.14.0",
	}

	var sortedLocs []string
	for loc := range locMap {
		sortedLocs = append(sortedLocs, loc)
	}
	sort.Strings(sortedLocs)

	// build dependencies
	deps := map[string]string{}
	for _, loc := range sortedLocs {

		// This is internal file import, so do not add
		// to package
		if strings.HasPrefix(loc, "./") {
			continue
		}
		if version, ok := packagesUsedVersions[loc]; ok {
			deps[loc] = version
		} else {
			deps[loc] = "latest"
		}
	}

	pkg := PackageJSON{
		Name:        packageName,
		Version:     "1.0.0",
		Description: "Auto-generated package.json",
		Main:        "index.js",
		Scripts: map[string]string{
			"build": "echo 'build script here'",
			"test":  "echo 'test script here'",
		},
		Dependencies: deps,
	}

	return json.MarshalIndent(pkg, "", "  ")
}

func CombineImportsJsWorld(chunk core.CodeChunkCompiled) string {
	// group by location
	locMap := map[string]map[string]struct{}{}

	for _, dep := range chunk.CodeChunkDependenies {
		if _, ok := locMap[dep.Location]; !ok {
			locMap[dep.Location] = map[string]struct{}{}
		}
		for _, obj := range dep.Objects {
			locMap[dep.Location][obj] = struct{}{}
		}
	}

	// build final import statements
	var importsList []string
	for loc, objs := range locMap {
		// sort objects for deterministic output
		objSlice := make([]string, 0, len(objs))
		for obj := range objs {
			objSlice = append(objSlice, obj)
		}
		sort.Strings(objSlice)
		statement := fmt.Sprintf("import { %s } from '%s';", strings.Join(objSlice, ", "), loc)
		importsList = append(importsList, statement)
	}

	// sort imports by location for consistency
	sort.Strings(importsList)

	// combine with actual script
	return strings.Join(importsList, "\r\n")
}

func commonJsActionStringCompiler(
	ctx core.MicroGenContext,
	callback func(action core.EmiRpcAction, ctx core.MicroGenContext, complexes []RecognizedComplex) (*core.CodeChunkCompiled, error),
) (string, error) {

	action, err := core.StringToEmiAction(ctx.Content)
	if err != nil {
		return "", err
	}

	result, err := callback(&action, ctx, []RecognizedComplex{})
	if err != nil {
		return "", err
	}

	return AsFullDocument(result), nil
}

func commonJsObjectStringCompiler(
	ctx core.MicroGenContext,
	callback func(fields []*core.EmiField, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error),
) (string, error) {

	fields, err := core.StringToEmiFields(ctx.Content)
	if err != nil {
		return "", err
	}

	// In this case, the only flag is the virtual class name which will be passed
	result, err := callback(fields, ctx, JsCommonObjectContext{RootClassName: ctx.Flags})
	if err != nil {
		return "", err
	}

	return AsFullDocument(result), nil
}

func commonJsDtoStringCompiler(
	ctx core.MicroGenContext,
	callback func(dto core.EmiDto, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error),
) (string, error) {

	fields, err := core.StringToEmiDto(ctx.Content)
	if err != nil {
		return "", err
	}

	// In this case, the only flag is the virtual class name which will be passed
	result, err := callback(fields, ctx, JsCommonObjectContext{RootClassName: ctx.Flags})
	if err != nil {
		return "", err
	}

	return AsFullDocument(result), nil
}

func commonJsModuleFileCompiler(
	ctx core.MicroGenContext,
	callback func(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error),
) ([]core.VirtualFile, error) {

	action, err := core.StringToEmi(ctx.Content)
	if err != nil {
		return nil, err
	}

	result, err := callback(&action, ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
