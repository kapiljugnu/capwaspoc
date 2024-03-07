/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/index.html", "./src/js/*.js",  "./go/templ/*.templ"],
  plugins: [require("daisyui"), require('tailwindcss-safe-area')],
  daisyui: {
    themes: false
    // themes: ["forest", "cyberpunk"],
    // lightTheme: "forest",
    // darkTheme: "cyberpunk",
  }
}

