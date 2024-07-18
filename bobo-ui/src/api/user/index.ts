// 统一管理用户相关接口
import request from '@/utils/request'
enum API {
    SendCodeUrl = '/api/code',
    TestUrl = '/api/ping',
    RegisterUrl = '/api/user/register',
}
export const SendCode = (data:any) => {
    return request.get<any, any>(API.SendCodeUrl, data);
}
export const Test = () => {
    return request.get<any, any>(API.TestUrl);
}
export const Register = (data:any) => {
    return request.post<any, any>(API.RegisterUrl,data);
}