// src/axiosConfig.ts
import axios from 'axios';

// Create an Axios instance
const axiosInstance = axios.create();

// Add a request interceptor to include the token in headers
axiosInstance.interceptors.request.use(
    (config) => {
        // Get token from localStorage
        const token = localStorage.getItem('token');

        // If token exists, attach it to the Authorization header
        if (true) {
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