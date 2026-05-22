import React, { useState } from "react";

class TestClass {
  private str: string = "";
  private num: number = 0;
  private obj: Record<string, unknown> = {};
  private arr: any[] = [];

  setStr(v: any) {
    this.str = typeof v === "string" ? v : String(v);
  }
  getStr() {
    return this.str;
  }

  setNum(v: any) {
    if (typeof v === "number") {
      this.num = v;
    } else {
      const parsed = Number(v);
      this.num = isNaN(parsed) ? 0 : parsed;
    }
  }
  getNum() {
    return this.num;
  }

  setObj(v: any) {
    this.obj = typeof v === "object" && !Array.isArray(v) ? v : { value: v };
  }
  getObj() {
    return this.obj;
  }

  setArr(v: any) {
    this.arr = Array.isArray(v) ? v : [v];
  }
  getArr() {
    return this.arr;
  }
}

export default function BenchmarkTest() {
  const [iterations, setIterations] = useState(100_000);
  const [results, setResults] = useState<{
    simple: number;
    classBased: number;
  } | null>(null);

  const runBenchmark = () => {
    // --- simple object ---
    const obj: any = {};
    let t1 = performance.now();
    for (let i = 0; i < iterations; i++) {
      obj.str = "Hello";
      obj.num = 42;
      obj.obj = { k: "v" };
      obj.arr = [1, 2, 3];
      const x1 = obj.str;
      const x2 = obj.num;
      const x3 = obj.obj;
      const x4 = obj.arr;
    }
    let t2 = performance.now();
    const simple = t2 - t1;

    // --- class object ---
    const o = new TestClass();
    t1 = performance.now();
    for (let i = 0; i < iterations; i++) {
      o.setStr("Hello");
      o.setNum(42);
      o.setObj({ k: "v" });
      o.setArr([1, 2, 3]);
      const x1 = o.getStr();
      const x2 = o.getNum();
      const x3 = o.getObj();
      const x4 = o.getArr();
    }
    t2 = performance.now();
    const classBased = t2 - t1;

    setResults({ simple, classBased });
  };

  const codeString = `// simple object
const obj: any = {};
for (let i = 0; i < ${iterations}; i++) {
  obj.str = "Hello";
  obj.num = 42;
  obj.obj = { k: "v" };
  obj.arr = [1, 2, 3];
  const x1 = obj.str;
  const x2 = obj.num;
  const x3 = obj.obj;
  const x4 = obj.arr;
}

// class object
const o = new TestClass();
for (let i = 0; i < ${iterations}; i++) {
  o.setStr("Hello");
  o.setNum(42);
  o.setObj({ k: "v" });
  o.setArr([1, 2, 3]);
  const x1 = o.getStr();
  const x2 = o.getNum();
  const x3 = o.getObj();
  const x4 = o.getArr();
}`;

  return (
    <div className="p-4 space-y-4">
      <div className="flex gap-2 items-center">
        <input
          type="number"
          className="border p-2 rounded"
          value={iterations}
          onChange={(e) => setIterations(Number(e.target.value))}
        />
        <button
          onClick={runBenchmark}
          className="px-4 py-2 bg-blue-600 text-white rounded"
        >
          Run
        </button>
      </div>

      {results && (
        <div className="space-y-2">
          <div>Object property: {results.simple.toFixed(2)} ms</div>
          <div>Class get/set: {results.classBased.toFixed(2)} ms</div>
        </div>
      )}

      <pre className="bg-gray-100 p-3 rounded overflow-x-auto text-sm">
        {codeString}
      </pre>
    </div>
  );
}
