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
