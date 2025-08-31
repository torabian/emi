import { SampleSseAction } from "../generated/SampleSseAction";
import { useSse } from "../generated/sdk/react";

export function SSESample() {
  const { messages } = useSse(SampleSseAction.Fetch);

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
