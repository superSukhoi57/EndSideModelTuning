interface AppConfig {
    auth: string;
    interaction: string;
    loginTimeoutSec: number;
}

let config: AppConfig = {} as AppConfig;

async function loadConfig() {
    try {
        const response = await fetch('/config/app.json');
        const data = await response.json();
        config = data;
        console.log('配置加载成功:', config);
    } catch (error) {
        console.error('Failed to load config:', error);
    }
}

//等待配置加载完成
await loadConfig();

export { config };
