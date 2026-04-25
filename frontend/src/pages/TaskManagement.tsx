import React from 'react';
import { Outlet } from 'react-router-dom';

const TaskManagement: React.FC = () => {
    return (
        <div className="page-content">
            <Outlet />
        </div>
    );
};

export default TaskManagement;
