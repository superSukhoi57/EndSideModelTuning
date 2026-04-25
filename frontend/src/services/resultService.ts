import { createApiClient } from './api.ts';
import { config } from '../stores/config.ts';
import { Result, ResultCreateReq, ResultUpdateReq, PageResponse, PageParams } from './types.ts';

const api = createApiClient(config.result);

export const resultService = {
    create: (data: ResultCreateReq) => {
        return api.post('/api/v1/result', data);
    },

    update: (data: ResultUpdateReq) => {
        return api.put('/api/v1/result', data);
    },

    delete: (id: number) => {
        return api.delete(`/api/v1/result/${id}`);
    },

    getById: (id: number) => {
        return api.get<Result>(`/api/v1/result/${id}`);
    },

    list: (params: PageParams & { userid?: number; machineid?: number; desc?: string }) => {
        return api.get<PageResponse<Result>>('/api/v1/results', { params });
    },
};
