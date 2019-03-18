require('dotenv').config()

const path = require('path')
const Dotenv = require('dotenv-webpack')

module.exports = {
  webpack: config => {
    /* eslint-disable */
    config.plugins = config.plugins || []

    config.plugins = [
      ...config.plugins,
    /* eslint-enable */

      // Read the .env file
      new Dotenv({
        path: path.join(__dirname, '.env'),
        systemvars: true,
      }),
    ]

    return config
  },
}
