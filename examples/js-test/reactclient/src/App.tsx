import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import "./App.css";
import { WebSocketEcho } from "./components/Websocket";

const queryClient = new QueryClient();

function AppContent() {
  return (
    <div>
      Demo
      <WebSocketEcho />
      {/* <SSESample />
      <GetPost /> */}
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
