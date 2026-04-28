import { useState, useCallback } from 'react';
import { machineService } from '../services/machineService.ts';
import { Machine, MachineCreateReq, MachineUpdateReq, PageResponse } from '../services/types.ts';

interface UseMachineReturn {
    loading: boolean;
    data: Machine | null;
    listData: PageResponse<Machine> | null;
    error: string | null;
    createMachine: (data: MachineCreateReq) => Promise<void>;
    updateMachine: (data: MachineUpdateReq) => Promise<void>;
    deleteMachine: (id: number) => Promise<void>;
    getMachineById: (id: number) => Promise<void>;
    listMachines: (params: { page: number; pageSize: number; ip?: string; isfinsh?: number }) => Promise<void>;
}

export const useMachine = (): UseMachineReturn => {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState<Machine | null>(null);
    const [listData, setListData] = useState<PageResponse<Machine> | null>(null);
    const [error, setError] = useState<string | null>(null);

    const createMachine = useCallback(async (req: MachineCreateReq) => {
        setLoading(true);
        setError(null);
        try {
            await machineService.create(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '创建失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const updateMachine = useCallback(async (req: MachineUpdateReq) => {
        setLoading(true);
        setError(null);
        try {
            await machineService.update(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '更新失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const deleteMachine = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            await machineService.delete(id);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '删除失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const getMachineById = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            const res = await machineService.getById(id);
            setData(res.data);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '获取详情失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const listMachines = useCallback(async (params: { page: number; pageSize: number; ip?: string; isfinsh?: number }) => {
        setLoading(true);
        setError(null);
        try {
            const res = await machineService.list(params);
            setListData(res.data);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '获取列表失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    return {
        loading,
        data,
        listData,
        error,
        createMachine,
        updateMachine,
        deleteMachine,
        getMachineById,
        listMachines,
    };
};
