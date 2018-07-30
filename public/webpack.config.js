const path = require('path');

var config = {
	mode: "none",
	entry: path.join(__dirname, "src", "main.js"),
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
				test: /\.exec\.js$/,
				use: [ 'script-loader' ]
			}
		]
	}
}

module.exports = config;