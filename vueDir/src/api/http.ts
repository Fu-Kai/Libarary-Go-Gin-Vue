import Vue from 'vue'
import axios, {AxiosRequestConfig, AxiosResponse} from 'axios';
import VueAxios from 'vue-axios'


export function request(config: any) {

    // 1.创建axios实例
    const instance = axios.create({
        //后端接口
        baseURL: 'http://localhost:7070',
        timeout: 5000,
        headers: {}
    })
    // 请求拦截
    instance.interceptors.request.use((config: AxiosRequestConfig) => {
        // 加载
        return config;
    })
    // 响应拦截
    instance.interceptors.response.use((response: AxiosResponse<any>) => {
        // 结束loading
        console.log(response)
       // window.location.href = "/";
        return response;
    }, error => {
        // 结束loading
        // 错误提醒
        return Promise.reject(error);
    })

    // 发送真正的网络请求
    return instance(config)
}
