/* eslint-disable import/no-extraneous-dependencies */
const merge = require('webpack-merge');
const webpack = require('webpack');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const { BaseHrefWebpackPlugin } = require('base-href-webpack-plugin');
const  S3Plugin = require('webpack-s3-plugin')
/* eslint-enable */
const path = require('path');
const common = require('./webpack.common');

const dist = path.resolve(__dirname, 'dist');

module.exports = merge(common, {
  mode: 'production',
  output: {
    path: dist,
    chunkFilename: '[name].[chunkhash].bundle.js',
    filename: '[name].[chunkhash].bundle.js',
    publicPath: '',
  },
  optimization: {
    runtimeChunk: 'single',
    splitChunks: {
      chunks: 'async',
      minSize: 30000,
      minChunks: 1,
      maxAsyncRequests: 5,
      maxInitialRequests: 3,
      name: true,
      cacheGroups: {
        default: {
          minChunks: 2,
          priority: -20,
          reuseExistingChunk: true,
        },
        vendors: {
          test: /[\\/]node_modules[\\/]/,
          name: 'vendors',
          enforce: true,
          chunks: 'all',
        },
      },
    },
  },
  plugins: [
    new webpack.DefinePlugin({
      __API_BASE_URL__: JSON.stringify('https://wipeinc.io/api/'),
    }),
    // new BundleAnalyzerPlugin(),
    new BaseHrefWebpackPlugin({ baseHref: 'https://s3-eu-west-1.amazonaws.com/wipeinc/' }),
    new CleanWebpackPlugin([dist]),
    new S3Plugin({
      // Exclude uploading of html
      exclude: /.*\.html$/,
      // s3Options are required
      s3Options: {
        region: 'eu-west-1',
      },
      s3UploadOptions: {
        Bucket: 'wipeinc',
      },
    }),
  ],
});
