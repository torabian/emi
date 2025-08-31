import { useEffect } from "react";
import {
  WebSocketOrgEchoAction,
  WebSocketOrgEchoReq,
} from "../generated/WebSocketOrgEchoAction";
import { useWebSocketX } from "../generated/sdk/react";

export function WebSocketEcho() {
  const { messages, isOpen, send } = useWebSocketX(
    WebSocketOrgEchoAction.Create
  );

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
