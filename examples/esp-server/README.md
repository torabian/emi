# esp-server: a device (or any C program) that serves its own emi API

`thermo.emi.yml` is the contract. With a server tag the C compiler generates the **server** side
(the same way `emi go` generates gin handlers): routing, path/query/body decoding, JSON and websocket
sessions are generated, you implement the behaviour. The generated action headers are identical for
both backends - only the runtime header differs:

| tag | backend | runs on |
|---|---|---|
| `--tags esp-server` | ESP-IDF `esp_http_server` | ESP32 family |
| `--tags mongoose-server` | [Mongoose](https://mongoose.ws) event loop | macOS, Linux, Windows, ESP-IDF |

Mongoose is GPLv2 / commercially licensed, so it is not shipped: put `mongoose.c`/`mongoose.h` next to
the generated files. With Mongoose you create an `EmiHttpdRouter` (`emihttpd_router_init`), pass it to
every `*_register()` instead of an `httpd_handle_t`, then `emihttpd_listen(&router, "http://0.0.0.0:8000")`
and run `mg_mgr_poll`. Websocket sends from other threads are queued with `mg_wakeup()`. A complete
program: the `server/` folder of the tinkerlane project (`make run`, `make test`).

```
make gen                      # esp-idf-thermo/main/emi/ from thermo.emi.yml
cd esp-idf-thermo
idf.py set-target esp32s3     # any target with Wi-Fi
idf.py build flash monitor    # join Wi-Fi "emi-thermo" (12345678), then http://192.168.4.1/api/state
```

Verified: builds with ESP-IDF v5.5.5 for `esp32s3` (no warnings), and `lib/c/c_server_test.go` runs the
generated handlers on the host against a fake `esp_http_server` (routing, decoding, status codes,
websocket open/message/close, broadcast, dropping a vanished client). Not run on a board.

## What is generated, per action

Plain action (`get`/`post`/...), e.g. `set_temperature_action.h`:

```c
typedef struct SetTemperatureActionCall {
    httpd_req_t* req;
    const SetTemperatureActionPathParameters* path;   // only if the url has :params
    const SetTemperatureActionQueryParams* query;     // only if qs: is declared
    const SetTemperatureActionRequest* body;          // only if the action has `in`
    void* user;
} SetTemperatureActionCall;

typedef SetTemperatureActionResponse* (*SetTemperatureActionHandler)(const SetTemperatureActionCall*, int* status);
SetTemperatureActionBinding* set_temperature_action_register(httpd_handle_t, SetTemperatureActionHandler, void* user);
```

You return a response struct (freed for you) or `NULL`. `*status` starts at 0: untouched means 200
(or 204 for `NULL`); any code you set is used, and `NULL` with a code >= 400 sends an error body.
Bad JSON becomes 400, an oversized body 413, an url that does not match 404 - before your code runs.

Reactive action (`method: reactive`) = a websocket, e.g. `state_stream_action.h`:

```c
StateStreamActionCallbacks cb = { .on_open = ..., .on_message = ..., .on_close = ... };
StateStreamActionBinding* b = state_stream_action_register(server, &cb);

state_stream_action_broadcast(b, &update);        // typed push to every client, from any task
state_stream_action_send(b, session->conn, &update);
state_stream_action_client_count(b);
```

`on_message` receives the `in` type parsed from each text frame; `binaryType: arraybuffer` actions get
`on_binary` and `_send_binary()`/`_broadcast_binary()` instead of raw text. Path and query values are
available on the session for the whole life of the connection.

## Setup notes

- esp_http_server only: `config.uri_match_fn = emihttpd_uri_match;` on the `httpd_config_t` - it understands `/users/:id`. (Mongoose routes match `:params` on their own.)
- `CONFIG_HTTPD_WS_SUPPORT=y` for reactive actions (the header `#error`s with a hint if it is off).
- Component: `SRCS ... "emi/cjson/cJSON.c"`, `INCLUDE_DIRS "." "emi"`, `REQUIRES esp_http_server`,
  and `FETCHX_NO_CURL` (see `main/CMakeLists.txt`).
- Limits (override with `-D`): `EMIHTTPD_MAX_BODY` 8192, `EMIHTTPD_WS_MAX_MESSAGE` 16384,
  `EMIHTTPD_MAX_CLIENTS` 8 per reactive action, `EMIHTTPD_MAX_PARAMS` 8.
- Clients that disappear without a close frame are dropped on the next failed send (`on_close` fires).
- Remotes are skipped (they are calls to somebody else's server). Request headers are not decoded:
  use `call->req` with `httpd_req_get_hdr_value_str`.
- Headers only: include each action header from as many files as you like, but keep the `Binding*`
  returned by `*_register()` - it is the handle for send/broadcast.
