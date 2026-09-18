// TS-side counterpart of ../../sdkgen/money.go's Money type (see that
// file's doc comment) - registered for the JS/TS compiler only via this
// module's `complexes:` (compiler: js), so the generated classes referencing
// `complex: Money` import a real type instead of a bare, unresolved name.
//
// A generated complex-typed setter always constructs via `new Money(value)`
// - one raw argument, not a multi-arg constructor (see
// examples/js-test/reactclient/src/Money.ts for the same convention) - so
// this takes whatever shape the field's value arrives as (typically the
// already-decoded {amountCents, currency} object) and copies it over.
export class Money {
  amountCents: number;
  currency: string;

  constructor(data: { amountCents: number; currency: string }) {
    this.amountCents = data.amountCents;
    this.currency = data.currency;
  }
}
