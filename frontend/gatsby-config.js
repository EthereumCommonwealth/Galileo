const rupture = require("rupture");
const nib = require("nib");

module.exports = {
  siteMetadata: {
    title: 'Galileo',
  },
  plugins: [
    'gatsby-plugin-react-helmet',
    'gatsby-transformer-json',
    'gatsby-plugin-google-analytics',
    {
      resolve: "gatsby-plugin-stylus",
      options: {
        use: [nib(), rupture()],
      },
    },
  ],
};
