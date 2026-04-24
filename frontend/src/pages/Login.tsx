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
    const isInitRef = useRef(false); // 新增：标记是否已初始化

    useEffect(() => {
        // 防止严格模式下重复执行
        if (isInitRef.current) return;

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
        isInitRef.current = true; // 标记已初始化

        return () => {
            // 清理：移除脚本、销毁二维码实例、移除监听
            const existingScript = document.querySelector('script[src*="LarkSSOSDKWebQRCode"]');
            if (existingScript?.parentNode) existingScript.parentNode.removeChild(existingScript);
            if (qrLoginObjRef.current?.destroy) qrLoginObjRef.current.destroy();
            window.removeEventListener('message', handleMessage);
            isInitRef.current = false; // 重置标记
        };
    }, []);

    // 抽离成独立函数，方便移除监听
    const handleMessage = (event: MessageEvent) => {
        if (
            qrLoginObjRef.current &&
            qrLoginObjRef.current.matchOrigin(event.origin) &&
            qrLoginObjRef.current.matchData(event.data)
        ) {
            const loginTmpCode = event.data.tmp_code;
            const state = generateState();
            localStorage.setItem('feishu_auth_state', state);
            const redirectUri = encodeURIComponent(REDIRECT_URI);
            const goto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${FEISHU_APP_ID}&redirect_uri=${redirectUri}&response_type=code&state=${state}`;
            window.location.href = `${goto}&tmp_code=${loginTmpCode}`;
        }
    };

    const initFeishuQRCode = () => {
        const container = document.getElementById('feishu-login-container');
        if (!container) {
            console.error('Container element not found');
            setError('二维码容器未找到');
            setLoading(false);
            return;
        }
        // 关键：清空容器，防止叠加
        container.innerHTML = '';

        const state = generateState();
        localStorage.setItem('feishu_auth_state', state);
        const redirectUri = encodeURIComponent(REDIRECT_URI);
        const goto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${FEISHU_APP_ID}&redirect_uri=${redirectUri}&response_type=code&state=${state}`;

        try {
            // 先销毁旧实例
            if (qrLoginObjRef.current?.destroy) qrLoginObjRef.current.destroy();
            qrLoginObjRef.current = window.QRLogin!({
                id: 'feishu-login-container',
                goto: goto,
                width: '300',
                height: '400',
                style: 'width:300px;height:400px;',
            });

            // 先移除旧监听，再添加新监听
            window.removeEventListener('message', handleMessage);
            window.addEventListener('message', handleMessage, false);

            setLoading(false);
        } catch (e: any) {
            console.error('QRLogin error:', e);
            setError('二维码组件初始化失败: ' + e.message);
            setLoading(false);
        }
    };

    const generateState = (): string => {
        return Math.random().toString(36).substring(2) + Date.now().toString(36);
    };

    return (
        <div className="login-container">
            <div className="login-card">
                <h1 className="login-title">登录模型调优系统</h1>
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