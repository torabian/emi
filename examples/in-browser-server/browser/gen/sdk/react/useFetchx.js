import { createContext, useContext } from "react";
const FetchxContextReact = createContext(null);
export const FetchxProvider = FetchxContextReact.Provider;
export function useFetchxContext() {
    return useContext(FetchxContextReact);
}
