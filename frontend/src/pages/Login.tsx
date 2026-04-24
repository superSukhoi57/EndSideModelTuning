import React, { useEffect, useState, useRef } from 'react';
import './Login.css';

const FEISHU_APP_ID = 'cli_a960a3aa0db89bd5';
const REDIRECT_URI = 'http://47.115.225.81:12300/callback';
const API_BASE_URL = 'http://47.115.225.81:12300';

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

    // 修复：使用锁，确保整个生命周期只初始化一次
    const initializedRef = useRef(false);

    useEffect(() => {
        // 严格模式下也强制只执行一次
        if (initializedRef.current) return;
        initializedRef.current = true;

        const init = async () => {
            // 防止重复插入脚本
            if (document.querySelector('script[src*="LarkSSOSDKWebQRCode"]')) {
                await new Promise(resolve => setTimeout(resolve, 300));
                initFeishuQRCode();
                return;
            }

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
        };

        init();

        // 清理逻辑
        return () => {
            if (qrLoginObjRef.current?.destroy) {
                qrLoginObjRef.current.destroy();
                qrLoginObjRef.current = null;
            }
            window.removeEventListener('message', handleMessage);
        };
    }, []);

    const fetchState = async (): Promise<string> => {
        try {
            const response = await fetch(`${API_BASE_URL}/state`);
            if (!response.ok) throw new Error('Failed to fetch state');
            const data = await response.json();
            return data.state;
        } catch (error) {
            console.error('Error fetching state:', error);
            throw error;
        }
    };

    // 缓存监听函数，防止重复绑定
    const handleMessage = useRef(async (event: MessageEvent) => {
        if (
            qrLoginObjRef.current &&
            qrLoginObjRef.current.matchOrigin(event.origin) &&
            qrLoginObjRef.current.matchData(event.data)
        ) {
            const loginTmpCode = event.data.tmp_code;
            const state = await fetchState();
            localStorage.setItem('feishu_auth_state', state);
            const redirectUri = encodeURIComponent(REDIRECT_URI);
            const goto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${FEISHU_APP_ID}&redirect_uri=${redirectUri}&response_type=code&state=${state}`;
            window.location.href = `${goto}&tmp_code=${loginTmpCode}`;
        }
    }).current;

    const initFeishuQRCode = async () => {
        const container = document.getElementById('feishu-login-container');
        if (!container) {
            setError('二维码容器未找到');
            setLoading(false);
            return;
        }

        // 强制清空容器 + 销毁旧实例
        container.innerHTML = '';
        if (qrLoginObjRef.current?.destroy) {
            qrLoginObjRef.current.destroy();
            qrLoginObjRef.current = null;
        }

        try {
            const state = await fetchState();
            localStorage.setItem('feishu_auth_state', state);
            const redirectUri = encodeURIComponent(REDIRECT_URI);
            const goto = `https://passport.feishu.cn/suite/passport/oauth/authorize?client_id=${FEISHU_APP_ID}&redirect_uri=${redirectUri}&response_type=code&state=${state}`;

            qrLoginObjRef.current = window.QRLogin!({
                id: 'feishu-login-container',
                goto: goto,
                width: '300',
                height: '400',
                style: 'width:300px;height:400px;',
            });

            // 只绑定一次监听
            window.removeEventListener('message', handleMessage);
            window.addEventListener('message', handleMessage, false);

            setLoading(false);
        } catch (e: any) {
            console.error('QRLogin error:', e);
            setError('二维码组件初始化失败: ' + e.message);
            setLoading(false);
        }
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