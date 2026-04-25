interface AppConfig {
    auth: string;
    machine: string;
    paramater: string;
    tasks: string;
    result: string;
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

await loadConfig();

export { config };
