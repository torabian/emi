import {
  SampleSseQueryParams,
  useSampleSse,
} from "../generated/SampleSseAction";

export function SSESample() {
  //   const { messages } = useSse(SampleSseAction.Fetch);

  const { messages } = useSampleSse({ qs: new SampleSseQueryParams() });

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
