import { useState, useCallback } from 'react';
import { taskService } from '../services/taskService.ts';
import { Task, TaskCreateReq, TaskUpdateReq, PageResponse } from '../services/types.ts';

interface UseTaskReturn {
    loading: boolean;
    data: Task | null;
    listData: PageResponse<Task> | null;
    error: string | null;
    createTask: (data: TaskCreateReq) => Promise<void>;
    updateTask: (data: TaskUpdateReq) => Promise<void>;
    deleteTask: (id: number) => Promise<void>;
    getTaskById: (id: number) => Promise<void>;
    listTasks: (params: { page: number; pageSize: number; paramterid?: number; desc?: string }) => Promise<void>;
}

export const useTask = (): UseTaskReturn => {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState<Task | null>(null);
    const [listData, setListData] = useState<PageResponse<Task> | null>(null);
    const [error, setError] = useState<string | null>(null);

    const createTask = useCallback(async (req: TaskCreateReq) => {
        setLoading(true);
        setError(null);
        try {
            await taskService.create(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '创建失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const updateTask = useCallback(async (req: TaskUpdateReq) => {
        setLoading(true);
        setError(null);
        try {
            await taskService.update(req);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '更新失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const deleteTask = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            await taskService.delete(id);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '删除失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const getTaskById = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            const res = await taskService.getById(id);
            setData(res.data);
        } catch (err: unknown) {
            const message = err instanceof Error ? err.message : '获取详情失败';
            setError(message);
        } finally {
            setLoading(false);
        }
    }, []);

    const listTasks = useCallback(async (params: { page: number; pageSize: number; paramterid?: number; desc?: string }) => {
        setLoading(true);
        setError(null);
        try {
            const res = await taskService.list(params);
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
        createTask,
        updateTask,
        deleteTask,
        getTaskById,
        listTasks,
    };
};
