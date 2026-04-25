export interface UserInfo {
    open_id: string;
    union_id: string;
    name: string;
    en_name: string;
    email: string;
    mobile: string;
    avatar_url: string;
    tenant_name: string;
}

interface AuthState {
    accessToken: string | null;
    refreshToken: string | null;
    userInfo: UserInfo | null;
}

const ACCESS_TOKEN_KEY = 'access_token';
const REFRESH_TOKEN_KEY = 'refresh_token';
const USER_INFO_KEY = 'user_info';

export const authStore = {
    saveAuth: (accessToken: string, refreshToken: string, userInfo: UserInfo) => {
        localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
        localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
        localStorage.setItem(USER_INFO_KEY, JSON.stringify(userInfo));
    },

    getAuth: (): AuthState => {
        const accessToken = localStorage.getItem(ACCESS_TOKEN_KEY);
        const refreshToken = localStorage.getItem(REFRESH_TOKEN_KEY);
        const userInfoStr = localStorage.getItem(USER_INFO_KEY);
        const userInfo = userInfoStr ? JSON.parse(userInfoStr) : null;

        return {
            accessToken,
            refreshToken,
            userInfo,
        };
    },

    clearAuth: () => {
        localStorage.removeItem(ACCESS_TOKEN_KEY);
        localStorage.removeItem(REFRESH_TOKEN_KEY);
        localStorage.removeItem(USER_INFO_KEY);
    },

    isAuthenticated: (): boolean => {
        return !!localStorage.getItem(ACCESS_TOKEN_KEY);
    },
};
