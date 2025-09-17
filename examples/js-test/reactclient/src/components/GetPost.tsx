import {
  GetSinglePostActionRes,
  useGetSinglePostActionQuery,
} from "../generated/GetSinglePostAction";

export function GetPost() {
  const { data, isLoading, error, isCompleted, result } =
    useGetSinglePostActionQuery({
      params: { id: 15 },
    });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {(error as any).message}</div>;

  return (
    <div>
      <pre
        style={{
          whiteSpace: "pre-wrap",
          wordBreak: "break-word",
          textAlign: "left",
        }}
      >
        Is instance: {data instanceof GetSinglePostActionRes ? "Yes" : "No"}
      </pre>
    </div>
  );
}
