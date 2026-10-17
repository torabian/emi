package c

import (
	"fmt"
	"slices"
	"strings"

	cInclude "github.com/torabian/emi/lib/c/c-include"
	"github.com/torabian/emi/lib/core"
)

// commonIncludes is emitted at the top of every generated header. An
// unused #include is never a compiler error in C, so a single fixed set -
// covering every standard header any of this package's templates might
// need - is simpler and just as correct as computing a precise one per
// file.
var commonIncludes = []string{
	`"cjson/cJSON.h"`,
	"<stdbool.h>",
	"<stddef.h>",
	"<stdint.h>",
	"<stdio.h>",
	"<stdlib.h>",
	"<string.h>",
}

// AsFullDocument wraps a compiled chunk into a standalone, header-guarded
// .h file: the guard, the common includes, plus any module-specific
// #include (a sibling generated header, or a resolved complex type's own
// header), then the actual body.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	guard := "EMISDK_" + strings.ToUpper(core.ToSnakeCase(x.SuggestedFileName)) + "_H"

	includes := make([]string, len(commonIncludes))
	copy(includes, commonIncludes)
	seen := map[string]struct{}{}
	for _, i := range includes {
		seen[i] = struct{}{}
	}
	for _, dep := range x.CodeChunkDependensies {
		if dep.Location == "" {
			continue
		}
		spelling := dep.Location
		if !strings.HasPrefix(spelling, "<") && !strings.HasPrefix(spelling, `"`) {
			spelling = `"` + spelling + `"`
		}
		if _, ok := seen[spelling]; ok {
			continue
		}
		seen[spelling] = struct{}{}
		includes = append(includes, spelling)
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "#ifndef %s\n#define %s\n\n", guard, guard)
	for _, i := range includes {
		fmt.Fprintf(&sb, "#include %s\n", i)
	}
	sb.WriteString("\n")
	sb.Write(x.ActualScript)
	fmt.Fprintf(&sb, "\n#endif /* %s */\n", guard)

	return string(core.EscapeLines([]byte(sb.String())))
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/C target, mirroring the same helper in the sibling generators.
// ImportLocation here is a header to #include (`"vector3.h"` or
// `<mymath/vector3.h>`), not a generated file path.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "c" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type CModuleGenerationFlags struct {
	Actions *string
	Remotes *string
	Dtos    *string
}

