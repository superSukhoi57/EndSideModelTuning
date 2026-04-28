import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, InputNumber, Popconfirm, message, Space } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ReloadOutlined } from '@ant-design/icons';
import { useParameter } from '../hooks/useParameter.ts';
import type { ColumnsType } from 'antd/es/table';
import type { Parameter } from '../services/types.ts';

const ParamScriptManagement: React.FC = () => {
    const { loading, listData, createParameter, updateParameter, deleteParameter, listParameters } = useParameter();
    const [form] = Form.useForm();
    const [modalVisible, setModalVisible] = useState(false);
    const [editingParam, setEditingParam] = useState<Parameter | null>(null);
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);

    useEffect(() => {
        fetchData();
    }, [page, pageSize]);

    const fetchData = () => {
        listParameters({ page, pageSize });
    };

    const handleCreate = () => {
        setEditingParam(null);
        form.resetFields();
        setModalVisible(true);
    };

    const handleEdit = (record: Parameter) => {
        setEditingParam(record);
        form.setFieldsValue(record);
        setModalVisible(true);
    };

    const handleDelete = async (id: number) => {
        try {
            await deleteParameter(id);
            message.success('删除成功');
            fetchData();
        } catch {
            message.error('删除失败');
        }
    };

    const handleSubmit = async () => {
        try {
            const values = await form.validateFields();
            if (editingParam) {
                await updateParameter({ id: editingParam.id, ...values });
                message.success('更新成功');
            } else {
                await createParameter(values);
                message.success('创建成功');
            }
            setModalVisible(false);
            fetchData();
        } catch {
            message.error('操作失败');
        }
    };

    const columns: ColumnsType<Parameter> = [
        {
            title: 'ID',
            dataIndex: 'id',
            key: 'id',
            width: 80,
        },
        {
            title: '参数配置',
            dataIndex: 'parameters',
            key: 'parameters',
            ellipsis: true,
        },
        {
            title: '脚本内容',
            dataIndex: 'script',
            key: 'script',
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
            render: (_, record: Parameter) => (
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
                <h1 style={{ margin: 0 }}>参数/脚本管理</h1>
                <Space>
                    <Button icon={<ReloadOutlined />} onClick={fetchData} loading={loading}>
                        刷新
                    </Button>
                    <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>
                        新增配置
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
                title={editingParam ? '编辑配置' : '新增配置'}
                open={modalVisible}
                onOk={handleSubmit}
                onCancel={() => setModalVisible(false)}
                width={700}
            >
                <Form form={form} layout="vertical">
                    <Form.Item name="id" label="ID" rules={[{ required: true, message: '请输入ID' }]}>
                        <InputNumber style={{ width: '100%' }} disabled={!!editingParam} />
                    </Form.Item>
                    <Form.Item name="parameters" label="参数配置(JSON)">
                        <Input.TextArea rows={4} />
                    </Form.Item>
                    <Form.Item name="script" label="脚本内容" rules={[{ required: true, message: '请输入脚本内容' }]}>
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

export default ParamScriptManagement;
