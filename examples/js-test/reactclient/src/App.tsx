import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import "./App.css";
import { SSESample } from "./components/SSE";
import {
  GetSinglePostQueryParams,
  useGetSinglePost,
} from "./generated/GetSinglePostAction";
import { WebSocketEcho } from "./components/Websocket";

const queryClient = new QueryClient();

function AppContent() {
  const { data, isLoading, error } = useGetSinglePost({
    params: { id: 10 },
    qs: new GetSinglePostQueryParams().set("pageSize", 100),
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {(error as any).message}</div>;

  return (
    <div>
      <WebSocketEcho />
      <SSESample />
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

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <AppContent />
    </QueryClientProvider>
  );
}

export default App;
