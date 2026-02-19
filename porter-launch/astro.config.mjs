import { defineConfig } from "astro/config";
import svelte from "@astrojs/svelte";
import tailwind from "@astrojs/tailwind";

import starlight from "@astrojs/starlight";

const disableSitemap = () => ({
  name: "@astrojs/sitemap",
  hooks: {}
});

export default defineConfig({
  site: "https://porter.dev",
  integrations: [
    disableSitemap(),
    starlight({
      title: "Porter",
      description: "GitHub-native orchestration docs for agent execution workflows.",
      favicon: "/images/porter-icon.png",
      customCss: ["./src/styles/starlight-custom.css"],
      components: {
        SiteTitle: "./src/components/starlight/SiteTitle.astro",
        ThemeSelect: "./src/components/starlight/ThemeSelect.astro",
        PageSidebar: "./src/components/starlight/PageSidebar.astro"
      },
      lastUpdated: true,
      sidebar: [
        {
          label: "Quick Start",
          autogenerate: { directory: "getting-started" }
        },
        {
          label: "Concepts",
          autogenerate: { directory: "concepts" }
        },
        {
          label: "Configuration",
          autogenerate: { directory: "configuration" }
        },
        {
          label: "API",
          autogenerate: { directory: "api" }
        }
      ]
    }),
    svelte(),
    tailwind()
  ],
  output: "static",
  build: {
    inlineStylesheets: "auto"
  },
  vite: {
    build: {
      rollupOptions: {
        output: {
          manualChunks: {
            landing: ["./src/components/landing"]
          }
        }
      }
    }
  }
});
