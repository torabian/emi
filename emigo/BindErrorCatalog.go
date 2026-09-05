package emigo

import (
	"fmt"
	"strings"
)

/**
* Translations for the fixed set of messages BindErrors.go's ToPublicJSON
* methods produce. Scoped to emi itself - no dependency on fireback/ferror
* or any application's own catalog - covering the languages emi's own
* tooling/docs target: English (the required fallback), Persian, Polish,
* Russian, and Spanish. An application with its own localization story
* switches on the concrete *Bind*Error type instead of using these at all
* (see BindErrors.go's package doc comment).
**/

// bindErrorCatalog maps a machine-readable error code (the same string
// returned as the "message" field) to per-language message templates.
// A template with fmt.Sprintf verbs is filled from the args its
// ToPublicJSON call site passes - see each call in BindErrors.go for the
// argument order. "en" must be present for every code; it's the fallback
// used for a language this catalog doesn't have a translation for.
var bindErrorCatalog = map[string]map[string]string{
	"BodyIsEmptyEof": {
		"en": "Body is empty. Please provide the necessary data and try again.",
		"fa": "بدنه درخواست خالی است. لطفاً داده‌های لازم را ارسال کرده و دوباره تلاش کنید.",
		"pl": "Treść żądania jest pusta. Podaj wymagane dane i spróbuj ponownie.",
		"ru": "Тело запроса пустое. Пожалуйста, укажите необходимые данные и попробуйте снова.",
		"es": "El cuerpo de la solicitud está vacío. Proporcione los datos necesarios e inténtelo de nuevo.",
	},
	"BodyUnexpectedEof": {
		"en": "Body unexpected EOF. The data you sent appears incomplete. Please check your request and try again.",
		"fa": "پایان غیرمنتظره بدنه درخواست. داده‌های ارسالی شما ناقص به نظر می‌رسند. لطفاً درخواست خود را بررسی کرده و دوباره تلاش کنید.",
		"pl": "Nieoczekiwany koniec treści żądania. Wysłane dane wydają się niekompletne. Sprawdź żądanie i spróbuj ponownie.",
		"ru": "Неожиданный конец тела запроса. Отправленные данные выглядят неполными. Проверьте запрос и попробуйте снова.",
		"es": "Fin inesperado del cuerpo de la solicitud. Los datos enviados parecen incompletos. Revise su solicitud e inténtelo de nuevo.",
	},
	"BodyReadAfterClose": {
		"en": "Body is read after closed. The request might have been processed incorrectly.",
		"fa": "بدنه درخواست پس از بسته شدن خوانده شده است. ممکن است درخواست به‌درستی پردازش نشده باشد.",
		"pl": "Treść żądania odczytano po zamknięciu. Żądanie mogło zostać przetworzone nieprawidłowo.",
		"ru": "Тело запроса прочитано после закрытия. Возможно, запрос был обработан некорректно.",
		"es": "El cuerpo de la solicitud se leyó después de cerrarse. Es posible que la solicitud se haya procesado incorrectamente.",
	},
	"UnknownErrorReadingBody": {
		"en": "We cannot read the body of your request.",
		"fa": "امکان خواندن بدنه درخواست شما وجود ندارد.",
		"pl": "Nie można odczytać treści żądania.",
		"ru": "Не удалось прочитать тело запроса.",
		"es": "No podemos leer el cuerpo de su solicitud.",
	},
	// args: field, expected type, actual value, line, col.
	"JsonInvalidFieldType": {
		"en": "The field %q must be of type %s, but received %v (line %d, column %d).",
		"fa": "فیلد %q باید از نوع %s باشد، اما %v دریافت شد (خط %d، ستون %d).",
		"pl": "Pole %q powinno być typu %s, ale otrzymano %v (linia %d, kolumna %d).",
		"ru": "Поле %q должно быть типа %s, но получено %v (строка %d, столбец %d).",
		"es": "El campo %q debe ser de tipo %s, pero se recibió %v (línea %d, columna %d).",
	},
	// args: line, col.
	"JsonMalformed": {
		"en": "Body is malformed at line %d, column %d. Check your commas, braces, tags, etc.",
		"fa": "بدنه درخواست در خط %d، ستون %d نامعتبر است. کاماها، آکولادها و سایر نشانه‌ها را بررسی کنید.",
		"pl": "Treść żądania jest nieprawidłowa w linii %d, kolumnie %d. Sprawdź przecinki, nawiasy klamrowe itp.",
		"ru": "Тело запроса содержит ошибку в строке %d, столбце %d. Проверьте запятые, скобки и т.д.",
		"es": "El cuerpo de la solicitud no es válido en la línea %d, columna %d. Revise las comas, llaves, etc.",
	},
	// args: line, col.
	"XmlMalformed": {
		"en": "The XML format is broken or incomplete at line %d, column %d. Please make sure all tags are properly opened and closed.",
		"fa": "قالب XML در خط %d، ستون %d ناقص یا نادرست است. لطفاً مطمئن شوید همه تگ‌ها به‌درستی باز و بسته شده‌اند.",
		"pl": "Format XML jest uszkodzony lub niekompletny w linii %d, kolumnie %d. Upewnij się, że wszystkie znaczniki są poprawnie otwarte i zamknięte.",
		"ru": "Формат XML повреждён или неполный в строке %d, столбце %d. Убедитесь, что все теги правильно открыты и закрыты.",
		"es": "El formato XML está dañado o incompleto en la línea %d, columna %d. Asegúrese de que todas las etiquetas estén correctamente abiertas y cerradas.",
	},
	"XmlUnmarshalError": {
		"en": "The XML structure doesn't match the expected format. Some elements may be missing or in the wrong place.",
		"fa": "ساختار XML با قالب مورد انتظار مطابقت ندارد. ممکن است برخی عناصر گم شده یا در جای نادرستی قرار گرفته باشند.",
		"pl": "Struktura XML nie odpowiada oczekiwanemu formatowi. Niektóre elementy mogą brakować lub znajdować się w niewłaściwym miejscu.",
		"ru": "Структура XML не соответствует ожидаемому формату. Некоторые элементы могут отсутствовать или быть расположены неправильно.",
		"es": "La estructura XML no coincide con el formato esperado. Es posible que falten algunos elementos o estén en el lugar equivocado.",
	},
	"JsonUnmarshalUnsupportedType": {
		"en": "Unsupported type when unmarshalling body.",
		"fa": "نوع پشتیبانی‌نشده هنگام پردازش بدنه درخواست.",
		"pl": "Nieobsługiwany typ podczas przetwarzania treści żądania.",
		"ru": "Неподдерживаемый тип при разборе тела запроса.",
		"es": "Tipo no compatible al procesar el cuerpo de la solicitud.",
	},
	"YamlTypeError": {
		"en": "One of the values is in the wrong format. For example, you might have entered text instead of a number or used quotes incorrectly.",
		"fa": "یکی از مقادیر دارای قالب نادرستی است. برای مثال ممکن است به‌جای عدد، متن وارد کرده یا از نقل‌قول‌ها به‌اشتباه استفاده کرده باشید.",
		"pl": "Jedna z wartości ma nieprawidłowy format. Na przykład mogłeś wpisać tekst zamiast liczby lub niepoprawnie użyć cudzysłowów.",
		"ru": "Одно из значений имеет неверный формат. Например, вы могли ввести текст вместо числа или неправильно использовать кавычки.",
		"es": "Uno de los valores tiene un formato incorrecto. Por ejemplo, es posible que haya introducido texto en lugar de un número o usado comillas incorrectamente.",
	},
	"YamlDecodingError": {
		"en": "There's something wrong with the format of your YAML. Please check indentation, colons, and line breaks to fix the formatting.",
		"fa": "مشکلی در قالب YAML شما وجود دارد. لطفاً تورفتگی‌ها، دونقطه‌ها و خطوط جدید را بررسی کنید.",
		"pl": "Coś jest nie tak z formatem YAML. Sprawdź wcięcia, dwukropki i podziały wierszy.",
		"ru": "Что-то не так с форматом YAML. Проверьте отступы, двоеточия и переносы строк.",
		"es": "Hay un problema con el formato de su YAML. Revise la sangría, los dos puntos y los saltos de línea.",
	},
	"XmlDecodingError": {
		"en": "Something went wrong while processing the XML. Please check the content or try again later.",
		"fa": "هنگام پردازش XML خطایی رخ داد. لطفاً محتوا را بررسی کرده یا بعداً دوباره تلاش کنید.",
		"pl": "Wystąpił błąd podczas przetwarzania XML. Sprawdź zawartość lub spróbuj ponownie później.",
		"ru": "При обработке XML произошла ошибка. Проверьте содержимое или повторите попытку позже.",
		"es": "Ocurrió un error al procesar el XML. Revise el contenido o inténtelo de nuevo más tarde.",
	},
	"FormDataMalformed": {
		"en": "The form data submitted is malformed or contains invalid fields. Please check the form and ensure all required fields are properly filled out.",
		"fa": "داده‌های فرم ارسالی نامعتبر است یا شامل فیلدهای نادرست است. لطفاً فرم را بررسی کرده و از پر بودن صحیح فیلدهای الزامی مطمئن شوید.",
		"pl": "Przesłane dane formularza są nieprawidłowe lub zawierają błędne pola. Sprawdź formularz i upewnij się, że wszystkie wymagane pola są poprawnie wypełnione.",
		"ru": "Отправленные данные формы неверны или содержат некорректные поля. Проверьте форму и убедитесь, что все обязательные поля заполнены правильно.",
		"es": "Los datos del formulario enviados no son válidos o contienen campos incorrectos. Revise el formulario y asegúrese de que todos los campos obligatorios estén completos.",
	},
	"JsonDecodingError": {
		"en": "Unknown error happened upon decoding.",
		"fa": "خطای ناشناخته‌ای هنگام پردازش رخ داد.",
		"pl": "Wystąpił nieznany błąd podczas dekodowania.",
		"ru": "При декодировании произошла неизвестная ошибка.",
		"es": "Ocurrió un error desconocido durante la decodificación.",
	},
}

// bindErrorMessage resolves code's template for lang - falling back to
// English when lang isn't in the catalog for that code, or when code
// itself isn't in the catalog at all (in which case code is returned
// as-is, so a caller always gets *something* rather than an empty
// string) - and formats it with args when there are any.
func bindErrorMessage(code string, lang string, args ...any) string {
	templates, ok := bindErrorCatalog[code]
	if !ok {
		return code
	}
	lang = strings.ToLower(strings.TrimSpace(lang))
	tmpl, ok := templates[lang]
	if !ok {
		tmpl = templates["en"]
	}
	if len(args) == 0 {
		return tmpl
	}
	return fmt.Sprintf(tmpl, args...)
}
