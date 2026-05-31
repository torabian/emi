// @ts-check
import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";
import react from "@astrojs/react";

// https://astro.build/config
export default defineConfig({
  // Production URL + base path for GitHub Pages (https://torabian.github.io/emi/)
  site: "https://torabian.github.io/emi",
  base: "/emi",

  integrations: [
    starlight({
      title: "Emi",
      tagline: "Emi compiler playground and documents",
      favicon: "/img/favicon.ico",
      logo: {
        src: "./src/assets/logo.svg",
        alt: "Emi",
      },
      customCss: ["./src/styles/custom.css"],
      social: [
        {
          icon: "github",
          label: "GitHub",
          href: "https://github.com/torabian/emi",
        },
      ],
      // "Edit this page" links point at the source file in the repo.
      editLink: {
        baseUrl: "https://github.com/torabian/emi/edit/main/examples/emi-web/",
      },
      sidebar: [
        {
          label: "Playground",
          link: "https://torabian.github.io/emi/playground",
          attrs: { target: "_blank" },
        },
        { label: "Motivation", slug: "intro" },
        { label: "Getting started", slug: "installation" },
        { label: "Emi definitions", slug: "emi-module-spec" },
        { label: "Emi Actions", slug: "emi-actions" },
        { label: "Complex Types", slug: "emi-complex-types" },
        {
          label: "Golang Compiler",
          items: [{ autogenerate: { directory: "golang" } }],
        },
        {
          label: "Javascript Compiler",
          items: [{ autogenerate: { directory: "js" } }],
        },
        {
          label: "Swift Compiler",
          items: [{ autogenerate: { directory: "swift" } }],
        },
        {
          label: "Kotlin Compiler",
          items: [{ autogenerate: { directory: "kotlin" } }],
        },
        {
          label: "Query Predict",
          items: [{ autogenerate: { directory: "query-predict" } }],
        },
        {
          label: "Documents & Export",
          items: [{ autogenerate: { directory: "exports" } }],
        },
      ],
    }),
    react(),
  ],
});
