import axios from 'axios';

const API = axios.create({
  baseURL: 'https://mystore-sx6n.onrender.com',
});

// Automatically add JWT Bearer token to requests
API.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default API;