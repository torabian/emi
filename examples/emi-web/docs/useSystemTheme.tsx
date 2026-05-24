import { useEffect, useState } from "react";

export function useSystemTheme(): "light" | "vs-dark" {
  const [theme, setTheme] = useState<"light" | "vs-dark">("light");

  useEffect(() => {
    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");

    const updateTheme = (e: MediaQueryListEvent | MediaQueryList) => {
      setTheme(e.matches ? "vs-dark" : "light");
    };

    updateTheme(mediaQuery); // initial value
    mediaQuery.addEventListener("change", updateTheme);

    return () => mediaQuery.removeEventListener("change", updateTheme);
  }, []);

  return theme;
}
