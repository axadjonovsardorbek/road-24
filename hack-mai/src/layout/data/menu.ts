import { AppstoreAddOutlined, PieChartOutlined, UsergroupAddOutlined } from "@ant-design/icons";
import { nanoid } from "@reduxjs/toolkit";
import React from "react";

export const menu = [
  {
    id: nanoid(),
    name: "Dashboard",
    path: '/',
    icon: PieChartOutlined,
  },
  {
    id: nanoid(),
    name: "Hodimlar",
    icon: UsergroupAddOutlined,
    path: '/users'
  },
  {
    id: nanoid(),
    name: "Management",
    icon: AppstoreAddOutlined,
    children: [
      {
        id: nanoid(),
        name: "Mashinalar",
        path: '/cars',
        icon: PieChartOutlined,
      },
      {
        id: nanoid(),
        name: "Guvohnomalar",
        path: '/licenses',
        icon: PieChartOutlined,
      },
      {
        id: nanoid(),
        name: "Servislar",
        path: '/services',
        icon: PieChartOutlined,
      },
    ]
  },
];

export const menuItems = menu.map(item => {
  if (item.children) {
    return {
      key: item.id,
      icon: React.createElement(item.icon),
      label: item.name,
      children: item.children.map(subItem => ({
        key: subItem.id,
        label: subItem.name,
        path: subItem.path,
      })),
    };
  } else {
    return {
      key: item.id,
      icon: React.createElement(item.icon),
      label: item.name,
      path: item.path,
    };
  }
});
