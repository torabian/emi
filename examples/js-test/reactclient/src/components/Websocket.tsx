import { useEffect } from "react";
import {
  useWebSocketOrgEcho,
  WebSocketOrgEchoReq,
} from "../generated/WebSocketOrgEchoAction";

export function WebSocketEcho() {
  const { messages, isOpen, send } = useWebSocketOrgEcho({});

  useEffect(() => {
    if (isOpen) {
      const x = setInterval(() => {
        send(new WebSocketOrgEchoReq({}).setLastName("Sara"));
      }, 1000);
      return () => clearInterval(x);
    }
  }, [isOpen]);

  return (
    <div>
      2
      <ul>
        <span>{isOpen ? "Open" : "Close"}</span>
        {messages.map((msg, i) => (
          <li key={i}>{msg.lastName}</li>
        ))}
      </ul>
    </div>
  );
}
