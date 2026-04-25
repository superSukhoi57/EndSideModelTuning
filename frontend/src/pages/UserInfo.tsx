import React from 'react';
import { authStore } from '../stores/auth.ts';

const UserInfo: React.FC = () => {
    const auth = authStore.getAuth();

    const handleLogout = () => {
        authStore.clearAuth();
        window.location.href = '/login';
    };

    return (
        <div className="page-content">
            <h1>用户信息</h1>
            {auth.userInfo ? (
                <div className="user-info-card">
                    <div className="user-avatar">
                        <img src={auth.userInfo.avatar_url} alt="avatar" />
                    </div>
                    <div className="user-details">
                        <p><strong>姓名:</strong> {auth.userInfo.name}</p>
                        <p><strong>英文名:</strong> {auth.userInfo.en_name}</p>
                        <p><strong>邮箱:</strong> {auth.userInfo.email}</p>
                        <p><strong>手机:</strong> {auth.userInfo.mobile}</p>
                        <p><strong>租户:</strong> {auth.userInfo.tenant_name}</p>
                    </div>
                    <button className="logout-button" onClick={handleLogout}>
                        退出登录
                    </button>
                </div>
            ) : (
                <p>暂无用户信息</p>
            )}
        </div>
    );
};

export default UserInfo;
