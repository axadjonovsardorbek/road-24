import { useState } from 'react';
import Search from 'antd/es/input/Search';
import './style.scss';
import { Button, Drawer, Form, Input, message, Select, Table, Popconfirm, Switch } from 'antd';
import { CopyOutlined, EditOutlined, PlusCircleOutlined, ShareAltOutlined } from '@ant-design/icons';
import {
    EmailShareButton,
    TelegramShareButton,
    TelegramIcon,
    EmailIcon,
} from 'react-share';
import { useCreateUsers } from '../../services/users/mutation/useCreateUsers';
import { useGetUsers } from '../../services/users/mutation/useGetUsers';
import CryptoJS from 'crypto-js';

export const Users = () => {
    const [visible, setVisible] = useState(false);
    const [createdUserPassword, setCreatedUserPassword] = useState<string | null>(null);

    const { data: users, refetch, isLoading } = useGetUsers();

    const { mutate: createUsers } = useCreateUsers();

    const [form] = Form.useForm();

    const unhashPassword = (hashedPassword: string, originalPassword: string) => {
        const hashedOriginal = CryptoJS.SHA256(originalPassword).toString();
        return hashedOriginal === hashedPassword;
    };

    const handleCopy = (username: string, password: string) => {
        const textToCopy = `username: ${username}\npassword: ${password}`;
        navigator.clipboard.writeText(textToCopy).then(() => {
            message.success('Copied to clipboard!');
        }).catch(() => {
            message.error('Failed to copy!');
        });
    };


    const columns = [
        {
            title: 'First Name',
            dataIndex: 'first_name',
            key: 'firstName',
        },
        {
            title: 'Last Name',
            dataIndex: 'last_name',
            key: 'lastName',
        },
        {
            title: 'User Role',
            dataIndex: 'role',
            key: 'user_role',
        },
        {
            title: "User Login",
            dataIndex: "username",
            key: "username",
            render: (username: string) => (
                <div className='user-login-container'>
                    <Input type="text" className='user-login' readOnly value={username} />
                </div>
            )
        },
        {
            title: "User Password",
            key: "password",
            render: (hashedPassword: string, record: any) => {
                const unhashedPassword = unhashPassword(hashedPassword, record.password); 
                return (
                    <div className='user-password-container'>
                        <Input.Password type="text" className='user-password' readOnly value={unhashedPassword} />
                    </div>
                );
            }
        },
        {
            title: "Share",
            dataIndex: "share_login_password",
            key: "share_login_password",
            // eslint-disable-next-line
            render: (_: any, record: any) => (
                <div className='share-buttons-container'>
                    <Button onClick={() => handleCopy(record.username, record.password)} className='copy-btn' type='link'>
                        <CopyOutlined />
                    </Button>
                    <Popconfirm
                        title="Share by"
                        icon={false}
                        placement="top"
                        okButtonProps={{ style: { display: 'none' } }}
                        cancelButtonProps={{ style: { display: 'none' } }}
                        description={(
                            <div className="share-buttons">
                                <EmailShareButton
                                    url={`http://localhost:3000/login`}
                                    subject="Your Login and Password"
                                    body={`\n\nLogin: ${record.username}\nPassword: ${record.password}`}
                                >
                                    <EmailIcon size={32} round />
                                </EmailShareButton>
                                <TelegramShareButton
                                    url={'http://localhost:3000/login'}
                                    title={`Your Login and Password: \n\nLogin: ${record.username}\nPassword: ${record.password}`}
                                >
                                    <TelegramIcon size={32} round />
                                </TelegramShareButton>
                            </div>
                        )}
                    >
                        <Button className='copy-btn' type="link">
                            <ShareAltOutlined />
                        </Button>
                    </Popconfirm>

                </div>
            )
        },
        {
            title: "Actions",
            dataIndex: "actions",
            key: "actions",
            render: () => (
                <Button type="primary" icon={<EditOutlined />}>Edit</Button>
            )
        }
    ];

    const userRoleOptions = [
        {
            value: 'yhxb',
            label: 'YHXB'
        },
        {
            value: 'yidxp',
            label: 'YIDXP'
        },
        {
            value: 'service',
            label: 'Service'
        }
    ];

    const handleOpen = () => {
        setVisible(true);
    };

    const handleClose = () => {
        setVisible(false);
        form.resetFields();
    };

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const onFinish = (values: any) => {
        setCreatedUserPassword(values.password);
        createUsers({
            ...values,
            role: values.role,
            is_admin: values.is_admin ? "true" : "false",
        }, {
            onSuccess: () => {
                refetch();
                message.success('User created successfully!');
                form.resetFields();
            }
        });
        handleClose();
    };

    const updatedUsers = createdUserPassword
        ? [...(users?.users || []), { ...form.getFieldsValue(), password: createdUserPassword }]
        : users?.users || [];

    return (
        <>
            <div className='users'>
                <div className='users-header'>
                    <Search className='users-search' placeholder="Search Users..." enterButton />
                    <Button onClick={handleOpen} type="primary" icon={<PlusCircleOutlined />}>
                        Add User
                    </Button>
                </div>
                <div className='users-table-container'>
                    <Table loading={isLoading} pagination={{ defaultPageSize: 7, showSizeChanger: false }} className='users-table' dataSource={updatedUsers} columns={columns} />
                </div>
            </div>
            <Drawer open={visible} onClose={handleClose} placement='right'>
                <Form onFinish={onFinish} form={form} layout='vertical'>
                    <Form.Item label="First Name" name="first_name" rules={[{ required: true, message: 'Please input user first name!' }]}>
                        <Input type="text" />
                    </Form.Item>
                    <Form.Item label="Last Name" name="last_name" rules={[{ required: true, message: 'Please input user last name!' }]}>
                        <Input type="text" />
                    </Form.Item>
                    <Form.Item label="User Role" name="role" rules={[{ required: true, message: 'Please select user role!' }]}>
                        <Select options={userRoleOptions} />
                    </Form.Item>
                    <Form.Item label="Username" name="username" rules={[{ required: true, message: 'Please input username!' }]}>
                        <Input type="text" />
                    </Form.Item>
                    <Form.Item label="Password" name="password" rules={[{ required: true, message: 'Please input user password!' }]}>
                        <Input.Password />
                    </Form.Item>
                    <Form.Item label="Admin" name="is_admin" valuePropName="checked">
                        <Switch />
                    </Form.Item>
                    <Form.Item>
                        <Button className='submit-btn' type="primary" htmlType="submit">
                            SUBMIT
                        </Button>
                    </Form.Item>
                </Form>
            </Drawer>
        </>
    )
}
