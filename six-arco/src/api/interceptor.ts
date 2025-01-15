import axios from 'axios';
import type { AxiosRequestConfig, AxiosResponse } from 'axios';
import { Message, Modal } from '@arco-design/web-vue';
import { useUserStore } from '@/store';
import { getToken } from '@/utils/auth';
import router from "@/router";

export interface HttpResponse<T = unknown> {
    status: number;
    msg: string;
    code: number;
    total: number;
    data: T;
}

if (import.meta.env.VITE_API_BASE_URL) {
    axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
}

axios.interceptors.request.use(
    (config: AxiosRequestConfig) => {
        // let each request carry token
        // this example using the JWT token
        // Authorization is a custom headers key
        // please modify it according to the actual situation
        const token = getToken();
        if (token) {
            if (!config.headers) {
                config.headers = {};
            }
            config.headers.Authorization = `Bearer ${token}`;
            config.headers['Six-Token'] = token;
        }
        return config;
    },
    (error) => {
        // do something
        return Promise.reject(error);
    }
);

// 以下code全部都是http请求成功之后 服务端的业务码 和 http 状态码处于解耦状态，完全不相关
const ServiceCode = {
    Success: 0,
    Warning: 1,
    Error: 2,
    NoAuth: 3,
    Logout: 9,
};

// add response interceptors
axios.interceptors.response.use(
    (response: AxiosResponse<HttpResponse>) => {
        const res = response.data;
        if (res.code === ServiceCode.Success) {
            return res;
        }
        switch (res.code) {
            case ServiceCode.Warning:
                Message.warning({
                    content: res.msg || 'Warning',
                    duration: 5 * 1000,
                });
                return Promise.reject(res);
            case ServiceCode.Error:
                Message.error({
                    content: res.msg || 'Error',
                    duration: 5 * 1000,
                });
                return Promise.reject(new Error(res.msg || 'Error'));
            case ServiceCode.NoAuth:
                Message.error({
                    content: '无权限',
                    duration: 5 * 1000,
                });
                return Promise.reject(new Error(res.msg || 'Error'));
            case ServiceCode.Logout:
                Modal.error({
                    title: '登录状态失效',
                    content: '您的登录状态已失效，请重新登录',
                    okText: '重新登陆',
                    async onOk() {
                        const userStore = useUserStore();
                        await userStore.logout();
                        window.location.reload();
                    },
                });
                return Promise.reject(new Error(res.msg || '登录状态失效'));
            default:
                Message.warning({
                    content: res.msg || '未知错误',
                    duration: 5 * 1000,
                });
                return Promise.reject(new Error(res.msg || 'Error'));
        }
    },
    (error) => {
        Message.error({
            content: error.msg || 'Request Error',
            duration: 5 * 1000,
        });
        return Promise.reject(error);
    }
);
