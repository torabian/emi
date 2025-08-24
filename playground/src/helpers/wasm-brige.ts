import { useEffect, useState, useCallback } from "react";

type GoWasmConfig = {
  wasmPath: string; // e.g. "/emi-compiler.wasm"
};

export function useGoWasm(config: GoWasmConfig) {
  const [ready, setReady] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const [mod, setMod] = useState<WebAssembly.Module | null>(null);
  const [_, setInst] = useState<WebAssembly.Instance | null>(null);

  useEffect(() => {
    // polyfill instantiateStreaming if not supported
    if (!(WebAssembly as any).instantiateStreaming) {
      (WebAssembly as any).instantiateStreaming = async (
        resp: Response | Promise<Response>,
        importObject: WebAssembly.Imports
      ) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
      };
    }

    let go: any;

    async function init() {
      try {
        // @ts-ignore: wasm_exec.js defines Go globally
        go = new Go();

        const result = await (WebAssembly as any).instantiateStreaming(
          fetch(config.wasmPath),
          go.importObject
        );

        setMod(result.module);
        setInst(result.instance);
        setReady(true);
      } catch (err: any) {
        setError(err);
      }
    }

    init();
  }, [config.wasmPath]);

  const run = useCallback(async () => {
    if (!mod) return;
    // @ts-ignore
    const go = new Go();
    const instance: any = await WebAssembly.instantiate(mod, go.importObject);
    await go.run(instance.instance || instance); // different browsers return shape
    setInst(instance.instance || instance);
  }, [mod]);

  useEffect(() => {
    if (ready) {
      run();
    }
  }, [ready]);

  return { ready, error, run };
}
