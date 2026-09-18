// Compile-time proof that a vsql's JSON-Logic filter type only accepts
// fields the query actually declares (InsertUsersVsqlFilterField, defaulted
// from batch_insert_users.emi.yml's `columns:` - see EmiVsql.Filters). Run
// `npm run typecheck` in this directory: every `@ts-expect-error` line below
// must actually fail to typecheck, or tsc reports an "unused directive"
// error and the run fails - so this file only stays green as long as the
// constraint genuinely holds, not just at the moment it was written.
//
// There's no third-party TypeScript typing of JSON-Logic itself to build on
// here: JsonLogic's own spec (https://jsonlogic.com) doesn't ship one, and a
// generic library couldn't know a *specific* query's allowed field names
// anyway - only Emi, generating per-vsql, can. So InsertUsersVsqlFilter
// below is Emi's own type, not a wrapped third-party one; see
// jsVsqlFilterTypeAppendix's doc comment (lib/js/js-vsql.go) for exactly
// what it does and doesn't constrain. It stays plain-JSON-shaped, though -
// nothing stops handing a value of this type straight to a real JsonLogic
// runtime (e.g. json-logic-js) exactly as any other JsonLogic value.
import type { InsertUsersVsqlFilter } from "./sdkgen/InsertUsersVsqlFilters";

// ---- Valid filters: every {var: ...} names a real, allow-listed field ----

const byEmail: InsertUsersVsqlFilter = {
  "==": [{ var: "email" }, "alice@example.com"],
};

const byStatusAndRating: InsertUsersVsqlFilter = {
  and: [
    { "==": [{ var: "status" }, "Active"] },
    { ">": [{ var: "rating" }, 4] },
  ],
};

const notSuspended: InsertUsersVsqlFilter = { "!": { var: "isSuspended" } };

// A bare {var: ...} node, not wrapped in an operator, is legal too.
const bareVar: InsertUsersVsqlFilter = { var: "money" };

// ---- Invalid filters: a field the query never declared -------------------

// Each of these is written as a single line on purpose: @ts-expect-error
// only suppresses (and requires) an error on the very next line, and a
// multi-line object literal can report its assignability error against a
// nested line rather than the first one - keeping the whole statement on
// one line means there's no ambiguity about which line the error lands on.

// @ts-expect-error - "isAdmin" isn't one of InsertUsersVsqlFilterField; this vsql never declared it as a column/filter.
const unknownFieldNested: InsertUsersVsqlFilter = { "==": [{ var: "isAdmin" }, true] };

// @ts-expect-error - same mistake, as a direct var node this time.
const unknownFieldDirect: InsertUsersVsqlFilter = { var: "isAdmin" };

// @ts-expect-error - a typo of a real field ("emial") is just as invalid as a field that was never real to begin with - the union only has room for the eleven exact names above.
const typoedField: InsertUsersVsqlFilter = { "==": [{ var: "emial" }, "alice@example.com"] };

// Referenced so a stricter lint config (unused-locals) never flags these -
// the point of each is the assignment itself typechecking (or not), not
// anything read from them afterwards.
void [
  byEmail,
  byStatusAndRating,
  notSuspended,
  bareVar,
  unknownFieldNested,
  unknownFieldDirect,
  typoedField,
];
