// prettier-ignore
const colors = require('tailwindcss/colors')

module.exports = {
	mode: 'jit',
	purge: ["web/templates/**/*.{html,tmpl}"],
	darkMode: false, // or 'media' or 'class'
	theme: {
		extend: {}
	},
	variants: {
		extend: {},
	},
	plugins: [],
};
