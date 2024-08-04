// 统一管理用户相关接口
import request from '@/utils/request'
enum API {
    DailyListUrl = '/api/daily/list',
    SaveOrUpdateUrl = '/api/daily/save_or_update',
    DailyInfoUrl = '/api/daily/info',
}
//列表
export const DailyList = (data:any) => {
    return request.get<any, any>(API.DailyListUrl,data);
}
//添加/创建
export const SaveOrUpdate = (data:any) => {
    return request.post<any, any>(API.SaveOrUpdateUrl,data);
}
//详情
export const DailyInfo = (data:any) => {
    return request.get<any, any>(API.DailyInfoUrl, data);
}