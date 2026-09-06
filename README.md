# Emi Compiler

### One YAML API definition → type-safe Go, TypeScript, Swift, Kotlin, Python, Dart, C#, Java, PHP, C & C++ SDKs

**Live playground:** https://torabian.github.io/emi/playground — try Emi in the browser, no install needed.

Emi is a code generator: you describe your API once — dtos, entities, actions, config —
in a single yaml file, and Emi compiles that one definition into working, type-safe code
for multiple languages, so the backend and every client SDK are always generated from
the same source of truth and never drift out of sync with each other.

Beyond basic DTO/action generation, Emi's goal is to cover the topics that usually get
bolted on by hand afterwards: **reactive programming** (real WebSocket/SSE actions,
typed on both ends), **nullability** (nullable-aware types across every target),
**envelopes** (a consistent response-wrapper shape, e.g. Google's JSON style guide),
**translations** (string-resource generation across languages), a first-class **CLI**
(Golang actions double as `urfave/cli` commands for free), and **WASM** (the same
compiler runs in the browser, no server round trip).

<img src="./emi-languages.png" style="max-width: 600px" />

## Quick start

**1. Install Emi**

```bash
go install github.com/torabian/emi/cmd/emi@latest
```

(or download a prebuilt binary from the [releases](https://github.com/torabian/emi/releases) page)

**2. Create the sample yaml**

Save this as `user.emi.yml`:

```yaml
dtos:
  - name: user
    fields:
      - name: id
        type: string
      - name: email
        type: string
      - name: age
        type: int?
```

**3. Compile it to multiple languages**

```bash
emi go     --path user.emi.yml --output ./out/go
emi js     --path user.emi.yml --output ./out/js --tags typescript
emi swift  --path user.emi.yml --output ./out/swift
emi python --path user.emi.yml --output ./out/python
```

Each command reruns the same `user.emi.yml` through a different target compiler and
writes matching, type-safe code to its own output folder. You can also list every
target once, under `targets:`, in the yaml itself and compile them all in one shot
with `emi compile --path user.emi.yml`.

## Language targets & features

| Target                      | What you get                                                                 |
| ---------------------------- | ----------------------------------------------------------------------------- |
| **Golang**                  | Full server + client: DTOs, `Gin` HTTP handlers, `urfave/cli` CLI, GORM entities, reactive/WebSocket support. The only target that generates a server. |
| **JavaScript / TypeScript** | Typed client SDK (vanilla or React + TanStack Query hooks), typed `fetch` layer, JSDoc typedefs for plain JS, reactive WebSocket/SSE hooks, NestJS decorators. Runs in the browser or Node. |
| **Swift**                   | `Codable` DTO structs, typed header structs, typed WebSocket client for reactive actions. HTTP actions are WIP. |
| **Kotlin**                  | Typed DTO/action client code. Reactive/WebSocket not yet supported. |
| **Python**                  | `dataclasses`-based DTOs, `httpx` HTTP client (sync or async), SSE for reactive actions. |
| **Dart**                    | DTO classes with hand-generated `toJson`/`fromJson`, `package:http` client, SSE for reactive actions. |
| **C#**                      | `System.Text.Json`-based DTOs, `HttpClient` transport, SSE via `IAsyncEnumerable<string>`. |
| **Java**                    | Jackson-based DTOs, `java.net.http.HttpClient` transport, SSE for reactive actions. |
| **PHP**                     | Reflection-based DTO hydration, `curl` HTTP transport, SSE for reactive actions. |
| **C**                       | Vendored `cJSON` (de)serialization, `libcurl` transport. No nullable value types without pointers (documented scope limit). |
| **C++ (generic)**           | Portable ISO C++17 DTOs/actions for desktop, ESP-IDF, and Arduino; `IEmiHttpTransport` seam; real RFC 6455 WebSocket client for reactive actions. |
| **C++ (Unreal Engine)**     | `USTRUCT`/`UPROPERTY` DTOs usable from Blueprint, Unreal's own JSON reflection, async HTTP via `FHttpModule`, real WebSocket via `IWebSocket`. |

Every target except Golang is client-only — no server, no CLI, and no entity/GORM
persistence is ever generated for them. Shared across all targets: nullable-aware
types, nested object/array/map fields, `one`/`collection` relations, and enums.

---

Documentation: https://torabian.github.io/emi
