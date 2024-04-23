/** @type {import('tailwindcss').Config} */

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      'dark-primary': '#131a1c',
      'dark-secondary': '#1b2224',
      red: '#e74c4c',
      green: '#6bb05d',
      blue: '#0183ff',
      grey: '#dddfe2',
      white: '#fff',
    },
  },
  plugins: [],
}

