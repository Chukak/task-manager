var path = require('path');

module.exports = {
	entry: path.resolve(__dirname, 'src', 'index.js'),
  output: {
		path: path.resolve(__dirname, 'dist'),
		publicPath: '/',
    filename: 'bundle.js'
	},
	module: {
		rules: [
			{
				test: /\.js/,
				exclude: /node_modules/,
        loader: 'babel-loader'
			},
			{
        test: /\.css$/i,
        loader: 'css-loader',
      }
		]
	}
}