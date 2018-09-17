const path = require('path');

var config = {
	mode: "development",
	entry: path.join(__dirname, "src", "App.jsx"),
	output: {
		path: path.join(__dirname, "dist"),
		filename: 'index.js',
	},
	devServer: {
		contentBase: path.join(__dirname, 'src'),
		port: 8080
	},
	module: {
		rules: [
			{
				test: /\.jsx?$/,
				exclude: /node_modules/,
				loader: 'babel-loader',
				query: {
					presets: ['es2015', 'react']
				}
			},
			{
				test: /\.css$/,
				use: [ 'style-loader', 'css-loader' ]
			}
		]
	}
}

module.exports = config;