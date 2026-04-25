import { createApiClient } from './api.ts';
import { config } from '../stores/config.ts';
import { Parameter, ParameterCreateReq, ParameterUpdateReq, PageResponse, PageParams } from './types.ts';

const api = createApiClient(config.paramater);

export const parameterService = {
    create: (data: ParameterCreateReq) => {
        return api.post('/api/v1/parameter', data);
    },

    update: (data: ParameterUpdateReq) => {
        return api.put('/api/v1/parameter', data);
    },

    delete: (id: number) => {
        return api.delete(`/api/v1/parameter/${id}`);
    },

    getById: (id: number) => {
        return api.get<Parameter>(`/api/v1/parameter/${id}`);
    },

    list: (params: PageParams & { userid?: number; desc?: string }) => {
        return api.get<PageResponse<Parameter>>('/api/v1/parameters', { params });
    },
};
