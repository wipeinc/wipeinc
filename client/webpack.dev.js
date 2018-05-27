/* eslint-disable import/no-extraneous-dependencies */
const merge = require('webpack-merge');
const webpack = require('webpack');
/* eslint-enable */
const common = require('./webpack.common');

module.exports = merge(common, {
  mode: 'developement',
  devServer: {
    historyApiFallback: true,
  },
  plugins: [
    new webpack.DefinePlugin({
      __API_BASE_URL__: JSON.stringify('http://localhost:8000'),
    }),
  ],
});
