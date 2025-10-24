package main

import (
	"fmt"
	"net/http"
	"strconv"

	"test.com/emi-go-header-demo/mathlib"
)

/**
* Anonymouse struct
* Auto-generated from emi go header module
 */
type Anonymouse struct {
	headers map[string]string
}

func NewAnonymouse() *Anonymouse {
	return &Anonymouse{headers: make(map[string]string)}
}

// A simple string in header, common as everywhere else.
func (h *Anonymouse) SetContentType(v string) {
	h.headers[http.CanonicalHeaderKey("content-type")] = v
}

// A simple string in header, common as everywhere else.
func (h *Anonymouse) ContentType() string {
	return h.headers[http.CanonicalHeaderKey("content-type")]
}

// int, how the golang converts int to make it easy for developer
func (h *Anonymouse) SetCacheExpire(v int) {
	h.headers[http.CanonicalHeaderKey("cache-expire")] = strconv.FormatInt(int64(v), 10)
}

// int, how the golang converts int to make it easy for developer
func (h *Anonymouse) CacheExpire() (int, error) {
	s := h.headers[http.CanonicalHeaderKey("cache-expire")]
	if s == "" {
		return 0, nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	return int(i), err
}

// Fully custom defined type, which would be able to access as Vector3
func (h *Anonymouse) SetPosition(v mathlib.Vector3) {
	// Complex type implements ToString()
	h.headers[http.CanonicalHeaderKey("position")] = v.ToString()
}

// Fully custom defined type, which would be able to access as Vector3
func (h *Anonymouse) Position() (mathlib.Vector3, error) {
	s := h.headers[http.CanonicalHeaderKey("position")]
	var v mathlib.Vector3
	if s != "" {
		if err := v.FromString(s); err != nil {
			return mathlib.Vector3{}, fmt.Errorf("invalid header 'position': %w", err)
		}
	}
	return v, nil
}

// Wrapping go internal big.Float into complex, and using it in the headers
func (h *Anonymouse) SetEnergy(v mathlib.BigFloat) {
	// Complex type implements ToString()
	h.headers[http.CanonicalHeaderKey("energy")] = v.ToString()
}

// Wrapping go internal big.Float into complex, and using it in the headers
func (h *Anonymouse) Energy() (mathlib.BigFloat, error) {
	s := h.headers[http.CanonicalHeaderKey("energy")]
	var v mathlib.BigFloat
	if s != "" {
		if err := v.FromString(s); err != nil {
			return mathlib.BigFloat{}, fmt.Errorf("invalid header 'energy': %w", err)
		}
	}
	return v, nil
}
func (h *Anonymouse) Set(key string, val string) {
	h.headers[http.CanonicalHeaderKey(key)] = val
}
func (h *Anonymouse) Get(key string) string {
	return h.headers[http.CanonicalHeaderKey(key)]
}
func (h *Anonymouse) PopulateFromHTTP(httpHeaders http.Header) {
	for k, values := range httpHeaders {
		if len(values) > 0 {
			// Canonicalize header key and take first value
			h.headers[http.CanonicalHeaderKey(k)] = values[0]
		}
	}
}
func (h *Anonymouse) WriteToHTTP(w http.ResponseWriter) {
	// Optionally, you can set default or computed values here
	for k, v := range h.headers {
		fmt.Println("Writing to response header:", k, v)
		w.Header().Set(k, v)
	}
}
