/* eslint-disable import/no-extraneous-dependencies */
const merge = require('webpack-merge');
const webpack = require('webpack');
/* eslint-enable */
const common = require('./webpack.common');

module.exports = merge(common, {
  mode: 'production',
  plugins: [
    new webpack.DefinePlugin({
      __API_BASE_URL__: JSON.stringify('https://api.wipeinc.io'),
    }),
  ],
});
