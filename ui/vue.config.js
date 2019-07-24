module.exports = {
  publicPath: process.env.VUE_APP_BASE_PATH === undefined || process.env.VUE_APP_BASE_PATH === null
    ? '/'
    : process.env.VUE_APP_BASE_PATH,
  runtimeCompiler: true
}
