import axios, { AxiosInstance, AxiosError } from "axios";
import Cookies from "js-cookie";
const BASE_URL = import.meta.env.VITE_APP_TITLE; 

const request: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  }
});

request.interceptors.request.use(
  (config) => {
    const token = Cookies.get('token');
    if (token && config.url !== "/login") {
      config.headers['Authorization'] = `Bearer ${token}`; 
    }
    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

request.interceptors.response.use(
  (response) => {
    return response;
  },
  (error: AxiosError) => {
    if (error.response?.status === 401 && window.location.pathname !== "/login") {
      window.location.href = "/login"; 
    }
    return Promise.reject(error);
  }
);

export default request;
