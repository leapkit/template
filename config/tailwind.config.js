module.exports = {
  content: [
    'web/*.html',
    'web/**/*.html',
    "web/public/*.js",  
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}