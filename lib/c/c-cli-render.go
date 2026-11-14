// Renders emi_cli.h (`--tags mongoose-server,cli`): the command table of the module. Every non reactive action with a
// url is a command, its path parameters, query fields and the top level fields of its request are the flags; emicli.h
// turns them into a request line and runs it through the router the handlers were registered on. Only data is
// generated per action, so a module of a hundred actions costs a table, not a hundred functions.
package c

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type cCliFlag struct {
	Flag  string
	Place string // EMICLI_PATH, EMICLI_QUERY, EMICLI_BODY
	Key   string
	Type  string // EMICLI_STRING, ...
	Help  string
}

type cCliCommand struct {
	Name, Alias, Method, Path, Help string
	Flags                           []cCliFlag
	HasBody                         bool
	Var                             string
	Full                            string // the command from the name of the action, when a cliName is not unique
}

// cKebab spells a camel case or bracketed name as a flag / command name: setPoint -> set-point, filter[page][n] ->
// filter-page-n.
func cKebab(s string) string {
	s = strings.NewReplacer("[", "_", "]", "", ".", "_", " ", "_").Replace(s)
	return strings.Trim(strings.ReplaceAll(core.ToSnakeCase(s), "_", "-"), "-")
}

func cCliScalarType(ctype string) string {
	switch ctype {
	case "bool":
		return "EMICLI_BOOL"
	case "float", "double":
		return "EMICLI_FLOAT"
	case "char*", "":
		return "EMICLI_STRING"
	}
	return "EMICLI_INT"
}

func cCliTypeWord(t string) string {
	switch t {
	case "EMICLI_BOOL":
		return "bool"
	case "EMICLI_FLOAT":
		return "number"
	case "EMICLI_INT":
		return "int"
	case "EMICLI_JSON":
		return "json"
	}
	return "string"
}

// cCliActionHelp is the first sentence of the description of the action, which is what a listing of the commands has room for.
func cCliActionHelp(action core.EmiRpcAction) string {
	a, ok := action.(*core.EmiAction)
	if !ok {
		return ""
	}
	text := strings.Join(strings.Fields(a.Description), " ")
	if i := strings.Index(text, ". "); i >= 0 {
		text = text[:i+1]
	}
	if len(text) > 110 {
		text = text[:107] + "..."
	}
	return text
}

func cCliCommandFor(action core.EmiRpcAction, ctx core.MicroGenContext, complexes []RecognizedComplex) (*cCliCommand, error) {
	realms, err := cGetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}
	if !realms.HasUrl || realms.IsReactive {
		return nil, nil
	}
	hasBody := realms.RequestTypeName != "" && !realms.RequestIsAmbiguous
	method, err := serverMethod(realms.HttpMethod, hasBody)
	if err != nil {
		return nil, err
	}
	cmd := &cCliCommand{
		Name:    action.GetCliName(),
		Alias:   action.GetCliShort(),
		Method:  strings.TrimPrefix(method, "EMIHTTPD_"),
		Path:    serverUrlPath(realms.Url),
		Help:    cCliActionHelp(action),
		HasBody: hasBody,
		Var:     "emi_cli_flags_" + realms.FunctionName,
	}
	cmd.Full = cKebab(strings.TrimSuffix(realms.ActionName, "Action"))
	if cmd.Name == "" {
		cmd.Name = cmd.Full
	}

	// the flag names of a command must be unique, and --help, --body and --pretty are the command's own
	taken := map[string]bool{"help": true, "body": true, "pretty": true}
	add := func(group, name, place, key, typ, help string) {
		flag := cKebab(name)
		if taken[flag] {
			flag = group + "-" + flag
		}
		for i := 2; taken[flag]; i++ {
			flag = fmt.Sprintf("%s-%s-%d", group, cKebab(name), i)
		}
		taken[flag] = true
		cmd.Flags = append(cmd.Flags, cCliFlag{Flag: flag, Place: place, Key: key, Type: typ, Help: help})
	}

	for _, p := range cBuildPathParams(action) {
		typ := cCliScalarType(p.Type)
		add("path", p.WireName, "EMICLI_PATH", p.WireName, typ, "path parameter, "+cCliTypeWord(typ)+", required")
	}
	_, leaves := cQueryTree(action.GetQuery())
	for _, l := range leaves {
		typ := cCliScalarType(l.Type)
		add("query", l.Key, "EMICLI_QUERY", l.Key, typ, "query, "+cCliTypeWord(typ))
	}
	if hasBody {
		for _, f := range action.GetRequestFields() {
			typ := "EMICLI_JSON"
			switch ft := f.GetType(); {
			case ft == core.FieldTypeString || ft == core.FieldTypeEnum:
				typ = "EMICLI_STRING"
			case extractPrimitive(string(ft)) != "":
				typ = cCliScalarType(extractPrimitive(string(ft)))
			}
			help := "request field, " + cCliTypeWord(typ)
			if typ == "EMICLI_JSON" {
				help += " (the text, @file or -)"
			}
			if d := strings.Join(strings.Fields(f.GetDescription()), " "); d != "" {
				help += ": " + d
			}
			add("body", f.GetName(), "EMICLI_BODY", f.GetName(), typ, help)
		}
	}
	return cmd, nil
}

