import { useState, useEffect } from 'react';
import API from '../api/axios';
import { Home, Package, DollarSign, CheckSquare, FileText, Plus, CheckCircle2, Clock } from 'lucide-react';
import { Link } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

export default function Dues() {
  const [dues, setDues] = useState([]);
  const [customerName, setCustomerName] = useState('');
  const [phone, setPhone] = useState('');
  const [totalDue, setTotalDue] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchDues();
  }, []);

  const fetchDues = async () => {
    try {
      const res = await API.get('/dues');
      setDues(res.data || []);
    } catch (err) {
      console.error('Failed to fetch dues', err);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateDue = async (e) => {
    e.preventDefault();
    try {
      await API.post('/dues', {
        customer_name: customerName,
        phone,
        total_due: parseFloat(totalDue),
      });
      setCustomerName('');
      setPhone('');
      setTotalDue('');
      fetchDues();
    } catch (err) {
      alert('Failed to record due');
    }
  };

  const handlePayDue = async (id) => {
    const payInput = prompt('Enter payment amount (৳):');
    if (!payInput) return;
    const amount = parseFloat(payInput);
    if (isNaN(amount) || amount <= 0) {
      alert('Invalid amount');
      return;
    }

    try {
      await API.put(`/dues/${id}`, { amount });
      fetchDues();
    } catch (err) {
      alert('Failed to process payment');
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex">
      {/* Sidebar */}
      <Sidebar />

      {/* Main Content */}
      <main className="flex-1 p-8">
        <h2 className="text-3xl font-bold mb-6">Customer Dues Tracker (বাকি হিসাব)</h2>

        {/* New Due Form */}
        <form onSubmit={handleCreateDue} className="bg-gray-800 p-6 rounded-xl border border-gray-700 mb-8 flex gap-4 items-end">
          <div className="flex-1">
            <label className="block text-sm mb-1 text-gray-400">Customer Name</label>
            <input
              type="text"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={customerName}
              onChange={(e) => setCustomerName(e.target.value)}
            />
          </div>
          <div className="w-48">
            <label className="block text-sm mb-1 text-gray-400">Phone</label>
            <input
              type="text"
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
            />
          </div>
          <div className="w-36">
            <label className="block text-sm mb-1 text-gray-400">Total Due (৳)</label>
            <input
              type="number"
              step="0.01"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={totalDue}
              onChange={(e) => setTotalDue(e.target.value)}
            />
          </div>
          <button type="submit" className="bg-red-600 hover:bg-red-700 text-white px-5 py-2.5 rounded flex items-center gap-2 font-medium transition">
            <Plus size={18} /> Add Due
          </button>
        </form>

        {/* Dues Table */}
        <div className="bg-gray-800 rounded-xl border border-gray-700 overflow-hidden">
          <table className="w-full text-left">
            <thead className="bg-gray-750 text-gray-400 text-sm border-b border-gray-700">
              <tr>
                <th className="p-4">Customer</th>
                <th className="p-4">Phone</th>
                <th className="p-4">Total Due</th>
                <th className="p-4">Paid</th>
                <th className="p-4">Remaining Due</th>
                <th className="p-4">Status</th>
                <th className="p-4 text-right">Action</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-700">
              {loading ? (
                <tr><td colSpan="7" className="p-4 text-center text-gray-400">Loading dues...</td></tr>
              ) : dues.length === 0 ? (
                <tr><td colSpan="7" className="p-4 text-center text-gray-400">No dues recorded yet.</td></tr>
              ) : (
                dues.map((d) => {
                  const remaining = d.total_due - d.paid_amount;
                  return (
                    <tr key={d.id} className="hover:bg-gray-750">
                      <td className="p-4 font-medium">{d.customer_name}</td>
                      <td className="p-4 text-gray-300">{d.phone || 'N/A'}</td>
                      <td className="p-4 text-gray-300">৳{d.total_due}</td>
                      <td className="p-4 text-green-400">৳{d.paid_amount}</td>
                      <td className="p-4 font-bold text-red-400">৳{remaining.toFixed(2)}</td>
                      <td className="p-4">
                        <span className={`px-2.5 py-1 rounded-full text-xs font-semibold flex items-center gap-1 w-max ${
                          d.status === 'Paid' ? 'bg-green-500/10 text-green-400' :
                          d.status === 'Partial' ? 'bg-yellow-500/10 text-yellow-400' :
                          'bg-red-500/10 text-red-400'
                        }`}>
                          {d.status === 'Paid' ? <CheckCircle2 size={14} /> : <Clock size={14} />}
                          {d.status}
                        </span>
                      </td>
                      <td className="p-4 text-right">
                        {d.status !== 'Paid' && (
                          <button
                            onClick={() => handlePayDue(d.id)}
                            className="bg-blue-600 hover:bg-blue-700 text-white text-xs px-3 py-1.5 rounded transition"
                          >
                            Collect Pay
                          </button>
                        )}
                      </td>
                    </tr>
                  );
                })
              )}
            </tbody>
          </table>
        </div>
      </main>
    </div>
  );
}