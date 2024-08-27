import React, { useState } from 'react';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  ExclamationCircleFilled,
  LogoutOutlined,
} from '@ant-design/icons';
import { Layout, Menu, Modal, message, Button } from 'antd';
import './style.scss';
import { Outlet, useNavigate } from 'react-router-dom';
import { menuItems } from './data/menu';
import Cookies from 'js-cookie';

const { confirm } = Modal;
const { Header, Sider, Content } = Layout;

export const MainLayout: React.FC = () => {
  const navigate = useNavigate();
  const token = Cookies.get('token');

  if (!token) {
    window.location.replace("/login");
  }


  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const onClickMenuItem = (item: any) => {
    if (item.item.props.path) {
      navigate(item.item.props.path);
    }
  };

  const showConfirm = () => {
    confirm({
      title: 'Do you want to log out?',
      icon: <ExclamationCircleFilled />,
      cancelText: 'No',
      onOk() {
        Cookies.remove("token");
        navigate("/login");
        message.success("You logged out successfully");
      }
    });
  };

  const [collapsed, setCollapsed] = useState(false);

  const toggleCollapse = () => {
    setCollapsed(!collapsed);
  };

  return (
    <Layout className='layout'>
      <Sider collapsedWidth={70} className='sidebar' trigger={null} collapsible collapsed={collapsed}>
        <div className='logo-block' style={{ padding: '16px', textAlign: 'center' }}>
          {collapsed ? <h3 className='logo-title' >TM</h3> : <h3 className='logo-title' >TechMinds</h3>}
        </div>
        <Menu
          className='layout__menu'
          theme="dark"
          mode="inline"
          items={menuItems}
          onClick={onClickMenuItem}
        />
        <div onClick={showConfirm} style={{ position: 'absolute', cursor: 'pointer', bottom: '0', width: '100%', padding: '16px', textAlign: 'left', display: 'flex', color: '#fff', alignItems: 'center', justifyContent: 'start', gap: '8px', paddingLeft: '24px' }}>
          <LogoutOutlined style={{ fontSize: 18, color: "#fff" }} /> {collapsed ? '' : 'Log out'}
        </div>
      </Sider>
      <Layout>
        <Header style={{ padding: 0 }}>
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined style={{ color: '#fff' }} /> : <MenuFoldOutlined style={{ color: '#fff' }} />}
            onClick={toggleCollapse}
            style={{
              fontSize: '16px',
              width: 64,
              height: 64,
            }}
          />
        </Header>
        <Content className='content'>
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};
