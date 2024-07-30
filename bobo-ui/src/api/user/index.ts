// 统一管理用户相关接口
import request from '@/utils/request'
enum API {
    SendCodeUrl = '/api/code',
    TestUrl = '/api/ping',
    RegisterUrl = '/api/user/register',
    LoginUrl = '/api/user/login',
    ProfileUrl = '/api/user/profile',
    UpdatePasswordUrl = '/api/user/up_pass',
    UpdateEmailUrl = '/api/user/up_email',
    UpdateProfileUrl = '/api/user/profile',
    
}
//发送验证码
export const SendCode = (data:any) => {
    return request.get<any, any>(API.SendCodeUrl, data);
}
export const Test = () => {
    return request.get<any, any>(API.TestUrl);
}
//注册
export const Register = (data:any) => {
    return request.post<any, any>(API.RegisterUrl,data);
}
//登陆
export const Login = (data:any) => {
    return request.post<any, any>(API.LoginUrl,data);
}
//个人详情
export const Profile = () => {
    return request.get<any, any>(API.ProfileUrl);
}
//修改密码
export const UpdatePassword = (data:any) => {
    return request.put<any, any>(API.UpdatePasswordUrl,data);
}
//修改邮箱
export const UpdateEmail = (data:any) => {
    return request.put<any, any>(API.UpdateEmailUrl,data);
}
//修改个人信息
export const UpdateProfile = (data:any) => {
    return request.put<any, any>(API.UpdateProfileUrl,data);
}

