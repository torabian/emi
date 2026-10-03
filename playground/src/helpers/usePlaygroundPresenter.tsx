import { useEffect, useRef, useState } from "react";

import { useGoWasm } from "./wasm-brige";
import type { VirtualFile } from "../definitions";
import { examplesForTarget, getExample } from "../examples";
import type { CompilerTag } from "../components/OptionsModal";

// How long to wait after the last keystroke before recompiling.
const RECOMPILE_DEBOUNCE_MS = 400;

export const usePlaygroundPresenter = () => {
  const { ready } = useGoWasm({ wasmPath: "emi-compiler.wasm" });
  const initial = examplesForTarget("jsGenModule")[0];
  const [assemblyFunction, setAssemblyFunction$] = useState(initial.target);
  const [definitionId, setDefinitionId] = useState(initial.id);
  // Options (tags) each compiler reports, loaded from the wasm module.
  const [compilerTags, setCompilerTags] = useState<
    Record<string, CompilerTag[]>
  >({});
  // Edits made to each example, so switching between them keeps your changes.
  const [drafts, setDrafts] = useState<Record<string, string>>({});

  const contentOf = (id: string) => drafts[id] ?? getExample(id).content;
  const value = contentOf(definitionId);

  const [features, setFeatures] = useState<string[]>(initial.tags);

  // Changing the compiler keeps the current definition if it still fits the new
  // target, otherwise falls back to the first example of the matching kind (yaml
  // vs sql). Options the new compiler doesn't know are dropped. Recompiling is
  // done by the effect below.
  const setAssemblyFunction = (play: string) => {
    setAssemblyFunction$(play);

    const available = examplesForTarget(play);
    const next = available.find((e) => e.id === definitionId) ?? available[0];
    setDefinitionId(next.id);

    const known = compilerTags[play];
    const sampleTags = getExample(next.id).tagsByTarget?.[play];
    if (sampleTags) {
      setFeatures(sampleTags);
    } else if (known) {
      setFeatures((prev) => prev.filter((f) => known.some((t) => t.Tag === f)));
    }
    cancelPendingRerender();
  };

  // Selecting an example switches to the compiler and options it is meant for.
  const selectDefinition = (id: string) => {
    const example = getExample(id);
    setDefinitionId(id);
    setAssemblyFunction$(example.target);
    setFeatures(example.tags);
    cancelPendingRerender();
  };

  const [files, setOutput] = useState<VirtualFile[]>([]);

  // Recompiles run asynchronously (prettier); only the latest one may publish.
  const runId = useRef(0);
  const debounceTimer = useRef<ReturnType<typeof setTimeout>>(undefined);
  const latestRerender = useRef<(data: string, func: string) => void>(() => {});

  const cancelPendingRerender = () => clearTimeout(debounceTimer.current);

  const rerender = (data: any, func = "") => {
    const currentRun = ++runId.current;
    try {
      console.log("Using function:", func);
      const res = (window as any)[func](data, {
        Tags: features.join(","),
      });

      if (!res) {
        return;
      }

      const formattedPromises = res.map(async (item: any) => {
        try {
          if (
            item.Name.endsWith(".ts") ||
            item.Name.endsWith(".js") ||
            item.Extension === ".ts" ||
            item.Extension === ".js"
          ) {
            item.ActualScript = await (window as any).prettier.format(
              item.ActualScript,
              {
                parser: "typescript",
                plugins: (window as any).prettierPlugins,
              },
            );
          }
        } catch (e) {
          console.error("Formatting failed for item:", item.Name, e);
        }
        return item;
      });

      // Wait for all formatting to finish
      Promise.all(formattedPromises).then((formattedRes) => {
        if (currentRun === runId.current) {
          setOutput(formattedRes);
        }
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
    } catch (err) {
      console.log("Error on generation:", err);
    }
  };

  // Always call the newest closure (it sees the current features) when the
  // debounce timer fires.
  latestRerender.current = rerender;

  const setValue = (data: string) => {
    setDrafts((prev) => ({ ...prev, [definitionId]: data }));
    cancelPendingRerender();
    debounceTimer.current = setTimeout(
      () => latestRerender.current(data, assemblyFunction),
      RECOMPILE_DEBOUNCE_MS,
    );
  };

  useEffect(() => cancelPendingRerender, []);

  // `ready` flips as soon as the wasm module is instantiated, but the Go program
  // only registers its globals (getCompilerTags, the generators) once it starts
  // running a moment later - so poll until the function shows up.
  useEffect(() => {
    if (!ready) return;
    let timer: ReturnType<typeof setTimeout>;
    const load = () => {
      const getTags = (window as any).getCompilerTags;
      if (typeof getTags === "function") {
        setCompilerTags(getTags());
      } else {
        timer = setTimeout(load, 50);
      }
    };
    load();
    return () => clearTimeout(timer);
  }, [ready]);

  useEffect(() => {
    if (ready) {
      const timer = setTimeout(() => {
        rerender(value, assemblyFunction);
      }, 100);
      return () => clearTimeout(timer);
    }
  }, [ready, features, assemblyFunction, definitionId]);

  return {
    files,
    value,
    definitionId,
    selectDefinition,
    setFeatures,
    ready,
    setAssemblyFunction,
    assemblyFunction,
    features,
    compilerTags,
    setValue,
  };
};
