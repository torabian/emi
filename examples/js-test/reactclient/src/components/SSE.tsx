import {
  SampleSseActionQueryParams,
  useSampleSseAction,
} from "../generated/SampleSseAction";

export function SSESample() {
  const { messages } = useSampleSseAction({
    qs: new SampleSseActionQueryParams(),
  });

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
