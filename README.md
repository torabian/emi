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
- Axios helper and Axios Bundle (--tags `axios,axiosbundle`)
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

### Javascript features

Here is a list of features which will be generated via js/ts compiler:

- [x] Type of action options, to determine the values as query
- [x] Url parameters type-safety
- [x] FetchMeta data class, containing the method, url, url generator, fetch and axios helper
- [x] Generate the object into typescript `type` only
- [x] Generate Request, and Response classes
- [x] Generate type-safe query strings (?param=1) and support deep nested
- [x] Nest.js decorators to use query, req, and headers directly in actions
- [x] React.js useQuery
- [x] React.js useMutation
- [x] Package.json generation
- [x] Playground eslint
- [x] Constructor for javascript needs to use it’s own functions
- [x] Make the partial compiler functions public
- [x] Module3 complete documentation to write.
- [x] Typesafe reactive methods, for WebSocket access.
- [x] Publish on npm and make available on npx
- [x] Generate the DTO and Entities standalone in emi, as well as remotes
- [x] Determine on the path selector, so in the options user can select details needed to be generated
- [x] Static fields for the common object to access fields via json path
- [x] Document static fields
- [ ] Make a complete demo of the code generation as a slideshow and video.

## Emi Golang

Emi for golang will be based on Fireback module3 files, it's a goal to use only emi for generating golang placeholder code.

## Emi Swift

Emi for swift and swiftui will be a similar structure to javascript compiler, to make writing swift features easier.

## Emi Kotlin

Emi kotlin wil be a generated client tools (and some shared, to be used in springboot) for writing mostly on Android environment.

## Emi syntax

Emi is a yaml file, and here you can find the complete schema:
- https://github.com/torabian/emi/blob/main/playground/public/emi-definitions.json

You can use the definition with Redhat yaml plugin in the vscode.
