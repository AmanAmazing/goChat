module.exports = {
    content: [
        './views/**/*.{html,templ}', // Adjust the path to match your template files
        './assets/**/*.{js,css}'
    ],
    theme: {
        extend: {},
    },
    plugins: [require("@tailwindcss/typography"), require('daisyui'),],
    daisyui: {
        themes: ["pastel",],
        base: true,
        styled: true,
        utils: true,
    },
}

