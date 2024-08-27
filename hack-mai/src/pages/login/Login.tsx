import React from 'react';
import { Form, Input, Button, Row, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import './style.scss';
import { useNavigate } from 'react-router-dom';
import { usePostUser } from '../../services/login';
import Cookies from 'js-cookie'

type FieldType = {
    username: string;
    password: string;
};
export const Login: React.FC = () => {


    const navigate = useNavigate();
    const { mutate, isPending } = usePostUser();

    React.useEffect(() => {
        if (Cookies.get("token")) {
            navigate('/', { replace: true })
        }
    }, [navigate])

    const onFinish = (values: FieldType) => {
        mutate(values, {
            onSuccess: (res) => {
                message.success("Login successfuly!")
                navigate("/", { replace: true });
                Cookies.set("token", res.token, { expires: 7 })
            },
            onError: () => {
                message.error("Incorrect Password or Username!");
            }
        });
    };

    return (
        <Row justify="center" align="middle" className="login-page">
            <Form
                className="login-form"
                initialValues={{
                    username: "sardorbekaxadjonov",
                    password: "1234"
                }}
                onFinish={onFinish}
                autoComplete='off'
            >
                <Form.Item
                    name="username"
                    className='input'
                    rules={[{ required: true, message: 'Please input your Username!' }]}
                >
                    <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
                </Form.Item>
                <Form.Item
                    className='input'
                    name="password"
                    rules={[{ required: true, message: 'Please input your Password!' }]}
                >
                    <Input
                        prefix={<LockOutlined className="site-form-item-icon" />}
                        type="password"
                        placeholder="Password"
                    />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" loading={isPending} htmlType="submit" className="login-form-button">
                        Login
                    </Button>
                </Form.Item>
            </Form>
        </Row>
    );
};
