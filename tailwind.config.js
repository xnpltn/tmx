/** @type {import('tailwindcss').Config} */
export const content = ["./**/*.{templ,go,html,js}"];
export const theme = {
  extend: {
    colors: {
      extra_dark: "#292E39",
      medium_dark: "#2E3440",
      dark: "#434C5E",
      light_dark: "#4C566A",
      extra_light: "#FFFFFF",
      medium_light: "#f8f9fb",
      dark_light: "#ECEFF4",
    },

    fontFamily: {
      rubik: ['Rubik', 'sans-serif'],
    },

  },
  backgroundImage: {
    'hero': "url('../img/banner11.jpeg')",
  },
};
export const plugins = [
  require('@tailwindcss/forms')
];

export const darkMode = "selector"

