import { request } from 'umi';

export interface LoginParamsType {
  userName: string;
  password: string;
  mobile: string;
  captcha: string;
  type: string;
}
/**
 *  用户登录
 * @param params 请求参数
 * @returns
 */
export async function fakeAccountLogin(params: LoginParamsType) {
  return request<API.LoginStateType>('/api/sys/user/login', {
    method: 'POST',
    data: params,
  });
}
/**
 * 获取手机验证码
 * @param mobile 手机号码
 * @returns
 */
export async function getFakeCaptcha(mobile: string) {
  return request(`/api/sys/login/captcha?mobile=${mobile}`);
}
/**
 * 退出登录
 * @returns
 */
export async function outLogin() {
  return request('/api/sys/login/outLogin');
}
