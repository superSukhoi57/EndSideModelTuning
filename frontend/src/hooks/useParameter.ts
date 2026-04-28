import { useState, useCallback } from 'react';
import { parameterService } from '../services/parameterService.ts';
import { Parameter, ParameterCreateReq, ParameterUpdateReq, PageResponse } from '../services/types.ts';

interface UseParameterReturn {
    loading: boolean;
    data: Parameter | null;
    listData: PageResponse<Parameter> | null;
    error: string | null;
    createParameter: (data: ParameterCreateReq) => Promise<void>;
    updateParameter: (data: ParameterUpdateReq) => Promise<void>;
    deleteParameter: (id: number) => Promise<void>;
    getParameterById: (id: number) => Promise<void>;
    listParameters: (params: { page: number; pageSize: number; desc?: string }) => Promise<void>;
}

export const useParameter = (): UseParameterReturn => {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState<Parameter | null>(null);
    const [listData, setListData] = useState<PageResponse<Parameter> | null>(null);
    const [error, setError] = useState<string | null>(null);

    const createParameter = useCallback(async (req: ParameterCreateReq) => {
        setLoading(true);
        setError(null);
        try {
            await parameterService.create(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '创建失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const updateParameter = useCallback(async (req: ParameterUpdateReq) => {
        setLoading(true);
        setError(null);
        try {
            await parameterService.update(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '更新失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const deleteParameter = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            await parameterService.delete(id);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '删除失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const getParameterById = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            const res = await parameterService.getById(id);
            setData(res.data);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '获取详情失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const listParameters = useCallback(async (params: { page: number; pageSize: number; desc?: string }) => {
        setLoading(true);
        setError(null);
        try {
            const res = await parameterService.list(params);
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
        createParameter,
        updateParameter,
        deleteParameter,
        getParameterById,
        listParameters,
    };
};
