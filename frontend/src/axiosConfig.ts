// src/axiosConfig.ts
import axios from 'axios';

const baseURL: string = import.meta.env.VITE_API_URL

// Create an Axios instance
const axiosInstance = axios.create({ baseURL });

// Add a request interceptor to include the token in headers
axiosInstance.interceptors.request.use(
    (config) => {
        // Get token from localStorage
        const token = localStorage.getItem('token');

        // If token exists, attach it to the Authorization header
        if (token !== null) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }

        return config;
    },
    (error) => {
        // Handle any error in the request setup
        return Promise.reject(error);
    }
);

export default axiosInstance;