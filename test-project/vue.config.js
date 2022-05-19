const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})

const GoogleFontsPlugin = require("@beyonk/google-fonts-webpack-plugin")

// vue.config.js
module.exports = {
    configureWebpack: {
        plugins: [
            new GoogleFontsPlugin({
                fonts: [
                    { family: "Poppins", variants: [ "500", "700" ] },
                    { family: "Roboto", variants: [ "500", "700" ] }
                ]
            })
        ]
    },
    devServer: {
        port: 80,
        host: '0.0.0.0'
    }
}
