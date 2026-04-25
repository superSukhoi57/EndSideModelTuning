import { createApiClient } from './api.ts';
import { config } from '../stores/config.ts';
import { Machine, MachineCreateReq, MachineUpdateReq, PageResponse, PageParams } from './types.ts';

const api = createApiClient(config.machine);

export const machineService = {
    create: (data: MachineCreateReq) => {
        return api.post('/api/v1/machine', data);
    },

    update: (data: MachineUpdateReq) => {
        return api.put('/api/v1/machine', data);
    },

    delete: (id: number) => {
        return api.delete(`/api/v1/machine/${id}`);
    },

    getById: (id: number) => {
        return api.get<Machine>(`/api/v1/machine/${id}`);
    },

    list: (params: PageParams & { ip?: string; userid?: number; isfinsh?: number }) => {
        return api.get<PageResponse<Machine>>('/api/v1/machines', { params });
    },
};
