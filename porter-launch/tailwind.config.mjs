/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{astro,html,js,jsx,ts,tsx,svelte,md,mdx}"],
  theme: {
    extend: {
      colors: {
        ink: {
          50: "#f5f5f5",
          100: "#e6e6e6",
          200: "#cfcfcf",
          300: "#b0b0b0",
          400: "#8a8a8a",
          500: "#6b6b6b",
          600: "#545454",
          700: "#3a3a3a",
          800: "#222222",
          900: "#121212"
        },
        accent: {
          500: "#36c5f0",
          600: "#1ba9d6"
        }
      }
    }
  },
  plugins: []
};
