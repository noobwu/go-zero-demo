// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';
import routes from './routes';

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  layout: {
    // https://umijs.org/zh-CN/plugins/plugin-layout
    locale: false,
    siderWidth: 208,
    ...defaultSettings,
  },

  // https://umijs.org/zh-CN/plugins/plugin-locale
  locale: {
    // default zh-CN
    default: 'zh-CN',
    antd: true,
    // default true, when it is true, will use `navigator.language` overwrite default
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@ant-design/pro-layout/es/PageLoading',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes,
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    'primary-color': defaultSettings.primaryColor,
  },
  // esbuild is father build tools
  // https://umijs.org/plugins/plugin-esbuild
  esbuild: {},
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy['proxy'],
  manifest: {
    basePath: '/',
  },
   // Fast Refresh 热更新
   //https://umijs.org/zh-CN/docs/fast-refresh
  fastRefresh: {},
  //https://umijs.org/guide/boost-compile-speed
  nodeModulesTransform: {
    type: 'none',
  },
  //https://umijs.org/zh-CN/docs/mfsu
  mfsu: {},
  //https://webpack.docschina.org/blog/2020-10-10-webpack-5-release/
  webpack5: {},
  exportStatic: {},

});
