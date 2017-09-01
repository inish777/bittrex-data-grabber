const path = require('path')

module.exports = {
    entry: './index.es6',
    output: {
        path: path.resolve(__dirname, "dist"),
        filename: "bundle.js",
        publicPath: 'dist/',
        sourceMapFilename: "[file].map"
    },
    devtool: 'source-map',
    devServer: {
        proxy: { // proxy URLs to backend development server
            '/api': 'http://localhost:3000'
        },
        watchOptions: {
            aggregateTimeout: 300,
            poll: 1000
        },
        historyApiFallback: true, // true for index.html upon 404, object for multiple paths
        hot: true, // hot module replacement. Depends on HotModuleReplacementPlugin
        https: false, // true for self-signed, object for cert authority
        noInfo: true, // only errors & warns on hot reload,
    },
    module: {
        rules: [
            {
                test: /\.es6$/,
                loader: "babel-loader",
            },
            {
                test: /\.vue$/,
                loader: "vue-loader"
            }
        ]
    }
}