/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/templates/web/templates/**/*.html",
  ],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
}

