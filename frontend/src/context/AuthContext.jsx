import { createContext, useState } from 'react';
import API from '../api/axios';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(localStorage.getItem('token') || '');

  const login = async (email, password) => {
    const res = await API.post('/login', { email, password });
    const userToken = res.data.token;
    localStorage.setItem('token', userToken);
    setToken(userToken);
    return res.data;
  };

  const signup = async (name, email, password, shopName) => {
    const res = await API.post('/signup', {
      name,
      email,
      password,
      shop_name: shopName,
    });
    return res.data;
  };

  const logout = () => {
    localStorage.removeItem('token');
    setToken('');
  };

  return (
    <AuthContext.Provider value={{ token, login, signup, logout }}>
      {children}
    </AuthContext.Provider>
  );
};