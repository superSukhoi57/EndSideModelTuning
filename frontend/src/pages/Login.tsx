import React, { useEffect, useState, useRef } from 'react';
import './Login.css';

const FEISHU_APP_ID = 'cli_a960a3aa0db89bd5';
const REDIRECT_URI = 'http://47.115.225.81:12300/callback';

declare global {
    interface Window {
        QRLogin?: (config: any) => any;
        LarkSSOSDKWebQRCode?: any;
    }
}

const Login: React.FC = () => {
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const qrLoginObjRef = useRef<any>(null);

    useEffect(() => {
        const script = document.createElement('script');
        script.src = 'https://lf-package-cn.feishucdn.com/obj/feishu-static/lark/passport/qrcode/LarkSSOSDKWebQRCode-1.0.3.js';
        script.async = true;
        script.onload = () => {
            setTimeout(() => {
                initFeishuQRCode();
            }, 300);
        };
        script.onerror = () => {
            setError('二维码加载失败');
            setLoading(false);
        };
        document.body.appendChild(script);

        return () => {
            const existingScript = document.querySelector('script[src*="LarkSSOSDKWebQRCode"]');
            if (existingScript && existingScript.parentNode) {
                existingScript.parentNode.removeChild(existingScript);
            }
        };
    }, []);

    const initFeishuQRCode = () => {
        const container = document.getElementById('feishu-login-container');
        if (!container) {
            console.error('Container element not found');
            setError('二维码容器未找到');
            setLoading(false);
            return;
        }

        const state = generateState();
        localStorage.setItem('feishu_auth_state', state);

        const redirectUri = encodeURIComponent(REDIRECT_URI);
        const goto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${FEISHU_APP_ID}&redirect_uri=${redirectUri}&response_type=code&state=${state}`;

        try {
            qrLoginObjRef.current = window.QRLogin!({
                id: 'feishu-login-container',
                goto: goto,
                width: '300',
                height: '400',
                style: 'width:300px;height:400px;',
            });

            setupMessageListener(goto);
            setLoading(false);
        } catch (e: any) {
            console.error('QRLogin error:', e);
            setError('二维码组件初始化失败: ' + e.message);
            setLoading(false);
        }
    };

    const setupMessageListener = (goto: string) => {
        const handleMessage = (event: MessageEvent) => {
            if (
                qrLoginObjRef.current &&
                qrLoginObjRef.current.matchOrigin(event.origin) &&
                qrLoginObjRef.current.matchData(event.data)
            ) {
                const loginTmpCode = event.data.tmp_code;
                window.location.href = `${goto}&tmp_code=${loginTmpCode}`;
            }
        };

        if (typeof window.addEventListener !== 'undefined') {
            window.addEventListener('message', handleMessage, false);
        }
    };

    const generateState = (): string => {
        return Math.random().toString(36).substring(2) + Date.now().toString(36);
    };

    return (
        <div className="login-container">
            <div className="login-card">
                <h1 className="login-title">飞书扫码登录</h1>
                <p className="login-subtitle">请使用飞书APP扫描二维码登录</p>
                {error ? (
                    <div className="error-container">
                        <p className="error-message">{error}</p>
                        <button className="retry-button" onClick={() => window.location.reload()}>
                            刷新重试
                        </button>
                    </div>
                ) : (
                    <>
                        {loading && (
                            <div className="loading-container">
                                <div className="loading-spinner"></div>
                                <p>正在加载二维码...</p>
                            </div>
                        )}
                        <div id="feishu-login-container" className="qr-code-container"></div>
                    </>
                )}
            </div>
        </div>
    );
};

export default Login;
