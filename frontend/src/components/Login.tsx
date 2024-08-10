import React, { useState } from 'react';
import { login } from '../services/authService';

interface LoginProps {
  onClose: () => void;
  onLogin: () => void;
}

const Login: React.FC<LoginProps> = ({ onClose, onLogin }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    try {
      await login(email, password);
      onLogin();
    } catch (err) {
      setError('Invalid credentials');
    }
  };

  return (
    <div className="max-w-md mx-auto bg-gray-800 p-8 rounded-lg">
      <h2 className="text-3xl font-bold mb-6 text-center text-green-300">Login</h2>
      {error && <p className="text-red-500 mb-4">{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="email" className="block text-green-300 mb-2">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-3 py-2 bg-gray-700 text-white rounded focus:outline-none focus:ring-2 focus:ring-green-300"
            required
          />
        </div>
        <div className="mb-6">
          <label htmlFor="password" className="block text-green-300 mb-2">Password</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-3 py-2 bg-gray-700 text-white rounded focus:outline-none focus:ring-2 focus:ring-green-300"
            required
          />
        </div>
        <div className="flex justify-between items-center">
          <button
            type="submit"
            className="bg-green-300 text-black px-4 py-2 rounded hover:bg-green-400 transition-colors"
          >
            Sign In
          </button>
          <button
            type="button"
            onClick={onClose}
            className="text-green-300 hover:text-green-400"
          >
            Cancel
          </button>
        </div>
      </form>
      <p className="mt-4 text-sm text-gray-400">
        Hint: Use test@test.com / test to log in
      </p>
    </div>
  );
};

export default Login;