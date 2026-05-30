# Website

This website is built with [Astro](https://astro.build/) and the
[Starlight](https://starlight.astro.build/) documentation theme.

## Installation

```bash
npm install
```

## Local Development

```bash
npm run dev
```

Starts a local dev server at `http://localhost:4321/emi/`. Most changes are
reflected live without restarting the server.

## Build

```bash
npm run build
```

Generates the static site into the `dist/` directory, which can be served by
any static host.

```bash
npm run preview
```

Locally previews the production build.

## Project structure

- `src/content/docs/` — documentation pages (`.md` / `.mdx`).
- `src/components/` — React components used as islands inside MDX (e.g. the
  object benchmark).
- `src/styles/custom.css` — brand theming via Starlight CSS variables.
- `public/` — static assets served at the site root.
- `astro.config.mjs` — site config, including the GitHub Pages `base` (`/emi`)
  and the sidebar layout.

## Deployment

The site is configured for GitHub Pages under the `/emi/` base path. Build with
`npm run build` and publish the `dist/` directory (e.g. via a GitHub Actions
Pages workflow).
