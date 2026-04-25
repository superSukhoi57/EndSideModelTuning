import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Form, Input, InputNumber, Popconfirm, message, Space, Tag } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined, ReloadOutlined } from '@ant-design/icons';
import { useMachine } from '../hooks/useMachine.ts';
import type { ColumnsType } from 'antd/es/table';
import type { Machine } from '../services/types.ts';

const DeviceManagement: React.FC = () => {
    const { loading, listData, createMachine, updateMachine, deleteMachine, listMachines } = useMachine();
    const [form] = Form.useForm();
    const [modalVisible, setModalVisible] = useState(false);
    const [editingMachine, setEditingMachine] = useState<Machine | null>(null);
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);

    useEffect(() => {
        fetchData();
    }, [page, pageSize]);

    const fetchData = () => {
        listMachines({ page, pageSize });
    };

    const handleCreate = () => {
        setEditingMachine(null);
        form.resetFields();
        setModalVisible(true);
    };

    const handleEdit = (record: Machine) => {
        setEditingMachine(record);
        form.setFieldsValue(record);
        setModalVisible(true);
    };

    const handleDelete = async (id: number) => {
        try {
            await deleteMachine(id);
            message.success('删除成功');
            fetchData();
        } catch {
            message.error('删除失败');
        }
    };

    const handleSubmit = async () => {
        try {
            const values = await form.validateFields();
            if (editingMachine) {
                await updateMachine({ id: editingMachine.id, ...values });
                message.success('更新成功');
            } else {
                await createMachine(values);
                message.success('创建成功');
            }
            setModalVisible(false);
            fetchData();
        } catch {
            message.error('操作失败');
        }
    };

    const columns: ColumnsType<Machine> = [
        {
            title: 'IP地址',
            dataIndex: 'ip',
            key: 'ip',
        },
        {
            title: 'CPU核心数',
            dataIndex: 'core',
            key: 'core',
        },
        {
            title: '内存(MB)',
            dataIndex: 'ram',
            key: 'ram',
        },
        {
            title: '磁盘(GB)',
            dataIndex: 'memory',
            key: 'memory',
        },
        {
            title: '操作系统',
            dataIndex: 'os',
            key: 'os',
        },
        {
            title: '状态',
            dataIndex: 'isfinsh',
            key: 'isfinsh',
            render: (val: number) => (
                <Tag color={val === 1 ? 'green' : 'orange'}>
                    {val === 1 ? '已完成' : '进行中'}
                </Tag>
            ),
        },
        {
            title: '描述',
            dataIndex: 'desc',
            key: 'desc',
        },
        {
            title: '操作',
            key: 'action',
            render: (_, record: Machine) => (
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
                <h1 style={{ margin: 0 }}>设备管理</h1>
                <Space>
                    <Button icon={<ReloadOutlined />} onClick={fetchData} loading={loading}>
                        刷新
                    </Button>
                    <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>
                        新增设备
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
                title={editingMachine ? '编辑设备' : '新增设备'}
                open={modalVisible}
                onOk={handleSubmit}
                onCancel={() => setModalVisible(false)}
            >
                <Form form={form} layout="vertical">
                    <Form.Item name="ip" label="IP地址" rules={[{ required: true, message: '请输入IP地址' }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item name="pwd" label="密码">
                        <Input.Password />
                    </Form.Item>
                    <Form.Item name="userid" label="用户ID" rules={[{ required: true, message: '请输入用户ID' }]}>
                        <InputNumber style={{ width: '100%' }} />
                    </Form.Item>
                    <Form.Item name="core" label="CPU核心数">
                        <InputNumber style={{ width: '100%' }} />
                    </Form.Item>
                    <Form.Item name="ram" label="内存(MB)">
                        <InputNumber style={{ width: '100%' }} />
                    </Form.Item>
                    <Form.Item name="memory" label="磁盘(GB)">
                        <InputNumber style={{ width: '100%' }} />
                    </Form.Item>
                    <Form.Item name="os" label="操作系统">
                        <Input />
                    </Form.Item>
                    <Form.Item name="desc" label="描述">
                        <Input.TextArea />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    );
};

export default DeviceManagement;
