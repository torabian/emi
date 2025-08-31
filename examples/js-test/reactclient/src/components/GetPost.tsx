import {
  GetSinglePostQueryParams,
  useGetSinglePost,
} from "../generated/GetSinglePostAction";

export function GetPost() {
  const { data, isLoading, error } = useGetSinglePost({
    params: { id: 10 },
    qs: new GetSinglePostQueryParams().set("pageSize", 100),
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
        {JSON.stringify(data, null, 2)}
      </pre>
    </div>
  );
}
