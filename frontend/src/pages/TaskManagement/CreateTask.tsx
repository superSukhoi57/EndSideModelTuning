import React, { useEffect, useState } from 'react';
import { Form, Select, Slider, InputNumber, Input, Button, message, Space } from 'antd';
import { useNavigate } from 'react-router-dom';
import { useMachine } from '../../hooks/useMachine.ts';
import { useParameter } from '../../hooks/useParameter.ts';
import { taskService } from '../../services/taskService.ts';

interface Machine {
    id: number;
    ip: string;
}

interface Parameter {
    id: number;
    desc: string;
}

const CreateTask: React.FC = () => {
    const [form] = Form.useForm();
    const navigate = useNavigate();
    const { listData: machineData, listMachines } = useMachine();
    const { listData: paramData, listParameters } = useParameter();
    const [submitting, setSubmitting] = useState(false);

    useEffect(() => {
        listMachines({ page: 1, pageSize: 1000 });
        listParameters({ page: 1, pageSize: 1000 });
    }, [listMachines, listParameters]);

    const handleSubmit = async () => {
        try {
            const values = await form.validateFields();
            setSubmitting(true);
            await taskService.create(values);
            message.success('创建任务成功');
            navigate('/tasks/list');
        } catch {
            message.error('创建任务失败');
        } finally {
            setSubmitting(false);
        }
    };

    const handleCancel = () => {
        navigate('/tasks/list');
    };

    return (
        <div className="page-content">
            <h1>创建任务</h1>
            <Form
                form={form}
                layout="vertical"
                style={{ maxWidth: 600 }}
                initialValues={{
                    memory_percent: 50,
                    cpu_percent: 50,
                }}
            >
                <Form.Item
                    name="machineid"
                    label="选择机器"
                    rules={[{ required: true, message: '请选择机器' }]}
                >
                    <Select
                        placeholder="请选择机器"
                        options={machineData?.list?.map((m: Machine) => ({
                            value: m.id,
                            label: `ID: ${m.id} - IP: ${m.ip}`,
                        }))}
                    />
                </Form.Item>

                <Form.Item
                    name="paramterid"
                    label="选择参数"
                    rules={[{ required: true, message: '请选择参数' }]}
                >
                    <Select
                        placeholder="请选择参数"
                        options={paramData?.list?.map((p: Parameter) => ({
                            value: p.id,
                            label: `ID: ${p.id} - ${p.desc || '无描述'}`,
                        }))}
                    />
                </Form.Item>

                <Form.Item
                    name="memory_percent"
                    label="内存占用 (%)"
                    rules={[{ required: true, message: '请设置内存占用' }]}
                >
                    <Slider min={0} max={100} marks={{ 0: '0%', 50: '50%', 100: '100%' }} />
                </Form.Item>

                <Form.Item
                    name="cpu_percent"
                    label="CPU使用率 (%)"
                    rules={[{ required: true, message: '请设置CPU使用率' }]}
                >
                    <Slider min={0} max={100} marks={{ 0: '0%', 50: '50%', 100: '100%' }} />
                </Form.Item>

                <Form.Item
                    name="completion_time"
                    label="完成时间 (秒)"
                >
                    <InputNumber
                        style={{ width: '100%' }}
                        min={0}
                        step={1}
                        precision={0}
                        placeholder="请输入完成时间"
                    />
                </Form.Item>

                <Form.Item
                    name="limit"
                    label="最大迭代次数"
                >
                    <InputNumber
                        style={{ width: '100%' }}
                        min={1}
                        step={1}
                        precision={0}
                        placeholder="请输入最大迭代次数"
                    />
                </Form.Item>

                <Form.Item name="desc" label="任务描述">
                    <Input.TextArea rows={3} placeholder="请输入任务描述" />
                </Form.Item>

                <Form.Item>
                    <Space>
                        <Button type="primary" onClick={handleSubmit} loading={submitting}>
                            创建任务
                        </Button>
                        <Button onClick={handleCancel}>
                            取消
                        </Button>
                    </Space>
                </Form.Item>
            </Form>
        </div>
    );
};

export default CreateTask;
