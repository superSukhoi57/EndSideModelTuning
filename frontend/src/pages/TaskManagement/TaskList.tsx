import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, InputNumber, Popconfirm, message, Space } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ReloadOutlined } from '@ant-design/icons';
import { useTask } from '../../hooks/useTask.ts';
import type { ColumnsType } from 'antd/es/table';
import type { Task } from '../../services/types.ts';

const TaskList: React.FC = () => {
    const { loading, listData, createTask, updateTask, deleteTask, listTasks } = useTask();
    const [form] = Form.useForm();
    const [modalVisible, setModalVisible] = useState(false);
    const [editingTask, setEditingTask] = useState<Task | null>(null);
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);

    useEffect(() => {
        fetchData();
    }, [page, pageSize]);

    const fetchData = () => {
        listTasks({ page, pageSize });
    };

    const handleCreate = () => {
        setEditingTask(null);
        form.resetFields();
        setModalVisible(true);
    };

    const handleEdit = (record: Task) => {
        setEditingTask(record);
        form.setFieldsValue(record);
        setModalVisible(true);
    };

    const handleDelete = async (id: number) => {
        try {
            await deleteTask(id);
            message.success('删除成功');
            fetchData();
        } catch {
            message.error('删除失败');
        }
    };

    const handleSubmit = async () => {
        try {
            const values = await form.validateFields();
            if (editingTask) {
                await updateTask({ id: editingTask.id, ...values });
                message.success('更新成功');
            } else {
                await createTask(values);
                message.success('创建成功');
            }
            setModalVisible(false);
            fetchData();
        } catch {
            message.error('操作失败');
        }
    };

    const columns: ColumnsType<Task> = [
        {
            title: '任务ID',
            dataIndex: 'id',
            key: 'id',
            width: 100,
        },
        {
            title: '参数ID',
            dataIndex: 'paramterid',
            key: 'paramterid',
            width: 100,
        },
        {
            title: '描述',
            dataIndex: 'desc',
            key: 'desc',
        },
        {
            title: '创建时间',
            dataIndex: 'createAt',
            key: 'createAt',
            width: 180,
        },
        {
            title: '更新时间',
            dataIndex: 'updateAt',
            key: 'updateAt',
            width: 180,
        },
        {
            title: '操作',
            key: 'action',
            width: 150,
            render: (_, record: Task) => (
                <Space>
                    <Button
                        type="link"
                        icon={<EditOutlined />}
                        onClick={() => handleEdit(record)}
                    >
                        编辑
                    </Button>
                    <Popconfirm
                        title="确定删除?"
                        onConfirm={() => handleDelete(record.id)}
                    >
                        <Button type="link" danger icon={<DeleteOutlined />}>
                            删除
                        </Button>
                    </Popconfirm>
                </Space>
            ),
        },
    ];

    return (
        <div className="page-content">
            <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <h1 style={{ margin: 0 }}>任务列表</h1>
                <Space>
                    <Button icon={<ReloadOutlined />} onClick={fetchData} loading={loading}>
                        刷新
                    </Button>
                    <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>
                        新增任务
                    </Button>
                </Space>
            </div>
            <Table
                columns={columns}
                dataSource={listData?.list || []}
                rowKey="id"
                loading={loading}
                pagination={{
                    current: page,
                    pageSize,
                    total: listData?.total || 0,
                    onChange: (p, ps) => {
                        setPage(p);
                        setPageSize(ps);
                    },
                }}
            />
            <Modal
                title={editingTask ? '编辑任务' : '新增任务'}
                open={modalVisible}
                onOk={handleSubmit}
                onCancel={() => setModalVisible(false)}
            >
                <Form form={form} layout="vertical">
                    <Form.Item name="id" label="任务ID" rules={[{ required: true, message: '请输入任务ID' }]}>
                        <InputNumber style={{ width: '100%' }} disabled={!!editingTask} />
                    </Form.Item>
                    <Form.Item name="paramterid" label="参数ID" rules={[{ required: true, message: '请输入参数ID' }]}>
                        <InputNumber style={{ width: '100%' }} disabled={!!editingTask} />
                    </Form.Item>
                    <Form.Item name="desc" label="描述">
                        <Input.TextArea />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    );
};

export default TaskList;
