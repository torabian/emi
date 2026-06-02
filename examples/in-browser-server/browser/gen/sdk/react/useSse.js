import { useEffect, useRef, useState } from "react";
export function useSse(fetchFn, props) {
    const acRef = useRef(null);
    const [state, setState] = useState({
        messages: [],
    });
    const create = () => {
        const ac = new AbortController();
        acRef.current = ac;
        setState({ messages: [], error: null }); // reset on new stream
        fetchFn((ev) => {
            setState((value) => {
                const next = ev.data;
                if ((value.messages || []).includes(next))
                    return value;
                return { ...value, messages: [...(value.messages || []), next] };
            });
        }, props === null || props === void 0 ? void 0 : props.qs, { ...((props === null || props === void 0 ? void 0 : props.init) || {}), signal: ac.signal }, props === null || props === void 0 ? void 0 : props.overrideUrl).catch((err) => setState((v) => ({ ...v, error: err })));
        return () => { var _a; return (_a = acRef.current) === null || _a === void 0 ? void 0 : _a.abort(); };
    };
    useEffect(() => {
        return create();
    }, []);
    const cancel = () => {
        var _a;
        (_a = acRef.current) === null || _a === void 0 ? void 0 : _a.abort();
    };
    const restart = () => {
        cancel();
        create();
    };
    return { ...state, cancel, restart, messages: state.messages };
}
