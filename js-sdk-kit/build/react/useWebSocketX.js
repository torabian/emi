import { useEffect, useRef, useState, useCallback } from "react";
export function useWebSocketX(fn) {
    const socketRef = useRef(null);
    const [state, setState] = useState({
        messages: [],
        isOpen: false,
        error: undefined,
    });
    const create = useCallback(() => {
        const ws = fn();
        socketRef.current = ws;
        setState({
            messages: [],
            error: undefined,
            isOpen: ws.readyState === ws.OPEN,
        });
        ws.addEventListener("message", (ev) => {
            setState((prev) => {
                return {
                    ...prev,
                    messages: [...prev.messages, ev.data],
                };
            });
        });
        ws.addEventListener("error", (ev) => {
            setState((prev) => {
                return {
                    ...prev,
                    error: ev,
                };
            });
        });
        ws.addEventListener("open", () => {
            setState((prev) => {
                return {
                    ...prev,
                    isOpen: true,
                };
            });
        });
        ws.addEventListener("close", () => {
            setState((prev) => {
                return {
                    ...prev,
                    isOpen: false,
                };
            });
        });
        return ws;
    }, []);
    useEffect(() => {
        const ws = create();
        return () => ws.close();
    }, [create]);
    const send = (msg) => {
        var _a;
        (_a = socketRef.current) === null || _a === void 0 ? void 0 : _a.send(msg);
    };
    const close = () => {
        var _a;
        (_a = socketRef.current) === null || _a === void 0 ? void 0 : _a.close();
    };
    const restart = () => {
        close();
        create();
    };
    return {
        ...state,
        send,
        close,
        restart,
        socket: socketRef.current || undefined,
    };
}
