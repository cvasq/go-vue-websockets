module.exports = {
  publicPath: process.env.PUBLIC_PATH == null 
    ? '/'
    : '/hello',
  runtimeCompiler: true
}
