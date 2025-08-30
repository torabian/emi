package js

import (
	"fmt"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// In this file we are going to put small functions which can generate code for different languages

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
	callback func(action *core.Module3Action, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error),
) (string, error) {

	action, err := core.StringToModule3Action(ctx.Content)
	if err != nil {
		return "", err
	}

	result, err := callback(&action, ctx)
	if err != nil {
		return "", err
	}

	return AsFullDocument(result), nil
}

func commonJsObjectStringCompiler(
	ctx core.MicroGenContext,
	callback func(fields []*core.Module3Field, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error),
) (string, error) {

	fields, err := core.StringToModule3Fields(ctx.Content)
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
	callback func(module *core.Module3, ctx core.MicroGenContext) ([]core.VirtualFile, error),
) ([]core.VirtualFile, error) {

	action, err := core.StringToModule3(ctx.Content)
	if err != nil {
		return nil, err
	}

	result, err := callback(&action, ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
