import { useEffect, useState } from "react";
import { SampleSseAction } from "../generated/SampleSseAction";

export function SSESample() {
  const [messages, setMessages] = useState<string[]>([]);

  useEffect(() => {
    const ac = new AbortController();

    // Fetch SSE
    SampleSseAction.Fetch(
      (event: MessageEvent) => {
        // Append each message to state
        setMessages((prev) => [...prev, event.data]);
      },
      undefined, // query paramsHow
      { signal: ac.signal }, // pass signal
      undefined
    );

    // Cancel after 5 seconds
    const timer = setTimeout(() => ac.abort(), 5000);

    return () => {
      clearTimeout(timer);
      ac.abort(); // cancel on unmount
    };
  }, []);

  return (
    <div>
      <h3>SSE Messages:</h3>
      <ul>
        {messages.map((msg, i) => (
          <li key={i}>{msg}</li>
        ))}
      </ul>
    </div>
  );
}
