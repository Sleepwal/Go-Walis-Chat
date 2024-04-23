import axios from "axios"
import {API_URL} from "@/utils/consts";

// 创建axios实例
const request = axios.create({
    // axios中请求配置有baseURL选项，表示请求URL公共部分
    baseURL: API_URL,
    // 超时
    timeout: 20000
})

request.interceptors.request.use(config => {
    // 在发送请求之前做某事
    return config
}, error => {
    // 请求错误时做些事
    return Promise.reject(error)
})


request.interceptors.response.use(res => {
    return res.data
}, error => {
    // 请求错误时做些事
    return Promise.reject(error)
})

export default request