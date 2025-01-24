import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
  darkMode: "class", // Enables class-based dark mode
  theme: {
    extend: {},
  },
  plugins: [],
};

export default config;