// CCliVirtualFile renders emi_cli.h, nil when the module has nothing to run.
func CCliVirtualFile(module *core.Emi, ctx core.MicroGenContext, complexes []RecognizedComplex) (*core.VirtualFile, error) {
	cmds := []*cCliCommand{}
	used := map[string]int{}
	for _, action := range module.Actions {
		cmd, err := cCliCommandFor(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if cmd == nil {
			continue
		}
		used[cmd.Name]++
		cmds = append(cmds, cmd)
	}
	// a cliName which several actions share (the `create` of every entity) is not a command: those are called by the name of their action
	seen := map[string]bool{}
	for _, cmd := range cmds {
		if used[cmd.Name] > 1 {
			cmd.Name, cmd.Alias = cmd.Full, ""
		}
		if seen[cmd.Name] {
			return nil, fmt.Errorf("cli: two actions are the command %q: give one a cliName", cmd.Name)
		}
		seen[cmd.Name] = true
	}
	if len(cmds) == 0 {
		return nil, nil
	}

	var b strings.Builder
	b.WriteString(`#ifndef EMISDK_EMI_CLI_H
#define EMISDK_EMI_CLI_H

#include "emihttpd_mongoose.h"
#include "emicli.h"

/* One command per action: ` + "`emi_cli_main(&router, argc, argv)`" + ` runs it on the router the *_register functions filled. */
`)
	for _, c := range cmds {
		if len(c.Flags) == 0 {
			continue
		}
		fmt.Fprintf(&b, "static const EmiCliFlag %s[] = {\n", c.Var)
		for _, f := range c.Flags {
			fmt.Fprintf(&b, "    {%s, %s, %s, %s, %s},\n", escapeDoubleQuoted(f.Flag), f.Place, escapeDoubleQuoted(f.Key), f.Type, escapeDoubleQuoted(f.Help))
		}
		b.WriteString("};\n")
	}
	b.WriteString("\nstatic const EmiCliCommand emi_cli_commands[] = {\n")
	for _, c := range cmds {
		flags, count := "NULL", "0"
		if len(c.Flags) > 0 {
			flags, count = c.Var, fmt.Sprintf("sizeof(%s) / sizeof(%s[0])", c.Var, c.Var)
		}
		fmt.Fprintf(&b, "    {%s, %s, %s, %s, %s, %s, %s, %v},\n", escapeDoubleQuoted(c.Name), escapeDoubleQuoted(c.Alias), escapeDoubleQuoted(c.Method),
			escapeDoubleQuoted(c.Path), escapeDoubleQuoted(c.Help), flags, count, c.HasBody)
	}
	b.WriteString(`};

#define EMI_CLI_COMMAND_COUNT (sizeof(emi_cli_commands) / sizeof(emi_cli_commands[0]))

static inline char* emi_cli_call(void* router, const char* line) {
    return emihttpd_text_call((EmiHttpdRouter*) router, line);
}

/* Returns the exit code: 0 for a 2xx answer, 1 for any other, 2 for a usage error. */
static inline int emi_cli_main(EmiHttpdRouter* router, int argc, const char** argv) {
    return emicli_main(emi_cli_commands, EMI_CLI_COMMAND_COUNT, argc, argv, emi_cli_call, router);
}

#endif /* EMISDK_EMI_CLI_H */
`)
	return &core.VirtualFile{Name: "emi_cli", Extension: ".h", ActualScript: b.String()}, nil
}
