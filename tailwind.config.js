/** @type {import('tailwindcss').Config} */
import daisyui from "daisyui"

export default {
  content: ["./view/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
  daisyui: {
    themes: ["light", "lofi", "dim"],
  },
  corePlugins: { preFlight: true },
}

// /** @type {import('tailwindcss').Config} */
// module.exports = {
//   content: ["./view/**/*.{html,js,templ,go}"],
//   theme: {
//     extend: {},
//   },
//   plugins: [require("daisyui")],
//   corePlugins = { preFlight: true };
// }

// export const content = ["./templates/**/*.templ", "./views/**/*.templ"]
// export const theme = {
//   extend: {
//     fontFamily: {
//       mono: ["Courier Prime", "monospace"],
//     },
//   },
// }
// export const plugins = []
// export const corePlugins = { preFlight: true }
