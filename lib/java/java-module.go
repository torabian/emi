package java

import (
	"slices"
	"strings"

	"github.com/torabian/emi/lib/core"
	javaInclude "github.com/torabian/emi/lib/java/java-include"
)

// javaPackage is the single package every generated type lives in. Since
// every generated file compiles into the same package, Java needs no
// per-file import wiring between them at all (unlike lib/dart or
// lib/python) - any type is reachable from any other file in the same
// package for free.
const javaPackage = "emisdk"

// commonImports is emitted at the top of every generated file. Unused
// imports are not even a compiler warning in Java (unlike C#'s mild
// analyzer suggestion), so a single fixed set - covering every bare
// `List`/`Map`/Jackson-annotation reference the templates in this package
// emit - is simpler and just as correct as computing a precise one per file.
var commonImports = []string{
	"java.util.List",
	"java.util.Map",
	"com.fasterxml.jackson.annotation.JsonCreator",
	"com.fasterxml.jackson.annotation.JsonProperty",
	"com.fasterxml.jackson.annotation.JsonValue",
}

// AsFullDocument wraps a compiled chunk into a standalone .java file: the
// package declaration, the common imports, plus any module-specific import
// (e.g. a resolved complex type's package), then the actual body.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	imports := make([]string, len(commonImports))
	copy(imports, commonImports)

	seen := map[string]struct{}{}
	for _, i := range imports {
		seen[i] = struct{}{}
	}
	for _, dep := range x.CodeChunkDependensies {
		if dep.Location == "" {
			continue
		}
		if _, ok := seen[dep.Location]; ok {
			continue
		}
		seen[dep.Location] = struct{}{}
		imports = append(imports, dep.Location)
	}

	var sb strings.Builder
	sb.WriteString("package " + javaPackage + ";\n\n")
	for _, i := range imports {
		sb.WriteString("import " + i + ";\n")
	}
	sb.WriteString("\n")
	sb.Write(x.ActualScript)

	return string(core.EscapeLines([]byte(sb.String())))
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/java target, mirroring the same helper in the sibling generators.
// ImportLocation here is a fully-qualified Java type or package to import,
// not a file path.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "java" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type JavaModuleGenerationFlags struct {
	Actions *string
	Remotes *string
	Dtos    *string
}

func (x JavaModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readJavaModuleFlags(ctx core.MicroGenContext) JavaModuleGenerationFlags {
	config := JavaModuleGenerationFlags{}
	if v, ok := ctx.Flags["actions"]; ok && v != "" {
		config.Actions = &v
	}
	if v, ok := ctx.Flags["remotes"]; ok && v != "" {
		config.Remotes = &v
	}
	if v, ok := ctx.Flags["dtos"]; ok && v != "" {
		config.Dtos = &v
	}
	return config
}

// chunksToVirtualFiles converts each compiled chunk (already one public
// Java type each) into a VirtualFile under srcDir, running it through
// AsFullDocument for the package declaration + imports.
func chunksToVirtualFiles(chunks []*core.CodeChunkCompiled, srcDir string) []core.VirtualFile {
	files := make([]core.VirtualFile, 0, len(chunks))
	for _, chunk := range chunks {
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     srcDir,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}
	return files
}

// JavaModuleFullVirtualFiles compiles an entire Emi module into a Java
// source tree: one file per dto/enum/action/remote (under
// src/main/java/emisdk/, matching the `emisdk` package), plus the embedded
// runtime (Fetchx) and a pom.xml, mirroring the equivalent full-module
// compilers in the sibling generators.
func JavaModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readJavaModuleFlags(ctx)
	files := []core.VirtualFile{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	const srcDir = "src/main/java/emisdk"

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}
		chunks, err := JavaCommonObjectGenerator(dto.Fields, ctx, JavaCommonObjectContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		files = append(files, chunksToVirtualFiles(chunks, srcDir)...)
	}

	for _, enum := range module.Enums {
		chunk, err := JavaStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     srcDir,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, action := range module.Actions {
		if len(actionFilter) > 0 && !slices.Contains(actionFilter, action.Name) {
			continue
		}
		chunks, err := JavaActionRender(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		files = append(files, chunksToVirtualFiles(chunks, srcDir)...)
	}

	for _, remote := range module.Remotes {
		if len(remoteFilter) > 0 && !slices.Contains(remoteFilter, remote.Name) {
			continue
		}
		chunks, err := JavaActionRender(remote, ctx, complexes)
		if err != nil {
			return nil, err
		}
		files = append(files, chunksToVirtualFiles(chunks, srcDir)...)
	}

	skipSdk := ctx.HasTag(NoSdk)
	if !skipSdk {
		files = append(files, core.FsEmbedToVirtualFile(&javaInclude.Content, srcDir)...)
	}

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "pom",
			Extension:    ".xml",
			ActualScript: javaPom,
		})
	}

	return files, nil
}

const javaPom = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated by the Emi compiler - do not edit by hand. -->
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.emisdk</groupId>
  <artifactId>emi-sdk</artifactId>
  <version>1.0.0</version>
  <packaging>jar</packaging>

  <properties>
    <maven.compiler.source>17</maven.compiler.source>
    <maven.compiler.target>17</maven.compiler.target>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>

  <dependencies>
    <dependency>
      <groupId>com.fasterxml.jackson.core</groupId>
      <artifactId>jackson-databind</artifactId>
      <version>2.17.0</version>
    </dependency>
  </dependencies>
</project>
`
