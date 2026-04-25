import { createApiClient } from './api.ts';
import { config } from '../stores/config.ts';
import { Task, TaskCreateReq, TaskUpdateReq, PageResponse, PageParams } from './types.ts';

const api = createApiClient(config.tasks);

export const taskService = {
    create: (data: TaskCreateReq) => {
        return api.post('/api/v1/task', data);
    },

    update: (data: TaskUpdateReq) => {
        return api.put('/api/v1/task', data);
    },

    delete: (id: number) => {
        return api.delete(`/api/v1/task/${id}`);
    },

    getById: (id: number) => {
        return api.get<Task>(`/api/v1/task/${id}`);
    },

    list: (params: PageParams & { userid?: number; paramterid?: number; desc?: string }) => {
        return api.get<PageResponse<Task>>('/api/v1/tasks', { params });
    },
};
