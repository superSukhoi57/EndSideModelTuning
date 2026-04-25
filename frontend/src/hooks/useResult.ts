import { useState, useCallback } from 'react';
import { resultService } from '../services/resultService.ts';
import { Result, ResultCreateReq, ResultUpdateReq, PageResponse } from '../services/types.ts';

interface UseResultReturn {
    loading: boolean;
    data: Result | null;
    listData: PageResponse<Result> | null;
    error: string | null;
    createResult: (data: ResultCreateReq) => Promise<void>;
    updateResult: (data: ResultUpdateReq) => Promise<void>;
    deleteResult: (id: number) => Promise<void>;
    getResultById: (id: number) => Promise<void>;
    listResults: (params: { page: number; pageSize: number; userid?: number; machineid?: number; desc?: string }) => Promise<void>;
}

export const useResult = (): UseResultReturn => {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState<Result | null>(null);
    const [listData, setListData] = useState<PageResponse<Result> | null>(null);
    const [error, setError] = useState<string | null>(null);

    const createResult = useCallback(async (req: ResultCreateReq) => {
        setLoading(true);
        setError(null);
        try {
            await resultService.create(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '创建失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const updateResult = useCallback(async (req: ResultUpdateReq) => {
        setLoading(true);
        setError(null);
        try {
            await resultService.update(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '更新失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const deleteResult = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            await resultService.delete(id);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '删除失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const getResultById = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            const res = await resultService.getById(id);
            setData(res.data);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '获取详情失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const listResults = useCallback(async (params: { page: number; pageSize: number; userid?: number; machineid?: number; desc?: string }) => {
        setLoading(true);
        setError(null);
        try {
            const res = await resultService.list(params);
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
        createResult,
        updateResult,
        deleteResult,
        getResultById,
        listResults,
    };
};
