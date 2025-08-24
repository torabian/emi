import { useEffect, useState } from "react";

import { useGoWasm } from "./wasm-brige";
import type { VirtualFile } from "../definitions";

export const usePlaygroundPresenter = () => {
  const { ready } = useGoWasm({ wasmPath: "emi-compiler.wasm" });
  const [value, setValue$] = useState(sampleDocument);
  const [assemblyFunction, setAssemblyFunction] = useState("js_module");
  const [features, setFeatures] = useState<string[]>([
    "nestjs",
    "axios",
    "typescript",
  ]);
  const [files, setOutput] = useState<VirtualFile[]>([]);

  const rerender = (data: any) => {
    try {
      const res = (window as any)[assemblyFunction](data, {
        Tags: features.join(","),
      });
      setOutput(res);
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
    method: get
    description: Get's an specific post from the endpoint
    out:
      fields:
        - name: userId
          type: int64
        - name: id
          type: int64
        - name: title
          type: string
        - name: body
          type: string`;
