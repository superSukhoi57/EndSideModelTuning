import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { message } from 'antd';
import { authStore } from '../stores/auth.ts';
import { config } from '../stores/config.ts';

let isRefreshing = false;
let failedQueue: Array<{
    resolve: (value?: unknown) => void;
    reject: (reason?: unknown) => void;
}> = [];

const processQueue = (error: unknown | null, token: string | null = null) => {
    failedQueue.forEach((prom) => {
        if (error) {
            prom.reject(error);
        } else {
            prom.resolve(token);
        }
    });
    failedQueue = [];
};

const createApiClient = (baseURL: string): AxiosInstance => {
    const api = axios.create({
        baseURL,
        timeout: 30000,
        headers: {
            'Content-Type': 'application/json',
        },
    });

    api.interceptors.request.use(
        (config) => {
            const auth = authStore.getAuth();
            if (auth.accessToken) {
                config.headers.Authorization = auth.accessToken;
            }
            return config;
        },
        (error) => {
            return Promise.reject(error);
        }
    );

    api.interceptors.response.use(
        (response: AxiosResponse) => {
            return response;
        },
        async (error) => {
            const originalRequest = error.config as AxiosRequestConfig & { _retry?: boolean };

            if (error.response?.status === 401 && !originalRequest._retry) {
                if (isRefreshing) {
                    return new Promise((resolve, reject) => {
                        failedQueue.push({ resolve, reject });
                    })
                        .then((token) => {
                            if (originalRequest.headers) {
                                originalRequest.headers.Authorization = token as string;
                            }
                            return api(originalRequest);
                        })
                        .catch((err) => {
                            return Promise.reject(err);
                        });
                }

                originalRequest._retry = true;
                isRefreshing = true;

                try {
                    const auth = authStore.getAuth();
                    if (!auth.refreshToken) {
                        authStore.clearAuth();
                        window.location.href = '/login';
                        return Promise.reject(error);
                    }

                    const refreshResponse = await axios.post(`${config.auth}/auth/refresh`, {
                        refresh_token: auth.refreshToken,
                    });

                    const { access_token, refresh_token } = refreshResponse.data;
                    authStore.saveAuth(access_token, refresh_token, auth.userInfo || {} as any);

                    processQueue(null, access_token);
                    if (originalRequest.headers) {
                        originalRequest.headers.Authorization = access_token;
                    }
                    return api(originalRequest);
                } catch (refreshError) {
                    processQueue(refreshError, null);
                    authStore.clearAuth();
                    window.location.href = '/login';
                    return Promise.reject(refreshError);
                } finally {
                    isRefreshing = false;
                }
            }

            // 其他错误显示 message 提示（包括网络错误、超时等）
            const errorMsg = error.response?.data?.message ||
                error.response?.data?.msg ||
                error.message ||
                '请求失败';
            message.error(errorMsg, 3);

            return Promise.reject(error);
        }
    );

    return api;
};

export { createApiClient };
