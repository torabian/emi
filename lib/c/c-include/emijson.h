/*
 * emijson.h - writes the JSON text of a generated struct without building a cJSON tree first.
 *
 * Every generated struct has `Type_emit(const Type*, EmiJson*)` beside `Type_to_json`: it writes the same text `cJSON_PrintUnformatted(Type_to_json(x))` makes, byte for byte
 * (the numbers and the escapes follow cJSON's own rules). It is run twice by emihttpd_reply_emit: with no destination it only counts the bytes (the Content-Length), then
 * it writes them straight into the connection's send buffer, which was made the size of the answer once. A tree of JSON nodes (about 280 bytes for each small value) and a
 * text which is printed into a buffer which doubles are never made: on a chip with little RAM they were three to ten times the answer in memory at the same moment.
 */
#ifndef EMIJSON_H
#define EMIJSON_H

#include <float.h>
#include <limits.h>
#include <math.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "cjson/cJSON.h"

typedef struct EmiJson {
    char* dst;  /* NULL: only count */
    size_t len; /* bytes written (or counted) so far */
    size_t cap; /* the room in dst: a write past it is counted but not made */
} EmiJson;

static inline void emijson_raw(EmiJson* j, const char* s, size_t n) {
    if (j->dst && j->len + n <= j->cap) memcpy(j->dst + j->len, s, n);
    j->len += n;
}

static inline void emijson_ch(EmiJson* j, char c) { emijson_raw(j, &c, 1); }

/* a string literal, as it is */
#define emijson_lit(j, lit) emijson_raw((j), lit, sizeof(lit) - 1)

/* a JSON string: quoted, with the escapes cJSON makes (" \ and the control characters; \b \f \n \r \t and \u00xx for the rest) */
static inline void emijson_str(EmiJson* j, const char* s) {
    emijson_ch(j, '"');
    const char* run = s;
    for (const char* p = s; p && *p; p++) {
        unsigned char c = (unsigned char) *p;
        if (c > 31 && c != '"' && c != '\\') continue;
        if (p > run) emijson_raw(j, run, (size_t) (p - run));
        run = p + 1;
        switch (c) {
            case '\\': emijson_raw(j, "\\\\", 2); break;
            case '"': emijson_raw(j, "\\\"", 2); break;
            case '\b': emijson_raw(j, "\\b", 2); break;
            case '\f': emijson_raw(j, "\\f", 2); break;
            case '\n': emijson_raw(j, "\\n", 2); break;
            case '\r': emijson_raw(j, "\\r", 2); break;
            case '\t': emijson_raw(j, "\\t", 2); break;
            default: {
                char esc[8];
                int n = snprintf(esc, sizeof(esc), "\\u%04x", c);
                emijson_raw(j, esc, (size_t) n);
            }
        }
    }
    if (s && *run) emijson_raw(j, run, strlen(run));
    emijson_ch(j, '"');
}

/* cJSON's compare_double: the same number to within the precision of a double */
static inline int emijson_same(double a, double b) {
    double m = fabs(a) > fabs(b) ? fabs(a) : fabs(b);
    return fabs(a - b) <= m * DBL_EPSILON;
}

/* a number the way cJSON prints it: an integer which fits an int as an int, else the shortest of %1.15g / %1.17g which reads back as the same double; null for NaN and Infinity */
static inline void emijson_num(EmiJson* j, double d) {
    char buf[32];
    int n;
    if (isnan(d) || isinf(d)) {
        n = snprintf(buf, sizeof(buf), "null");
    } else {
        int vi = d >= (double) INT_MAX ? INT_MAX : (d <= (double) INT_MIN ? INT_MIN : (int) d); /* cJSON's valueint */
        if (d == (double) vi) {
            n = snprintf(buf, sizeof(buf), "%d", vi);
        } else {
            double test = 0.0;
            n = snprintf(buf, sizeof(buf), "%1.15g", d);
            if (sscanf(buf, "%lg", &test) != 1 || !emijson_same(test, d)) n = snprintf(buf, sizeof(buf), "%1.17g", d);
        }
    }
    if (n < 0 || n >= (int) sizeof(buf)) n = snprintf(buf, sizeof(buf), "null");
    for (int i = 0; i < n; i++) if (buf[i] == ',') buf[i] = '.'; /* a locale with a decimal comma */
    emijson_raw(j, buf, (size_t) n);
}

static inline void emijson_bool(EmiJson* j, int b) { b ? emijson_raw(j, "true", 4) : emijson_raw(j, "false", 5); }
static inline void emijson_null(EmiJson* j) { emijson_raw(j, "null", 4); }

/* the key of the next member of an object: a comma unless it is the first */
static inline void emijson_key(EmiJson* j, int* first, const char* key) {
    if (!*first) emijson_ch(j, ',');
    *first = 0;
    emijson_str(j, key);
    emijson_ch(j, ':');
}

/* a value which is already a cJSON tree (an `any` field): printed for the moment */
static inline void emijson_cjson(EmiJson* j, const cJSON* v) {
    char* text = v ? cJSON_PrintUnformatted(v) : NULL;
    if (text) emijson_raw(j, text, strlen(text)); else emijson_null(j);
    free(text);
}

#endif /* EMIJSON_H */
