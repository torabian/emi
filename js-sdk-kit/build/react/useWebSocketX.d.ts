import { WebSocketX } from "../common/WebSocketX";
export type UseWebSocketResult<TSend, TRecv> = {
    send: (msg: TSend) => void;
    close: () => void;
    restart: () => void;
    socket?: WebSocketX<TSend, TRecv>;
} & UseWebSocketState<TRecv>;
interface UseWebSocketState<TRecv> {
    messages: TRecv[];
    isOpen: boolean;
    error?: Event | null;
}
export declare function useWebSocketX<TSend = any, TRecv = any, TQuery = any>(fn: (overrideUrl?: string | undefined, qs?: TQuery | undefined) => WebSocketX<TSend, TRecv>): UseWebSocketResult<TSend, TRecv>;
export {};
