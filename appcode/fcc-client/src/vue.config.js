module.exports = {
  productionSourceMap: false,
  configureWebpack: {
    devServer: {
      proxy: {
        '/api': {
          target: 'http://localhost:1206', //
          changeOrigin: true, //
          pathRewrite: {
            '^/api': '/api'
          }
        }
      }
    }
  },
}
