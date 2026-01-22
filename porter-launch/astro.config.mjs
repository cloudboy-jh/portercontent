import { defineConfig } from "astro/config";
import svelte from "@astrojs/svelte";
import tailwind from "@astrojs/tailwind";
import sitemap from "@astrojs/sitemap";

export default defineConfig({
  site: "https://porter.dev",
  integrations: [svelte(), tailwind(), sitemap()],
  output: "static",
  build: {
    inlineStylesheets: "auto"
  },
  vite: {
    build: {
      rollupOptions: {
        output: {
          manualChunks: {
            landing: ["./src/components/landing"],
            docs: ["./src/components/docs"]
          }
        }
      }
    }
  }
});
