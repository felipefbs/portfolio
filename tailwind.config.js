/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{templ,html}"],
  theme: {
    fontFamily: {
      body: ['Lato'],
    },
    extend: {
      colors: {
        background: {
          DEFAULT: "#09090b",
          muted: "rgba(24, 24, 27, 0.5)",
        },
        border: {
          DEFAULT: "#27272a",
          hover: "#3f3f46",
        },
        text: {
          primary: "#f4f4f5",
          secondary: "#d4d4d8",
          muted: "#a1a1aa",
          subtle: "#71717a",
        },
        accent: {
          DEFAULT: "#34d399",
          hover: "#6ee7b7",
          shadow: "rgba(16, 185, 129, 0.1)",
        },
      },
    },
  },
  plugins: [],
};
