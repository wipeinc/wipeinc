const HtmlWebpackPlugin = require('html-webpack-plugin');
// const webpack = require('webpack');
const path = require('path');
const autoprefixer = require('autoprefixer');

const config = {
  entry: './src/index.jsx',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        },
      },
      // css
      {
        test: /\.css$/,
        include: /node_modules/,
        loader: [
          'style-loader',
          'css-loader',
        ],
      },
      // sass
      {
        test: /\.scss$/,
        use: [
          {
            loader: 'style-loader',
          },
          {
            loader: 'css-loader',
            options: {
              sourceMap: true,
            },
          },
          {
            loader: 'postcss-loader',
            options: {
              sourceMap: true,
              plugins() {
                return [autoprefixer('last 2 version')];
              },
            },
          },
          {
            loader: 'sass-loader',
            options: {
              sourceMap: true,
            },
          },
        ],
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({ template: './src/index.html' }),
  ],
};

module.exports = config;