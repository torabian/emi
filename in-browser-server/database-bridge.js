// This code will inject necessary code for pglite to become available
// on golang side.

import { PGlite } from "./pglite/index.js";

const db = await PGlite.create("idb://my-pgdata");
console.log("[pglite] ready");

window.queryDatabase = async function (query, args) {
  console.log("Incoming db:", query, args);
  try {
    if (query === "COPY") {
      const payload = JSON.parse(args);
      const cols = payload.columns.map((c) => `"${c}"`).join(", ");
      const table = Array.isArray(payload.table)
        ? payload.table.map((t) => `"${t}"`).join(".")
        : `"${payload.table}"`;

      let affected = 0;
      for (const row of payload.rows || []) {
        const placeholders = row.map((_, i) => `$${i + 1}`).join(", ");
        const res = await db.query(
          `INSERT INTO ${table} (${cols}) VALUES (${placeholders})`,
          row,
        );
        affected += res.affectedRows ?? 0;
      }

      return JSON.stringify({
        rows: [],
        fields: [],
        affectedRows: affected,
        error: null,
      });
    }

    const params = Array.isArray(args) ? args : [];
    const res = await db.query(query, params);

    return JSON.stringify({
      rows: res.rows ?? [],
      fields: res.fields ?? [],
      affectedRows: res.affectedRows ?? 0,
      error: null,
    });
  } catch (e) {
    return JSON.stringify({
      rows: [],
      fields: [],
      affectedRows: 0,
      error: {
        Message: String(e && e.message ? e.message : e),
        Code: (e && e.code) || "",
        Severity: (e && e.severity) || "ERROR",
      },
    });
  }
};
