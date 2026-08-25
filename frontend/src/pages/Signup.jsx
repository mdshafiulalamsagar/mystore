import { useState, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import { useNavigate, Link } from 'react-router-dom';

export default function Signup() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [shopName, setShopName] = useState('');
  const [error, setError] = useState('');
  const { signup } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await signup(name, email, password, shopName);
      navigate('/login');
    } catch (err) {
      setError(err.response?.data?.error || 'Signup failed');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <form onSubmit={handleSubmit} className="bg-gray-800 p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold mb-6 text-center text-green-400">Create Account</h2>
        {error && <p className="bg-red-500/20 text-red-400 p-2 rounded mb-4 text-sm">{error}</p>}

        <div className="mb-3">
          <label className="block text-sm mb-1">Full Name</label>
          <input
            type="text"
            required
            className="w-full p-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-green-500"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>

        <div className="mb-3">
          <label className="block text-sm mb-1">Shop Name</label>
          <input
            type="text"
            required
            className="w-full p-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-green-500"
            value={shopName}
            onChange={(e) => setShopName(e.target.value)}
          />
        </div>

        <div className="mb-3">
          <label className="block text-sm mb-1">Email</label>
          <input
            type="email"
            required
            className="w-full p-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-green-500"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>

        <div className="mb-6">
          <label className="block text-sm mb-1">Password</label>
          <input
            type="password"
            required
            className="w-full p-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-green-500"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <button type="submit" className="w-full bg-green-600 hover:bg-green-700 text-white p-2 rounded font-semibold transition">
          Sign Up
        </button>

        <p className="mt-4 text-sm text-center text-gray-400">
          Already have an account? <Link to="/login" className="text-green-400 underline">Login</Link>
        </p>
      </form>
    </div>
  );
}