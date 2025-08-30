import { useEffect, useState } from "react";
import {
  WebSocketOrgEchoAction,
  WebSocketOrgEchoReq,
  WebSocketOrgEchoRes,
} from "../generated/WebSocketOrgEchoAction";

export function WebSocketEcho() {
  const [messages, setMessages] = useState<WebSocketOrgEchoRes[]>([]);

  useEffect(() => {
    const ws = WebSocketOrgEchoAction.Create();

    // Listen for messages
    ws.addEventListener("message", (event) => {
      console.log(1, event.data, typeof event.data);
      setMessages((prev) => [...prev, event.data]);
    });

    // Send messages when connected
    ws.addEventListener("open", () => {
      ws.send(new WebSocketOrgEchoReq({ firstName: "Hi", lastName: "Torabi" }));
      ws.send(new WebSocketOrgEchoReq({ firstName: "Hi", lastName: "Sedri" }));
    });

    // Cleanup on unmount
    return () => {
      ws.close();
    };
  }, []);

  return (
    <div>
      <h3>WebSocket Echo Messages:</h3>
      <ul>
        {messages.map((msg, i) => (
          <li key={i}>{msg.lastName}</li>
        ))}
      </ul>
    </div>
  );
}
