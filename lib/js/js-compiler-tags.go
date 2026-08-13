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
const NoJsDoc core.CTag = "no-jsdoc" // skip the JSDoc @typedef section on plain JS output (see js-common-object-jsdoc.go)
