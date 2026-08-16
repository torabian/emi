package js

import "github.com/torabian/emi/lib/core"

// List of all compiler tags javascript supports. Add them here before using them.
const Typescript core.CTag = "typescript"
const IncludeExt core.CTag = "include-ext"
const React core.CTag = "react"
const Nestjs core.CTag = "nestjs"
const NoPackage core.CTag = "no-package"
const NoEnvelope core.CTag = "no-envelope"
const NoSdk core.CTag = "no-sdk"
const NoJsDoc core.CTag = "no-jsdoc"           // skip the JSDoc @typedef section on plain JS output (see js-common-object-jsdoc.go)
const NoClass core.CTag = "no-class"           // skip the generated dto class body (both JS and TS) - keep only the type declaration (see js-common-object.go)
const NoDefinition core.CTag = "no-definition" // skip the `static Definition = {...}` JSON dump on every generated action class (see js-action-main-class.go)

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: Typescript, Description: "Generate TypeScript (.ts) output - typed classes/interfaces - instead of plain JavaScript"},
	{Tag: IncludeExt, Description: "Append the file extension (.js or .ts) to generated import paths"},
	{Tag: React, Description: "Also generate React Query hooks (useQuery/useMutation) for each action"},
	{Tag: Nestjs, Description: "Emit NestJS decorators (e.g. on headers/query params) for server-side integration"},
	{Tag: NoPackage, Description: "Skip generating package.json"},
	{Tag: NoEnvelope, Description: "Skip embedding the envelope (js-envelopes/ts-envelopes) runtime files"},
	{Tag: NoSdk, Description: "Skip embedding the SDK runtime (fetchx and friends) files entirely"},
	{Tag: NoJsDoc, Description: "Skip the JSDoc @typedef section on plain JS output"},
	{Tag: NoClass, Description: "Skip the generated DTO class body (both JS and TS) - keep only the type declaration"},
	{Tag: NoDefinition, Description: "Skip the `static Definition = {...}` JSON dump on every generated action class"},
}
