package c

// cPrintfArg wraps a value expression for snprintf with the given format
// literal. Integer formats are always "%lld" with a (long long) cast: int32_t
// is `long` on 32-bit targets such as Xtensa/ESP-IDF (a plain "%d" there fails
// -Werror=format), `int` elsewhere, and a cast keeps all of them portable.
func cPrintfArg(format, expr string) string {
	if format == `"%lld"` {
		return "(long long) " + expr
	}
	return expr
}
