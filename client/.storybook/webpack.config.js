// const autoprefixer = require('autoprefixer');
//
module.exports = {
  module: {
    rules: [
      // css
      {
        test: /\.css$/,
        include: /node_modules/,
        loader: [
          'style-loader',
          'css-loader',
        ],
      },
    ],
  }
}
