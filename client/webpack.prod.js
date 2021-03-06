/* eslint-disable import/no-extraneous-dependencies */
const merge = require('webpack-merge');
const webpack = require('webpack');
const CleanWebpackPlugin = require('clean-webpack-plugin');
// const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const { BaseHrefWebpackPlugin } = require('base-href-webpack-plugin');
const WebpackGoogleCloudStoragePlugin = require('webpack-google-cloud-storage-plugin');
/* eslint-enable */
const path = require('path');
const common = require('./webpack.common');
const gcloudIdentity = require('./gcloud-identity.json');

const dist = path.resolve(__dirname, 'dist');

module.exports = merge(common, {
  mode: 'production',
  // warning about size...
  performance: {
    hints: false,
  },
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
      __API_BASE_URL__: JSON.stringify('https://wipeinc.io/api'),
    }),
    // new BundleAnalyzerPlugin(),
    new BaseHrefWebpackPlugin({ baseHref: 'https://static.wipeinc.io/' }),
    new CleanWebpackPlugin([dist]),
    new WebpackGoogleCloudStoragePlugin({
      directory: './dist/',
      exclude: ['index.html'],
      include: [/\.js$/],
      storageOptions: {
        projectId: 'atomic-legacy-189222',
        credentials: gcloudIdentity,
      },
      uploadOptions: {
        bucketName: 'wipeinc',
        makePublic: true,
        gzip: false,
        destinationNameFn: file => file.name,
      },
    }),
  ],
});
