//go:build wasm

package emigo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"syscall/js"
)

type WasmServerConfig struct {

	// Customize the function which is being globally exposed to js side. Default is handleWasmRequest
	JsFunctionName string
}

func LiftWasmServer(mux *http.ServeMux, config *WasmServerConfig) {

	// Single entry point exposed to JS. It turns a fetch-style call into a real
	// *http.Request, runs it through the ServeMux, and captures the response with
	// an in-memory httptest.ResponseRecorder. This is exactly the request loop a
	// real server runs, one call at a time, driven by the browser.
	//
	//   handleWasmRequest(method, url, body, headersJSON) -> JSON string

	jsFunctionName := "handleWasmRequest"
	if config != nil && config.JsFunctionName != "" {
		jsFunctionName = config.JsFunctionName
	}

	js.Global().Set(jsFunctionName, js.FuncOf(func(_ js.Value, args []js.Value) any {
		method := args[0].String()
		url := args[1].String()
		bodyStr := args[2].String()
		headersJSON := args[3].String()

		// Return a Promise rather than the response directly. Handlers may make
		// async calls back into JS (e.g. a pglite database query exposed as a
		// promise-returning function). Those promises can only settle while the
		// JS event loop is running — but a synchronous Go function invoked from
		// JS never yields the event loop, so awaiting one inside it deadlocks.
		// By running ServeHTTP in a goroutine and resolving a Promise when it
		// finishes, we hand control back to the event loop immediately, letting
		// nested promises resolve and the blocked goroutine make progress.
		executor := js.FuncOf(func(_ js.Value, promiseArgs []js.Value) any {
			resolve := promiseArgs[0]
			reject := promiseArgs[1]

			go func() {
				defer func() {
					if r := recover(); r != nil {
						reject.Invoke(fmt.Sprintf("wasm handler panic: %v", r))
					}
				}()

				req := httptest.NewRequest(method, url, strings.NewReader(bodyStr))

				var hdrs map[string]string
				if json.Unmarshal([]byte(headersJSON), &hdrs) == nil {
					for k, v := range hdrs {
						req.Header.Set(k, v)
					}
				}

				rec := httptest.NewRecorder()
				mux.ServeHTTP(rec, req) // real routing, real handler, real ResponseWriter

				res := rec.Result()
				bodyBytes, _ := io.ReadAll(res.Body)

				out, _ := json.Marshal(map[string]any{
					"status":  res.StatusCode,
					"headers": res.Header, // map[string][]string serializes naturally
					"body":    string(bodyBytes),
				})
				resolve.Invoke(string(out))
			}()

			return nil
		})
		// The executor runs synchronously inside Promise's constructor, so it is
		// safe to release once New returns; the goroutine it spawned lives on.
		defer executor.Release()

		return js.Global().Get("Promise").New(executor)
	}))

	fmt.Println("Wasm router ready: window." + jsFunctionName + " (method, url, body, headersJSON)")

}
