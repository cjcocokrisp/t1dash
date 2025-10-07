/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/templates/web/templates/**/*.html",
  ],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: [
      {
        t1dash: {
          "primary": "#E63946",
          "secondary": "#ffb703",
          "accent": "#f77f00",
          "neutral": "#ffffff",
          "base-100": "#ffffff",
          "info": "#e63946",
          "success": "#4cAf50",
          "warning": "#00ff00",
          "error": "#ff0000",
        },
      },
    ],
  },
}

