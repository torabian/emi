package js

var TOKEN_ROOT_CLASS = "root.class"
var TOKEN_OBJ_CLASS = "root.request.class"
var TOKEN_OBJ_TYPE = "root.request.class"

// TOKEN_TYPEDEF_NAME is the real, `Type`-suffixed identifier of the standalone
// type declaration a dto/fields chunk produces alongside its class - the
// `export type XType = {...}` in TypeScript mode (js-common-object-types.go),
// or the `@typedef {Object} XType` in plain JS mode (js-common-object-jsdoc.go).
// Deliberately NOT the same string as TOKEN_OBJ_CLASS/TOKEN_OBJ_TYPE above: those
// two intentionally collide (a TS class name doubles as its own type, so "the
// canonical name" is correct for both), but the *_Type_-suffixed standalone
// declaration is a genuinely different identifier that needs its own,
// unambiguous token - see the `no-class` tag (js-compiler-tags.go), the one
// consumer that actually needs to resolve it.
var TOKEN_TYPEDEF_NAME = "root.typedef.name"
var TOKEN_ORIGINAL_NAME = "root.original.name"
var TOKEN_NEW_URL_FN = "new.url.fn"
var TOKEN_URL_METHOD = "new.url.method"
var TOKEN_ACTUAL_METHOD = "new.actual.method"
var TOKEN_CREATOR_FN = "new.creatorFn"
var TOKEN_RESPONSE_ENVELOPE = "response.envelope"
