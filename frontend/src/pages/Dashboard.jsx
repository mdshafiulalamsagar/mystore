import { useState, useEffect, useContext } from 'react';
import API from '../api/axios';
import { AuthContext } from '../context/AuthContext';
import { DollarSign, TrendingUp, TrendingDown, LogOut, Package, CheckSquare, FileText, Home } from 'lucide-react';
import { Link } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

export default function Dashboard() {
  const { logout } = useContext(AuthContext);
  const [summary, setSummary] = useState({ total_income: 0, total_expense: 0, net_profit: 0 });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchSummary();
  }, []);

  const fetchSummary = async () => {
    try {
      const res = await API.get('/transactions/summary');
      setSummary(res.data);
    } catch (err) {
      console.error('Failed to fetch financial summary', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex">
      {/* Sidebar Navigation */}
      <Sidebar />

      {/* Main Content Area */}
      <main className="flex-1 p-8">
        <header className="mb-8 flex justify-between items-center">
          <div>
            <h2 className="text-3xl font-bold">Store Financial Overview</h2>
          </div>
        </header>

        {loading ? (
          <p className="text-gray-400">Loading financial data...</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
            {/* Total Income Card */}
            <div className="bg-gray-800 p-6 rounded-xl border border-gray-700 shadow-lg">
              <div className="flex justify-between items-center mb-4">
                <span className="text-gray-400 font-medium">Total Income</span>
                <div className="p-2 bg-green-500/10 rounded-lg">
                  <TrendingUp className="text-green-400" size={20} />
                </div>
              </div>
              <p className="text-3xl font-bold text-green-400">৳{summary.total_income}</p>
            </div>

            {/* Total Expense Card */}
            <div className="bg-gray-800 p-6 rounded-xl border border-gray-700 shadow-lg">
              <div className="flex justify-between items-center mb-4">
                <span className="text-gray-400 font-medium">Total Expense</span>
                <div className="p-2 bg-red-500/10 rounded-lg">
                  <TrendingDown className="text-red-400" size={20} />
                </div>
              </div>
              <p className="text-3xl font-bold text-red-400">৳{summary.total_expense}</p>
            </div>

            {/* Net Profit Card */}
            <div className="bg-gray-800 p-6 rounded-xl border border-gray-700 shadow-lg">
              <div className="flex justify-between items-center mb-4">
                <span className="text-gray-400 font-medium">Net Profit</span>
                <div className="p-2 bg-blue-500/10 rounded-lg">
                  <DollarSign className="text-blue-400" size={20} />
                </div>
              </div>
              <p className={`text-3xl font-bold ${summary.net_profit >= 0 ? 'text-blue-400' : 'text-red-400'}`}>
                ৳{summary.net_profit}
              </p>
            </div>
          </div>
        )}
      </main>
    </div>
  );
}