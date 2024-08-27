import { RoutesType } from "../types";
import { Users } from "../pages/users";
import { Cars } from "../pages/cars";
import { Lisences } from "../pages/lisences";

export const routes: RoutesType[] = [
    {
        id: '1',
        name: 'Users',
        path: '/users',
        component: Users
    },
    {
        id: '2',
        name: 'Cars',
        path: '/cars',
        component: Cars
    },
    {
        id: '3',
        name: 'Licenses',
        path: '/licenses',
        component: Lisences
    }
]