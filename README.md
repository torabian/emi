# Emi - Backend-for-Frontend with automatic SDK generation.

Live preview: https://torabian.github.io/emi/playground

Backend for front-end code generation library based on Emi definition files. Define http actions, dtos, enums, and
other sharable details with a single yaml definition, and then compile multiple
targets out of the same definition. Golden tool for creating SDKs which work
with an API system.

Main focus is on **JavaScript, TypeScript, Golang** environment, also supports a basic
version of **Swift** and **Kotlin** which is more limited (PR Welcome)

Code written in Emi aims to be framework agnostic, but still there are following
libraries used:

- React.js code flag. Emi generates react.js, and tanstack react query code,
upon flag `react` provided to `js` compiler. You can change the library import
location, or version, if needed.
- On Golang side, code generated uses `Gin` and possibly `urfave/cli` for giving
context on cli status, which are both very well known and mature libraries.
- On js side, `qs` has been used for parsing query strings.
- Small amount of helpers are added for next.js, upon providing `--tags nextjs` for
js compiler, which is a famous backend framework on node.js. Helps
to use the classes generated directly in req, headers arguments.

Notes:
- Golang and JavaScript(Ts) are having an special focus.
- Emi generates most type-safe javascript code, which ensures typesafety through
class generation of dtos which is very unique feature.
- Kotlin and Swift compilers require more work, you can modify them and open a pull
request.

<img src="./emi-flowchart.png" style="max-width: 600px" />

## Live playground: https://torabian.github.io/emi 

You can try to compile Emi live here: https://torabian.github.io/emi/playground
Emi is written in Golang, and wasm file exported to be used in javascript environment (only for generation, no runtime needed.)


## Feature support in a glance.

| Feature / Language        | Golang | JavaScript | JavaScript (TS) | JavaScript (Node.js) | Kotlin | Swift | Notes |
|----------------------------|--------|------------|----------------|---------------------|--------|-------|-------|
| DTO Generation             | ✅     | ✅         | ✅             | ✅                  | ✅     | ✅    | Supported in all languages |
| HTTP Actions               | ✅     | ✅         | ✅             | ✅                  | WIP     | WIP    | Works with HTTP client libraries |
| Command Line               | ✅     | ❌         | ❌             | ❌                  | ❌     | ❌    | Only Golang has CLI support currently |


## Javascript/TypeScript generation

Emi javascript generator can be used for both typescript or javascript without need of recompiling to javascript
via external tools. It would cover:

- Pure vanilla javascript, with classes, and fetch api. (always enabled, not tags needed)
- Nest.js extra decorators (--tags `nestjs`)
- React query tools (--tags `react`) generates mutations and useQuery tools.

Emi compiler generates code which is compatible with javascript standard library, such as UrlSearchParams, Headers,
and would convert res/req into very typesafe matter by creating classes, and introducing getters and setters.

Example fetch created via emi compiler, which is 100% compatible with `fetch`:

```ts

export class FetchGetSinglePostAction {
  static URL = 'https://jsonplaceholder.typicode.com/posts/:id';
  static NewUrl = (
    params: FetchGetSinglePostActionPathParameter,
    qs?: GetSinglePostQueryParams,
  ) => buildUrl(FetchGetSinglePostAction.URL, params, qs);
  static Method = 'get';
  static Fetch = async (
    params: FetchGetSinglePostActionPathParameter,
    qs?: GetSinglePostQueryParams,
    init?: TypedRequestInit<GetSinglePostRes, GetSinglePostReqHeaders>,
    overrideUrl?: string,
  ) => {
    const res = await fetchx<
      GetSinglePostRes,
      unknown,
      GetSinglePostReqHeaders
    >(overrideUrl ?? FetchGetSinglePostAction.NewUrl(params, qs), {
      method: FetchGetSinglePostAction.Method,
      ...(init || {}),
    });
    const result = await res.json();
    res.result = new GetSinglePostRes(result);
    return res;
  };
}
```


## Emi syntax

Emi is a yaml file, and here you can find the complete schema:
- https://github.com/torabian/emi/blob/main/playground/public/emi-module-spec.json

You can use the definition with Redhat yaml plugin in the vscode.
