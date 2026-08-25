import { Link, useLocation, useNavigate } from 'react-router-dom';
import { Home, Package, DollarSign, CheckSquare, FileText, LogOut } from 'lucide-react';

export default function Sidebar() {
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  const navItems = [
    { name: 'Overview', path: '/dashboard', icon: Home },
    { name: 'Inventory', path: '/inventory', icon: Package },
    { name: 'Transactions', path: '/transactions', icon: DollarSign },
    { name: 'Tasks / Orders', path: '/orders', icon: CheckSquare },
    { name: 'Dues', path: '/dues', icon: FileText },
  ];

  return (
    <aside className="w-64 bg-gray-800 p-6 flex flex-col justify-between border-r border-gray-700 min-h-screen">
      <div>
        <h1 className="text-2xl font-bold text-blue-400 mb-8 tracking-wide">myStore</h1>
        <nav className="space-y-2">
          {navItems.map((item) => {
            const Icon = item.icon;
            const isActive = location.pathname === item.path;
            return (
              <Link
                key={item.path}
                to={item.path}
                className={`flex items-center space-x-3 p-2.5 rounded-lg font-medium transition ${
                  isActive
                    ? 'bg-blue-600 text-white'
                    : 'text-gray-400 hover:text-white hover:bg-gray-700'
                }`}
              >
                <Icon size={18} />
                <span>{item.name}</span>
              </Link>
            );
          })}
        </nav>
      </div>

      <button
        onClick={handleLogout}
        className="flex items-center space-x-3 text-red-400 hover:text-red-300 hover:bg-gray-700/50 w-full p-2.5 rounded-lg font-medium transition"
      >
        <LogOut size={18} />
        <span>Logout</span>
      </button>
    </aside>
  );
}