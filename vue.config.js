module.exports = {
  pluginOptions: {
    electronBuilder: {
      nodeIntegration: true,
      productName: "ngs_wb_timer",
      appId: "com.sample.ngs_wb_timer",
      win: {
        icon: 'src/assets/app.ico',
        target: [
          {
            target: 'portable',
            arch: ['x64']
          }
        ]
      }
    }
  }
}