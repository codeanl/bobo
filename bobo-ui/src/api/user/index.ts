// 统一管理用户相关接口
import request from '@/utils/request'
enum API {
    SendCodeUrl = '/api/code',
    TestUrl = '/api/ping',
    RegisterUrl = '/api/user/register',
    LoginUrl = '/api/user/login',
    ProfileUrl = '/api/user/profile',
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
export const Login = (data:any) => {
    return request.post<any, any>(API.LoginUrl,data);
}
export const Profile = () => {
    return request.get<any, any>(API.ProfileUrl);
}