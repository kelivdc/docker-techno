/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/*.html",
    "./views/**/*.html",    
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["dark","cupcake"],
  },
  plugins: [require('daisyui')],
}

