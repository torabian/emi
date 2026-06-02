import "./database-bridge.js";
import { WebSocketWasm } from "./WebSocketWasm.js";
import { ChatAction } from "./gen/ChatAction.js";

import { SubstringAction, SubstringActionRes } from "./gen/SubstringAction.js";
console.log(1, SubstringAction);

const go = new Go();
const r = await WebAssembly.instantiateStreaming(
  fetch("in-browser-server.wasm"),
  go.importObject,
);
// Do NOT await — main() blocks on select{} so the Go runtime stays alive
// and the exposed callback keeps working.
go.run(r.instance);

// window.handleWasmRequest is registered inside Go's main(), after the async
// DBTest finishes — so wait for it before wiring the UI.
async function waitForHandler() {
  while (typeof window.handleWasmRequest !== "function") {
    await new Promise((res) => setTimeout(res, 20));
  }
}

// Drop-in fetch replacement backed by the in-browser Go net/http router.
// The Go side speaks real HTTP semantics (ServeMux + httptest), so we can
// hand back a genuine Response object — callers can't tell it isn't a
// network server.
async function localFetch(url, opts = {}) {
  // handleWasmRequest returns a Promise — the Go side runs the request on a
  // goroutine so DB queries (which await JS promises) can resolve instead of
  // deadlocking a synchronous JS→Go call.
  const raw = await window.handleWasmRequest(
    opts.method || "GET",
    url,
    opts.body || "",
    JSON.stringify(opts.headers || {}),
  );
  const { status, headers, body } = JSON.parse(raw);
  const h = new Headers();
  for (const [k, vs] of Object.entries(headers || {})) {
    for (const v of vs) h.append(k, v);
  }
  return new Response(body, { status, headers: h });
}
// Expose it so you can also poke at it from the devtools console.
window.localFetch = localFetch;

await waitForHandler();

const $ = (id) => document.getElementById(id);
$("status").textContent = "ready — window.handleWasmRequest is live";
$("call").disabled = false;

$("call").addEventListener("click", async () => {
  const payload = {
    input: $("input").value,
    start: Number($("start").value),
    end: Number($("end").value),
  };
  const res = await localFetch("/", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  const text = await res.text();

  console.log(7, text, new SubstringActionRes(text));

  $("out").textContent =
    `POST / -> ${res.status} ${res.headers.get("Content-Type") || ""}\n` +
    `request : ${JSON.stringify(payload)}\n` +
    `response: ${new SubstringActionRes(text)}`;
});

// Fire one call automatically so you can see it working on load.
$("call").click();

async function waitForLocalFetch() {
  while (typeof window.localFetch !== "function") {
    await new Promise((res) => setTimeout(res, 20));
  }
}

async function refreshUsers() {
  const res = await window.localFetch("/users", { method: "GET" });
  const data = JSON.parse(await res.text());
  const users = data.users || [];
  $("usersOut").textContent =
    `GET /users -> ${res.status} (${users.length} row(s))\n` +
    users
      .map(
        (u) => `#${u.id}  ${u.firstName} ${u.lastName}  (born ${u.birthDate})`,
      )
      .join("\n");
}

await waitForLocalFetch();

$("createUser").disabled = false;
$("deleteUser").disabled = false;
$("listUsers").disabled = false;

$("createUser").addEventListener("click", async () => {
  const payload = {
    firstName: $("firstName").value,
    lastName: $("lastName").value,
    birthDate: $("birthDate").value,
  };
  const res = await window.localFetch("/users", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  console.log("POST /users ->", res.status, await res.clone().text());
  await refreshUsers();
});

$("deleteUser").addEventListener("click", async () => {
  const res = await window.localFetch("/users", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id: Number($("deleteId").value) }),
  });
  console.log("DELETE /users ->", res.status, await res.clone().text());
  await refreshUsers();
});

$("listUsers").addEventListener("click", refreshUsers);

// --- Reactive chat over WebSocketWasm ------------------------------------
// Note how this block uses the WebSocket API verbatim. Replace WebSocketWasm
// with the native WebSocket (and point at a real ws:// server) and it behaves
// identically.
let chat = null;

function setChatState(label, connected) {
  $("chatState").textContent = label;
  $("chatConnect").disabled = connected;
  $("chatDisconnect").disabled = !connected;
  $("chatSend").disabled = !connected;
}

function chatLog(line) {
  $("chatOut").textContent += line + "\n";
}

$("chatConnect").addEventListener("click", () => {
  chatLog("connecting…");
  // Drive the GENERATED SDK client (ChatAction) — but tell it to use our
  // WebSocketWasm transport instead of the browser's native WebSocket. Drop the
  // SocketClass option and it would hit a real ws:// server unchanged.
  chat = ChatAction.Create(undefined, undefined, {
    SocketClass: WebSocketWasm,
  });

  chat.onopen = () => {
    chatLog("● open");
    setChatState("open", true);
  };
  chat.onmessage = (e) => chatLog("← " + e.data);
  chat.onerror = (e) => chatLog("✕ error: " + (e.message || "unknown"));
  chat.onclose = (e) => {
    chatLog("○ close" + (e.reason ? " (" + e.reason + ")" : ""));
    setChatState("closed", false);
    chat = null;
  };
});

$("chatSend").addEventListener("click", () => {
  if (!chat) return;
  const text = $("chatInput").value;
  chat.send(text);
  chatLog("→ " + text);
});

$("chatDisconnect").addEventListener("click", () => chat?.close());
