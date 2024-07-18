import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import { ConfigEnv, UserConfigExport, loadEnv } from 'vite'
import path from 'path'

export default ({ command, mode }: ConfigEnv): UserConfigExport => {
  // 获取各种环境下对应的变量
  let env = loadEnv(mode, process.cwd())
  return {
    plugins: [
      vue(),
      AutoImport({
        imports: [
          'vue',
          {
            'naive-ui': [
              'useDialog',
              'useMessage',
              'useNotification',
              'useLoadingBar'
            ]
          }
        ]
      }),
      Components({
        resolvers: [NaiveUiResolver()]
      })
    ],
    resolve: { alias: { '@': path.resolve('./src') } },
    // server: {
    //   proxy: {
    //     [env.VITE_APP_BASE_API]: {
    //       target: env.VITE_SERVE,
    //       // 需要代理跨域
    //       changeOrigin: true,
    //       rewrite: (path) => path.replace(/^\/api/, ''),
    //     },
    //   },
    // },
  }
}

  