module.exports = {
  publicPath: process.env.PUBLIC_PATH == null 
    ? '/'
    : '/demo/vue-websocket-echo',
  runtimeCompiler: true
}
