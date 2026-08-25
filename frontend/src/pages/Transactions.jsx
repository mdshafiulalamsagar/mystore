import { useState, useEffect } from 'react';
import API from '../api/axios';
import { Home, Package, DollarSign, Plus, ArrowUpCircle, ArrowDownCircle } from 'lucide-react';
import { Link } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

export default function Transactions() {
  const [transactions, setTransactions] = useState([]);
  const [type, setType] = useState('income');
  const [amount, setAmount] = useState('');
  const [description, setDescription] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchTransactions();
  }, []);

  const fetchTransactions = async () => {
    try {
      const res = await API.get('/transactions');
      setTransactions(res.data || []);
    } catch (err) {
      console.error('Failed to fetch transactions', err);
    } finally {
      setLoading(false);
    }
  };

  const handleAddTransaction = async (e) => {
    e.preventDefault();
    try {
      await API.post('/transactions', {
        type,
        amount: parseFloat(amount),
        description,
      });
      setAmount('');
      setDescription('');
      fetchTransactions();
    } catch (err) {
      alert('Failed to add transaction');
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex">
      {/* Sidebar */}
      <Sidebar />

      {/* Main Content */}
      <main className="flex-1 p-8">
        <h2 className="text-3xl font-bold mb-6">Income & Expense Transactions</h2>

        {/* Add Transaction Form */}
        <form onSubmit={handleAddTransaction} className="bg-gray-800 p-6 rounded-xl border border-gray-700 mb-8 flex gap-4 items-end">
          <div className="w-40">
            <label className="block text-sm mb-1 text-gray-400">Type</label>
            <select
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={type}
              onChange={(e) => setType(e.target.value)}
            >
              <option value="income">Income (+)</option>
              <option value="expense">Expense (-)</option>
            </select>
          </div>
          <div className="w-36">
            <label className="block text-sm mb-1 text-gray-400">Amount (৳)</label>
            <input
              type="number"
              step="0.01"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
            />
          </div>
          <div className="flex-1">
            <label className="block text-sm mb-1 text-gray-400">Description</label>
            <input
              type="text"
              placeholder="e.g. Sold 2 Inks / Rent payment"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
          </div>
          <button type="submit" className="bg-blue-600 hover:bg-blue-700 text-white px-5 py-2.5 rounded flex items-center gap-2 font-medium transition">
            <Plus size={18} /> Record
          </button>
        </form>

        {/* Transactions Table */}
        <div className="bg-gray-800 rounded-xl border border-gray-700 overflow-hidden">
          <table className="w-full text-left">
            <thead className="bg-gray-750 text-gray-400 text-sm border-b border-gray-700">
              <tr>
                <th className="p-4">Type</th>
                <th className="p-4">Description</th>
                <th className="p-4">Amount</th>
                <th className="p-4">Date</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-700">
              {loading ? (
                <tr><td colSpan="4" className="p-4 text-center text-gray-400">Loading history...</td></tr>
              ) : transactions.length === 0 ? (
                <tr><td colSpan="4" className="p-4 text-center text-gray-400">No transactions recorded yet.</td></tr>
              ) : (
                transactions.map((t) => (
                  <tr key={t.id} className="hover:bg-gray-750">
                    <td className="p-4 flex items-center gap-2 font-medium">
                      {t.type === 'income' ? (
                        <span className="flex items-center gap-1 text-green-400"><ArrowUpCircle size={16} /> Income</span>
                      ) : (
                        <span className="flex items-center gap-1 text-red-400"><ArrowDownCircle size={16} /> Expense</span>
                      )}
                    </td>
                    <td className="p-4 text-gray-200">{t.description}</td>
                    <td className={`p-4 font-bold ${t.type === 'income' ? 'text-green-400' : 'text-red-400'}`}>
                      {t.type === 'income' ? '+' : '-'}৳{t.amount}
                    </td>
                    <td className="p-4 text-gray-400 text-sm">
                      {new Date(t.created_at).toLocaleDateString()}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </main>
    </div>
  );
}