import { useMutation } from "@tanstack/react-query";
import request from "../../config/request";

type LoginType = {
    username: string;
    password: string;
}

export const usePostUser = () => {
    return useMutation({
        mutationKey: ['login'],
        mutationFn: (data:LoginType) => request.post<{token:string}>("/login", data).then((res) => res.data)
    });
}