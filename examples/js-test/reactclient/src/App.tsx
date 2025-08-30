import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import "./App.css";
import {
  GetSinglePostQueryParams,
  useGetSinglePost,
} from "./generated/GetSinglePostAction";
import { FetchSampleSseAction } from "./generated/SampleSseAction";

const queryClient = new QueryClient();

function AppContent() {
  const [id, setId] = useState(42);

  useEffect(() => {
    const ac = new AbortController(); // create signal

    FetchSampleSseAction.Fetch(
      undefined, // query params
      { signal: ac.signal }, // pass the signal in init
      undefined,
      (event: MessageEvent) => {
        console.log(event.data); // onMessage
      }
    ).then(() => {
      console.log("SSE completed");
    });

    // // Cancel after 5 seconds
    // const timer = setTimeout(() => ac.abort(), 5000);

    return () => {
      // clearTimeout(timer); // cleanup timer
      ac.abort(); // cancel SSE on unmount
    };
  }, []);

  // useEffect(() => {
  //   const interval = setInterval(() => {
  //     setId((prev) => prev + 1);
  //   }, 2000);
  //   return () => clearInterval(interval);
  // }, []);

  const { data, isLoading, error } = useGetSinglePost({
    params: { id },
    qs: new GetSinglePostQueryParams().set("pageSize", 100),
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {(error as any).message}</div>;

  console.log(typeof data);
  console.log(data);

  return (
    <pre style={{ whiteSpace: "pre-wrap", wordBreak: "break-word" }}>
      {JSON.stringify(data, null, 2)}
      {typeof data}
    </pre>
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
