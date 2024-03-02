/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/index.html", "./src/js/*.js"],
  plugins: [require("daisyui"), require('tailwindcss-safe-area')],
  safelist: [ // list of classes added to go templ 
    "drawer-side",
    "drawer-overlay",
    "menu",
    "p-4",
    "w-80",
    "min-h-full",
    "bg-base-200",
    "text-base-content"
  ],
  daisyui: {
    themes: false
    // themes: ["forest", "cyberpunk"],
    // lightTheme: "forest",
    // darkTheme: "cyberpunk",
  }
}
