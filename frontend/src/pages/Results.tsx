import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, InputNumber, Popconfirm, message, Space } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ReloadOutlined } from '@ant-design/icons';
import { useResult } from '../hooks/useResult.ts';
import type { ColumnsType } from 'antd/es/table';
import type { Result } from '../services/types.ts';

const Results: React.FC = () => {
    const { loading, listData, createResult, updateResult, deleteResult, listResults } = useResult();
    const [form] = Form.useForm();
    const [modalVisible, setModalVisible] = useState(false);
    const [editingResult, setEditingResult] = useState<Result | null>(null);
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);

    useEffect(() => {
        fetchData();
    }, [page, pageSize]);

    const fetchData = () => {
        listResults({ page, pageSize });
    };

    const handleCreate = () => {
        setEditingResult(null);
        form.resetFields();
        setModalVisible(true);
    };

    const handleEdit = (record: Result) => {
        setEditingResult(record);
        form.setFieldsValue(record);
        setModalVisible(true);
    };

    const handleDelete = async (id: number) => {
        try {
            await deleteResult(id);
            message.success('删除成功');
            fetchData();
        } catch {
            message.error('删除失败');
        }
    };

    const handleSubmit = async () => {
        try {
            const values = await form.validateFields();
            if (editingResult) {
                await updateResult({ id: editingResult.id, ...values });
                message.success('更新成功');
            } else {
                await createResult(values);
                message.success('创建成功');
            }
            setModalVisible(false);
            fetchData();
        } catch {
            message.error('操作失败');
        }
    };

    const columns: ColumnsType<Result> = [
        {
            title: '结果ID',
            dataIndex: 'id',
            key: 'id',
            width: 80,
        },
        {
            title: '用户ID',
            dataIndex: 'userid',
            key: 'userid',
            width: 100,
        },
        {
            title: '机器ID',
            dataIndex: 'machineid',
            key: 'machineid',
            width: 100,
        },
        {
            title: '结果数据',
            dataIndex: 'result',
            key: 'result',
            ellipsis: true,
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
            title: '操作',
            key: 'action',
            width: 150,
            render: (_, record: Result) => (
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
                <h1 style={{ margin: 0 }}>结果集</h1>
                <Space>
                    <Button icon={<ReloadOutlined />} onClick={fetchData} loading={loading}>
                        刷新
                    </Button>
                    <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>
                        新增结果
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
                title={editingResult ? '编辑结果' : '新增结果'}
                open={modalVisible}
                onOk={handleSubmit}
                onCancel={() => setModalVisible(false)}
                width={700}
            >
                <Form form={form} layout="vertical">
                    <Form.Item name="userid" label="用户ID" rules={[{ required: true, message: '请输入用户ID' }]}>
                        <InputNumber style={{ width: '100%' }} />
                    </Form.Item>
                    <Form.Item name="machineid" label="机器ID" rules={[{ required: true, message: '请输入机器ID' }]}>
                        <InputNumber style={{ width: '100%' }} />
                    </Form.Item>
                    <Form.Item name="result" label="结果数据(JSON)" rules={[{ required: true, message: '请输入结果数据' }]}>
                        <Input.TextArea rows={6} />
                    </Form.Item>
                    <Form.Item name="desc" label="描述">
                        <Input />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    );
};

export default Results;
