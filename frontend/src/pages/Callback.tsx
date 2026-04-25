import React, { useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { authStore, UserInfo } from '../stores/auth.ts';

const Callback: React.FC = () => {
    const [searchParams] = useSearchParams();
    const navigate = useNavigate();

    useEffect(() => {
        const handleCallback = () => {
            const accessToken = searchParams.get('access_token');
            const refreshToken = searchParams.get('refresh_token');
            const userInfoStr = searchParams.get('user_info');

            console.log('Callback received - accessToken:', accessToken ? 'present' : 'missing');
            console.log('Callback received - refreshToken:', refreshToken ? 'present' : 'missing');
            console.log('Callback received - userInfo:', userInfoStr);

            if (!accessToken || !refreshToken || !userInfoStr) {
                console.error('Missing auth data');
                navigate('/login');
                return;
            }

            try {
                const userInfo = JSON.parse(userInfoStr) as UserInfo;
                console.log('Saving auth data...');
                authStore.saveAuth(accessToken, refreshToken, userInfo);
                console.log('Auth data saved, navigating to /');
                navigate('/');
            } catch (error) {
                console.error('Failed to parse user info:', error);
                navigate('/login');
            }
        };

        handleCallback();
    }, [searchParams, navigate]);

    return (
        <div className="callback-container">
            <div className="loading-spinner"></div>
            <p>正在登录...</p>
        </div>
    );
};

export default Callback;