func (x CModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readCModuleFlags(ctx core.MicroGenContext) CModuleGenerationFlags {
	config := CModuleGenerationFlags{}
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

// CModuleFullVirtualFiles compiles an entire Emi module into a C header
// library: one header per dto/enum/action/remote, plus the embedded
// runtime (vendored cJSON + fetchx.h) and a short README on how to build,
// mirroring the equivalent full-module compilers in the sibling
// generators.
func CModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readCModuleFlags(ctx)
	files := []core.VirtualFile{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}
		chunk, err := CCommonObjectGenerator(dto.Fields, ctx, CCommonObjectContext{
			RootTypeName:        dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, enum := range module.Enums {
		chunk, err := CStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, action := range module.Actions {
		if len(actionFilter) > 0 && !slices.Contains(actionFilter, action.Name) {
			continue
		}
		chunk, err := CActionRender(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	if ctx.HasTag(Sqlite) || ctx.HasTag(Pack) {
		entityFiles, err := CEntityVirtualFiles(module, ctx, complexes)
		if err != nil {
			return nil, err
		}
		files = append(files, entityFiles...)
	}

	if ctx.HasTag(Kconfig) {
		kfile, err := CKconfigVirtualFile(module)
		if err != nil {
			return nil, err
		}
		if kfile != nil {
			files = append(files, *kfile)
		}
	}

	for _, remote := range module.Remotes {
		// a remote is an outbound call to someone else's server; the esp-server
		// target only generates what the device itself serves.
		if serverBackendHeader(ctx) != "" {
			break
		}
		if len(remoteFilter) > 0 && !slices.Contains(remoteFilter, remote.Name) {
			continue
		}
		chunk, err := CActionRender(remote, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	skipSdk := ctx.HasTag(NoSdk)
	if !skipSdk {
		backend := serverBackendHeader(ctx)
		for _, f := range core.FsEmbedToVirtualFile(&cInclude.Content, "") {
			// the server runtime comes as a core plus one backend per target;
			// ship only the backend that was asked for (none for the client).
			name := f.Name + f.Extension
			if strings.HasPrefix(name, "emihttpd_") && name != "emihttpd_core.h" && name != backend {
				continue
			}
			if name == "emihttpd_core.h" && backend == "" {
				continue
			}
			if name == "emisql.h" && !ctx.HasTag(Sqlite) {
				continue
			}
			if name == "emipack.h" && !ctx.HasTag(Pack) {
				continue
			}
			if name == "emilayout.h" && !ctx.HasTag(Sqlite) && !ctx.HasTag(Pack) {
				continue
			}
			files = append(files, f)
		}
	}

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "README",
			Extension:    ".md",
			ActualScript: cReadme,
		})
	}

	return files, nil
}

const cReadme = `# Generated Emi C SDK

Generated by the Emi compiler - do not edit by hand.

Every ` + "`.h`" + ` file here is header-only (` + "`static inline`" + ` throughout, like the vendored
` + "`cjson/cJSON.h`" + `) - just ` + "`#include`" + ` whichever ones you need, from as many ` + "`.c`" + ` files as you
like, and compile ` + "`cjson/cJSON.c`" + ` into your program once:

` + "```" + `sh
cc -I. your_program.c cjson/cJSON.c -lcurl -o your_program
` + "```" + `

Without libcurl (ESP-IDF, bare metal): compile with ` + "`-DFETCHX_NO_CURL`" + `. That keeps the url helpers
and ` + "`FetchxContext`" + ` and drops the http/SSE helpers. ` + "`method: reactive`" + ` actions with
` + "`binaryType: arraybuffer`" + ` need no libcurl at all: they generate a websocket opener over
` + "`emiws.h`" + ` (RFC 6455 client, BSD sockets or your own ` + "`EmiStream`" + `) and ` + "`emimedia.h`" + `
(microphone/camera frame header, ` + "`EmiMediaSender`" + `). See examples/media-stream.

Every generated struct ` + "`Type`" + ` gets: ` + "`Type_new()`" + `/` + "`Type_free()`" + ` (heap
alloc/dealloc, always safe to call), ` + "`Type_from_json(const cJSON*)`" + ` /
` + "`Type_to_json(const Type*)`" + ` for (de)serialization, and - since every field
always gets a safe default - casting one dto's data into a different (but
JSON-compatible) struct is just going through cJSON as the interchange
format:

` + "```" + `c
cJSON* json = SourceDto_to_json(source);
TargetDto* target = TargetDto_from_json(json);
cJSON_Delete(json);
/* ... use target ... */
TargetDto_free(target);
` + "```" + `

Server mode (` + "`--tags esp-server`" + `, ESP-IDF): instead of client functions every action header
gets ` + "`<action>_register(server, handler, user)`" + ` plus a typed ` + "`<Action>Call`" + ` / ` + "`<Action>Handler`" + `
that you implement; ` + "`method: reactive`" + ` actions get websocket sessions with ` + "`_send()`/`_broadcast()`" + `.
Set ` + "`config.uri_match_fn = emihttpd_uri_match`" + ` on the ` + "`httpd_config_t`" + ` so ` + "`/users/:id`" + ` routes match,
enable ` + "`CONFIG_HTTPD_WS_SUPPORT`" + ` for reactive actions (` + "`--tags mongoose-server`" + ` instead targets Mongoose: macOS/Linux/Windows/ESP-IDF, bring your own mongoose.c/.h), and add the ` + "`cjson`" + `/` + "`emihttpd.h`" + ` files to
your component. Example: examples/esp-server.

SQLite storage (` + "`--tags sqlite`" + `): every entity of the module also gets, in its own header, ` + "`Entity_insert/_get/_update/_patch/_delete/_list/_query`" + `
over sqlite3 (` + "`_query`" + ` takes an ` + "`EmiSqlQuery`" + `: a json-logic filter, a sort, paging, and returns the total;
` + "`_patch`" + ` changes only the fields of a json object), plus ` + "`<module>_schema_apply(db)`" + ` to create (and later extend) the tables - the same
tables ` + "`emi entity-sql --tags sqlite`" + ` writes as a script. Link with ` + "`-lsqlite3`" + ` (ESP-IDF: a sqlite3 component).
Scalars, flattened ` + "`object`" + ` fields and ` + "`array`" + `s of objects are stored; ` + "`one`/`collection`" + ` columns and
join tables are created but not read or written by these functions.

Packed entities (` + "`--tags pack`" + `, alone or with sqlite): every entity gets ` + "`Entity_pack(self, &buf, &len)`" + ` and
` + "`Entity_unpack(buf, len)`" + ` (` + "`_unpack_prefix`" + ` for several in a row), a compact binary string (no field names, varints) for a
file or an NVS blob where there is no SQLite; no sqlite3.h needed (` + "`-DEMI_NO_SQLITE`" + ` leaves the SQLite functions of a
header out). Only append fields to an entity which has been stored: a newer program reads an older blob. See emipack.h.

Action functions (` + "`create_widget_action(...)`" + ` etc.) return ` + "`NULL`" + ` on any
failure (transport-level or a non-2xx status) - pass a ` + "`FetchxResponse*`" + ` as
their last argument to inspect ` + "`.status_code`" + `/` + "`.error_message`" + ` on failure (and
` + "`FetchxResponse_free()`" + ` it when you're done either way).
`
