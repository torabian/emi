import { createContext, useContext } from "react";
import { FetchxContext } from "../common/fetchx";

const FetchxContextReact = createContext<FetchxContext<
  unknown,
  unknown
> | null>(null);

export const FetchxProvider = FetchxContextReact.Provider;

export function useFetchxContext(): FetchxContext<unknown, unknown> | null {
  return useContext(FetchxContextReact);
}
