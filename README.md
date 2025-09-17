# Emi compiler

Backend for front-end code generation library based on Emi definition files. Define http actions, dtos, entities, and many more details, and get the code boilerplate, http requests, in different languages. 

Emi generated code aims to be framework agnostic on different targets, but in the same time would provide extra
helpers based on popularity of a library.

## Live playground: https://torabian.github.io/emi 

You can try to compile Emi live here: https://torabian.github.io/emi
Emi is written in Golang, and wasm file exported to be used in javascript environment (only for generation, no runtime needed.)


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
