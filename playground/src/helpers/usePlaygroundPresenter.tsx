import { useEffect, useState } from "react";

import { useGoWasm } from "./wasm-brige";
import type { VirtualFile } from "../definitions";

export const usePlaygroundPresenter = () => {
  const { ready } = useGoWasm({ wasmPath: "emi-compiler.wasm" });
  const [value, setValue$] = useState(sampleDocument);
  const [assemblyFunction, setAssemblyFunction$] = useState("jsGenModule");

  const setAssemblyFunction = (play: string) => {
    setAssemblyFunction$(play);
    const newValue = play === "sqlQueryPredict" ? sampleSql : sampleDocument;

    setValue$(newValue);
    rerender(newValue, play);
  };

  const [features, setFeatures] = useState<string[]>([
    "nestjs",
    "typescript",
    "react",
  ]);
  const [files, setOutput] = useState<VirtualFile[]>([]);

  const rerender = (data: any, func = "") => {
    try {
      console.log("Using function:", func);
      const res = (window as any)[func](data, {
        Tags: features.join(","),
      });

      console.log(func, res);

      const formattedPromises = res.map(async (item: any) => {
        try {
          if (item.Name.endsWith(".ts") || item.Name.endsWith(".js")) {
            item.ActualScript = await (window as any).prettier.format(
              item.ActualScript,
              {
                parser: "typescript",
                plugins: (window as any).prettierPlugins,
              }
            );
          }
        } catch (e) {
          console.error("Formatting failed for item:", item.Name, e);
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
    } catch (err) {
      console.log("Error on generation:", err);
    }
  };

  const setValue = (data: string) => {
    rerender(data, assemblyFunction);
    setValue$(data);
  };

  useEffect(() => {
    if (ready) {
      setTimeout(() => {
        rerender(value, assemblyFunction);
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

const sampleSql = `SELECT 
        u.user_id as user_id,
        field(u.user_name, 'string') as UserName,
        u.user_email,
        field(COUNT(o.order_id), 'int64') AS total_orders,
        COALESCE(SUM(o.total), 0) AS total_spent,
        MAX(o.total) AS max_order,
        (
            SELECT COUNT(*) 
            FROM (
                SELECT 101 AS order_id, 1 AS user_id, 120.5 AS total
                UNION ALL
                SELECT 102, 1, 50.0
                UNION ALL
                SELECT 103, 2, 75.0
                UNION ALL
                SELECT 104, 3, 200.0
                UNION ALL
                SELECT 105, 3, 25.0
            ) o2
            WHERE o2.user_id = u.user_id AND o2.total > 50
        ) AS big_orders_count
    FROM (
        SELECT 1 AS user_id, 'Alice' AS user_name, 'alice@example.com' AS user_email
        UNION ALL
        SELECT 2, 'Bob', 'bob@example.com'
        UNION ALL
        SELECT 3, 'Carol', 'carol@example.com'
    ) u
    LEFT JOIN (
        SELECT 101 AS order_id, 1 AS user_id, 120.5 AS total
        UNION ALL
        SELECT 102, 1, 50.0
        UNION ALL
        SELECT 103, 2, 75.0
        UNION ALL
        SELECT 104, 3, 200.0
        UNION ALL
        SELECT 105, 3, 25.0
    ) o
    ON u.user_id = o.user_id
    -- WHERE  u.user_name != 'Alice'        -- filter rows before aggregation
    WHERE filter()
    GROUP BY u.user_id, u.user_name, u.user_email
    HAVING MAX(o.total) > 100        -- filter groups after aggregation
    ORDER BY total_spent DESC
    limit useval('limit')
`;
