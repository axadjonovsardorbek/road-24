import { useMutation } from "@tanstack/react-query";
import request from "../../../config/request";

type UserType = {
    first_name: string;
    last_name: string;
    username: string;
    password: string;
    role: string;
    is_admin: string;
}

export const useCreateUsers = () => {
    return useMutation({
        mutationKey: ['add-users'],
        mutationFn: (data: UserType[]) => request.post<{token:string}>("/admin/user/add", data).then((res) => res.data)
    });
}