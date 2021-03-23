'use strict';

const path = require('path');
const {merge} = require('webpack-merge');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const common = require('./webpack.common.js');


module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
  devServer: {
    contentBase: path.resolve(__dirname, 'frontend', 'public'),
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, 'frontend', 'index.gohtml'),
      filename: path.resolve(__dirname, 'frontend', 'public', 'index.html'),
      scriptLoading: 'blocking',
      inject: 'body',
      publicPath: '/_assets/',
    })
  ]
});
