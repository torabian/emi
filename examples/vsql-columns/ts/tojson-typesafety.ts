// Compile-time (and, per the comment below, runtime-verified) proof that a
// generated dto class's toJSON() actually returns the same shape as its
// exported plain *Type - not just something JSON.stringify eventually
// resolves to correctly after further recursion.
//
// Before this file's underlying fix (JsCommonObjectClassGenerator's toJSON
// template, lib/js/js-common-object-class.go), this was NOT true: a nested
// array field's toJSON() returned the raw emigo.Array wrapper
// (MArray<Users>), not a plain UsersType[] - assigning
// `const x: InsertUsersVsqlParamsType = params.toJSON()` failed to compile
// with "Type 'MArray<Users>' is missing the following properties from type
// 'UsersType[]': length, pop, push, ...". The fix: every field goes through
// a private #toPlainJSON helper that recurses into anything carrying its own
// toJSON (a nested dto instance, or an emigo Array/One/Collection wrapper)
// instead of returning it as-is.
//
// Runtime behavior was checked too, not just types: for a users array with
// one entry, `params.toJSON().users` is a genuine plain Array (not an
// MArray instance), and its one item's `.constructor.name` is "Object" (not
// "Users") - confirmed by actually compiling and running this kind of
// assignment with node, once, before landing the fix.
import type { InsertUsersVsqlParams, InsertUsersVsqlParamsType } from "./sdkgen/InsertUsersVsqlParams";
import type { InsertUsersVsqlRow, InsertUsersVsqlRowType } from "./sdkgen/InsertUsersVsqlRow";
import type { InsertUsersVsqlFilters, InsertUsersVsqlFiltersType } from "./sdkgen/InsertUsersVsqlFilters";

// Each function's return statement is the actual check: if toJSON()'s real
// return type ever stops matching the declared plain Type, the assignment
// below fails to compile. Nothing here needs to run - only compile (see
// package.json's "typecheck").

function paramsToJSONMatchesType(p: InsertUsersVsqlParams): InsertUsersVsqlParamsType {
  return p.toJSON();
}

// InsertUsersVsqlRow is the interesting one: it carries an object column
// (profile, always present), a nullable object column (preferences), and a
// complex column (money) - all three go through the same #toPlainJSON path
// as the plain array case above.
function rowToJSONMatchesType(r: InsertUsersVsqlRow): InsertUsersVsqlRowType {
  return r.toJSON();
}

function filtersToJSONMatchesType(f: InsertUsersVsqlFilters): InsertUsersVsqlFiltersType {
  return f.toJSON();
}

void [paramsToJSONMatchesType, rowToJSONMatchesType, filtersToJSONMatchesType];
