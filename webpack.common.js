'use strict';

const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const {VueLoaderPlugin} = require('vue-loader');

const FRONTEND_DIR = path.resolve(__dirname, 'frontend');
const STATIC_DIR = path.resolve(FRONTEND_DIR, 'public');
const TEMPLATES_DIR = path.resolve(FRONTEND_DIR, 'templates');

const config = {
  entry: path.resolve(FRONTEND_DIR, 'main.mjs'),
  output: {
    filename: '[name]-[contenthash].js',
    path: path.resolve(STATIC_DIR, '_assets'),
    hashDigestLength: 8
  },
  externals: {
    'lodash': '_',
    'jquery': 'jQuery',
    'vue': 'Vue',
    'bootstrap-vue': 'BootstrapVue',
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        use: ['vue-loader']
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      },
      {
        test: /\.m?js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              ['@babel/preset-env', {targets: 'defaults'}]
            ]
          }
        }
      }
    ]
  },
  plugins: [
    new VueLoaderPlugin(),
    new HtmlWebpackPlugin({
      template: path.resolve(FRONTEND_DIR, 'index.gohtml'),
      filename: path.resolve(TEMPLATES_DIR, 'webpack-index.gohtml'),
      scriptLoading: 'blocking',
      inject: 'body',
      publicPath: '/_assets/',
    })
  ]
};

module.exports = config;
