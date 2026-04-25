import React, { useState } from 'react';
import { Button, Space, message, Spin } from 'antd';
import { ReloadOutlined } from '@ant-design/icons';
import { authStore } from '../stores/auth.ts';
import axios from 'axios';
import { config } from '../stores/config.ts';

interface UserInfoData {
    open_id: string;
    union_id: string;
    name: string;
    en_name: string;
    email: string;
    mobile: string;
    avatar_url: string;
    tenant_name: string;
}

const UserInfo: React.FC = () => {
    const auth = authStore.getAuth();
    const [userInfo, setUserInfo] = useState<UserInfoData | null>(auth.userInfo);
    const [refreshing, setRefreshing] = useState(false);

    const handleRefresh = async () => {
        setRefreshing(true);
        try {
            const authData = authStore.getAuth();
            if (!authData.accessToken) {
                message.error('未登录');
                return;
            }
            const response = await axios.get(`${config.auth}/auth/token/verify`, {
                headers: { Authorization: authData.accessToken },
                params: { access_token: authData.accessToken },
            });
            if (response.data.user_info) {
                setUserInfo(response.data.user_info);
                authStore.saveAuth(authData.accessToken, authData.refreshToken || '', response.data.user_info);
                message.success('刷新成功');
            }
        } catch {
            message.error('刷新失败');
        } finally {
            setRefreshing(false);
        }
    };

    const handleLogout = () => {
        authStore.clearAuth();
        window.location.href = '/login';
    };

    return (
        <div className="page-content">
            <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <h1 style={{ margin: 0 }}>用户信息</h1>
                <Button icon={<ReloadOutlined />} onClick={handleRefresh} loading={refreshing}>
                    刷新
                </Button>
            </div>
            <Spin spinning={refreshing}>
                {userInfo ? (
                    <div className="user-info-card">
                        <div className="user-avatar">
                            <img src={userInfo.avatar_url} alt="avatar" />
                        </div>
                        <div className="user-details">
                            <p><strong>姓名:</strong> {userInfo.name}</p>
                            <p><strong>英文名:</strong> {userInfo.en_name}</p>
                            <p><strong>邮箱:</strong> {userInfo.email}</p>
                            <p><strong>手机:</strong> {userInfo.mobile}</p>
                            <p><strong>租户:</strong> {userInfo.tenant_name}</p>
                        </div>
                        <button className="logout-button" onClick={handleLogout}>
                            退出登录
                        </button>
                    </div>
                ) : (
                    <p>暂无用户信息</p>
                )}
            </Spin>
        </div>
    );
};

export default UserInfo;
