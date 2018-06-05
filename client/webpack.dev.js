/* eslint-disable import/no-extraneous-dependencies */
const merge = require('webpack-merge');
const webpack = require('webpack');
/* eslint-enable */
const common = require('./webpack.common');

module.exports = merge(common, {
  mode: 'developement',
  devServer: {
    historyApiFallback: true,
    port: 8082,
    hot: true,
    overlay: true,
    proxy: {
      '/api': 'http://localhost:8080',
      '/twitter': 'http://localhost:8080',
    },
  },
  plugins: [
    new webpack.DefinePlugin({
      __API_BASE_URL__: JSON.stringify('/api'),
    }),
  ],
});
