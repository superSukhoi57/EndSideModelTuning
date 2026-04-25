import React, { useState } from 'react';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import { Layout, Menu, Typography } from 'antd';
import {
    AppstoreOutlined,
    ScheduleOutlined,
    UserOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    FileTextOutlined,
    FileOutlined,
} from '@ant-design/icons';
import './MainLayout.css';

const { Sider, Content } = Layout;
const { Title } = Typography;

interface NavItem {
    key: string;
    label: string;
    icon: React.ReactNode;
    path?: string;
    children?: { key: string; label: string; path: string }[];
}

const navItems: NavItem[] = [
    { key: 'devices', label: '设备管理', icon: <AppstoreOutlined />, path: '/devices' },
    { key: 'param_script', label: '参数/脚本管理', icon: <FileOutlined />, path: '/param_script' },
    {
        key: 'tasks',
        label: '任务管理',
        icon: <ScheduleOutlined />,
        children: [
            { key: 'create-task', label: '创建任务', path: '/tasks/create' },
            { key: 'task-list', label: '任务列表', path: '/tasks/list' },
        ],
    },
    { key: 'results', label: '结果集', icon: <FileTextOutlined />, path: '/results' },
    { key: 'profile', label: '用户信息', icon: <UserOutlined />, path: '/profile' },
];

const MainLayout: React.FC = () => {
    const [collapsed, setCollapsed] = useState(false);
    const navigate = useNavigate();
    const location = useLocation();

    const handleMenuClick = ({ key, keyPath }: { key: string; keyPath: string[] }) => {
        const findPath = (items: NavItem[], key: string): string | undefined => {
            for (const item of items) {
                if (item.key === key && item.path) return item.path;
                if (item.children) {
                    const childPath = findPath(item.children.map(c => ({ ...c, icon: undefined })), key);
                    if (childPath) return childPath;
                }
            }
            return undefined;
        };
        const path = findPath(navItems, key);
        if (path) {
            navigate(path);
        }
    };

    const getMenuItems = () => {
        return navItems.map((item) => {
            if (item.children) {
                return {
                    key: item.key,
                    icon: item.icon,
                    label: item.label,
                    children: item.children.map((child) => ({
                        key: child.key,
                        label: child.label,
                    })),
                };
            }
            return {
                key: item.key,
                icon: item.icon,
                label: item.label,
            };
        });
    };

    const getSelectedKey = () => {
        const pathname = location.pathname;
        if (pathname.startsWith('/tasks/create')) return 'create-task';
        if (pathname.startsWith('/tasks/list')) return 'task-list';
        return pathname.split('/').filter(Boolean)[0] || 'devices';
    };

    return (
        <Layout className="main-layout">
            <Sider
                trigger={null}
                collapsible
                collapsed={collapsed}
                className="sidebar"
                width={240}
            >
                <div className="sidebar-header">
                    {!collapsed && <Title level={4} className="brand-text">SZTU 2026</Title>}
                    <div
                        className="collapse-btn"
                        onClick={() => setCollapsed(!collapsed)}
                    >
                        {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                    </div>
                </div>
                <Menu
                    mode="inline"
                    selectedKeys={[getSelectedKey()]}
                    items={getMenuItems()}
                    onClick={handleMenuClick}
                />
            </Sider>
            <Layout>
                <Content className="main-content">
                    <Outlet />
                </Content>
            </Layout>
        </Layout>
    );
};

export default MainLayout;
