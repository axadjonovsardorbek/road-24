export type UsersType = {
    id: number;
    first_name: string;
    last_name: string;
    username: string;
    company: string;
    department: string;
    is_active: boolean;
}

export type UpdateUserType = {
    id: number;
    first_name: string;
    last_name: string;
    username: string;
    company: number;
    department: number;
    is_active: boolean;
}