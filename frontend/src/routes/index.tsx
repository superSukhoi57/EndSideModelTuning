import React from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from '../pages/Login.tsx';
import Callback from '../pages/Callback.tsx';
import MainLayout from '../layouts/MainLayout.tsx';
import DeviceManagement from '../pages/DeviceManagement.tsx';
import TaskManagement from '../pages/TaskManagement/TaskManagement.tsx';
import CreateTask from '../pages/TaskManagement/CreateTask.tsx';
import TaskList from '../pages/TaskManagement/TaskList.tsx';
import UserInfo from '../pages/UserInfo.tsx';
import ParamScriptManagement from '../pages/ParamScriptManagement.tsx';
import Results from '../pages/Results.tsx';
import { authStore } from '../stores/auth.ts';

const PrivateRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    return authStore.isAuthenticated() ? <>{children}</> : <Navigate to="/login" />;
};

const AppRoutes: React.FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/callback" element={<Callback />} />
                <Route
                    path="/"
                    element={
                        <PrivateRoute>
                            <MainLayout />
                        </PrivateRoute>
                    }
                >
                    <Route index element={<Navigate to="/devices" replace />} />
                    <Route path="devices" element={<DeviceManagement />} />
                    <Route path="param_script" element={<ParamScriptManagement />} />
                    <Route path="tasks" element={<TaskManagement />}>
                        <Route index element={<Navigate to="/tasks/create" replace />} />
                        <Route path="create" element={<CreateTask />} />
                        <Route path="list" element={<TaskList />} />
                    </Route>
                    <Route path="results" element={<Results />} />
                    <Route path="profile" element={<UserInfo />} />
                </Route>
                <Route path="*" element={<Navigate to="/" replace />} />
            </Routes>
        </BrowserRouter>
    );
};

export default AppRoutes;
