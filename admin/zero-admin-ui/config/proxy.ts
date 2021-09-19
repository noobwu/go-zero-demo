/**
 * 在生产环境 代理是无法生效的，所以这里没有生产环境的配置
 * The agent cannot take effect in the production environment
 * so there is no configuration of the production environment
 * For details, please see
 * https://pro.ant.design/docs/deploy
 */
/**
 const serveUrlMap = {
  dev: 'https://dev.pro.ant.design/',
  pre: 'https://pre.pro.ant.design/',
  test: 'https://test.pro.ant.design/',
  idc: 'https://idc.pro.ant.design/',
};
*/
const serveUrlMap = {
  dev: 'http://127.0.0.1:8888',
  pre: 'https://pre.pro.ant.design/',
  test: 'https://test.pro.ant.design/',
  idc: 'https://idc.pro.ant.design/',
};

const { SERVE_ENV = 'dev' } = process.env;

export default {
  proxy: {
    '/api/': {
      target: serveUrlMap[SERVE_ENV],
      changeOrigin: true,
      pathRewrite: { '^': '' },
    },
  },
};
