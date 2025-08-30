import { useEffect, useState } from "react";

import { useGoWasm } from "./wasm-brige";
import type { VirtualFile } from "../definitions";

export const usePlaygroundPresenter = () => {
  const { ready } = useGoWasm({ wasmPath: "emi-compiler.wasm" });
  const [value, setValue$] = useState(sampleDocument);
  const [assemblyFunction, setAssemblyFunction] = useState("jsGenAction");
  const [features, setFeatures] = useState<string[]>([
    "nestjs",
    "axios",
    "typescript",
    "axiosbundle",
    "angular",
    "react",
  ]);
  const [files, setOutput] = useState<VirtualFile[]>([]);

  const rerender = (data: any) => {
    try {
      const res = (window as any)[assemblyFunction](data, {
        Tags: features.join(","),
      });

      const formattedPromises = res.map(async (item: any) => {
        try {
          item.ActualScript = await (window as any).prettier.format(
            item.ActualScript,
            {
              parser: "typescript",
              plugins: (window as any).prettierPlugins,
            }
          );
        } catch (e) {
          console.error("Formatting failed for item:", item, e);
        }
        return item;
      });

      // Wait for all formatting to finish
      Promise.all(formattedPromises).then((formattedRes) => {
        setOutput(formattedRes);
      });

      // for (const item of res) {
      //   // item.ActualScript is string, call it with prettier, which returns promise,
      //   // and when all of them is resoled, put it back to item.ActionScript
      // }

      // // This is a promise. Run it for all files
      // // (window as any).prettier
      // // .format("var x = 10", {
      // //   parser: "typescript",
      // //   plugins: (window as any).prettierPlugins,
      // // })

      // setOutput(res);
    } catch (err) {}
  };

  const setValue = (data: string) => {
    rerender(data);
    setValue$(data);
  };

  useEffect(() => {
    if (ready) {
      setTimeout(() => {
        rerender(value);
      }, 100);
    }
  }, [ready, features]);

  return {
    files,
    value,
    setFeatures,
    ready,
    setAssemblyFunction,
    assemblyFunction,
    features,
    setValue,
  };
};

const sampleDocument = `name: sampleModule
actions:
  - name: getSinglePost
    url: https://jsonplaceholder.typicode.com/posts/1
    cliName: get-single-post
    method: post
    description: Get's an specific post from the endpoint
    in:
      headers:
        - name: accept-language
          type: string
    out:
      headers:
        - name: content-type
          type: string
      fields:
        - name: userId
          type: int64?
        - name: id
          type: int64
        - name: title
          type: string
        - name: body
          type: string
        - name: user
          type: object
          fields:
          - name: firstName
            type: string?
          - name: age
            type: int64
        - name: histories
          type: array
          fields:
          - name: firstName
            type: string?
          - name: age
            type: int64
          - name: info
            type: object
            fields:
            - name: memorySize
              type: int64
        `;
