import { useSampleSseActionMutation } from "../generated/SampleSseAction";

export function SSESample() {
  const { data } = useSampleSseActionMutation({
    onMessage: () => {
      alert(1);
    },
  });

  return (
    <div>
      <h3>SSE Messages:</h3>
      <pre>{JSON.stringify(data, null, 2)}</pre>
      {/* <ul>
        {data.map((msg, i) => (
          <li key={i}>{msg}</li>
        ))}
      </ul> */}
    </div>
  );
}
