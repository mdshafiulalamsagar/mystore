import { useState, useEffect } from 'react';
import API from '../api/axios';
import { Home, Package, DollarSign, CheckSquare, Plus, Clock, CheckCircle, XCircle } from 'lucide-react';
import { Link } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

export default function Orders() {
  const [orders, setOrders] = useState([]);
  const [customerName, setCustomerName] = useState('');
  const [customerPhone, setCustomerPhone] = useState('');
  const [totalAmount, setTotalAmount] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {
      const res = await API.get('/orders');
      setOrders(res.data || []);
    } catch (err) {
      console.error('Failed to fetch orders', err);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateOrder = async (e) => {
    e.preventDefault();
    try {
      await API.post('/orders', {
        customer_name: customerName,
        customer_phone: customerPhone,
        total_amount: parseFloat(totalAmount),
        status: 'Pending',
      });
      setCustomerName('');
      setCustomerPhone('');
      setTotalAmount('');
      fetchOrders();
    } catch (err) {
      alert('Failed to create order');
    }
  };

  const handleStatusChange = async (id, newStatus) => {
    try {
      await API.put(`/orders/${id}`, { status: newStatus });
      fetchOrders();
    } catch (err) {
      alert('Failed to update status');
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex">
      {/* Sidebar */}
      <Sidebar />

      {/* Main Content */}
      <main className="flex-1 p-8">
        <h2 className="text-3xl font-bold mb-6">Customer Orders & Tasks</h2>

        {/* New Order Form */}
        <form onSubmit={handleCreateOrder} className="bg-gray-800 p-6 rounded-xl border border-gray-700 mb-8 flex gap-4 items-end">
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
              value={customerPhone}
              onChange={(e) => setCustomerPhone(e.target.value)}
            />
          </div>
          <div className="w-36">
            <label className="block text-sm mb-1 text-gray-400">Amount (৳)</label>
            <input
              type="number"
              step="0.01"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={totalAmount}
              onChange={(e) => setTotalAmount(e.target.value)}
            />
          </div>
          <button type="submit" className="bg-blue-600 hover:bg-blue-700 text-white px-5 py-2.5 rounded flex items-center gap-2 font-medium transition">
            <Plus size={18} /> New Order
          </button>
        </form>

        {/* Orders Table */}
        <div className="bg-gray-800 rounded-xl border border-gray-700 overflow-hidden">
          <table className="w-full text-left">
            <thead className="bg-gray-750 text-gray-400 text-sm border-b border-gray-700">
              <tr>
                <th className="p-4">Order ID</th>
                <th className="p-4">Customer</th>
                <th className="p-4">Phone</th>
                <th className="p-4">Amount</th>
                <th className="p-4">Status</th>
                <th className="p-4 text-right">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-700">
              {loading ? (
                <tr><td colSpan="6" className="p-4 text-center text-gray-400">Loading orders...</td></tr>
              ) : orders.length === 0 ? (
                <tr><td colSpan="6" className="p-4 text-center text-gray-400">No orders found. Create one above!</td></tr>
              ) : (
                orders.map((o) => (
                  <tr key={o.id} className="hover:bg-gray-750">
                    <td className="p-4 text-gray-400">#{o.id}</td>
                    <td className="p-4 font-medium">{o.customer_name}</td>
                    <td className="p-4 text-gray-300">{o.customer_phone || 'N/A'}</td>
                    <td className="p-4 font-semibold text-blue-400">৳{o.total_amount}</td>
                    <td className="p-4">
                      <span className={`px-2.5 py-1 rounded-full text-xs font-semibold flex items-center gap-1 w-max ${
                        o.status === 'Delivered' ? 'bg-green-500/10 text-green-400' :
                        o.status === 'Cancelled' ? 'bg-red-500/10 text-red-400' :
                        'bg-yellow-500/10 text-yellow-400'
                      }`}>
                        {o.status === 'Delivered' && <CheckCircle size={14} />}
                        {o.status === 'Cancelled' && <XCircle size={14} />}
                        {o.status === 'Pending' && <Clock size={14} />}
                        {o.status}
                      </span>
                    </td>
                    <td className="p-4 text-right">
                      <select
                        value={o.status}
                        onChange={(e) => handleStatusChange(o.id, e.target.value)}
                        className="bg-gray-700 text-sm rounded p-1 border border-gray-600 focus:outline-none"
                      >
                        <option value="Pending">Pending</option>
                        <option value="Delivered">Delivered</option>
                        <option value="Cancelled">Cancelled</option>
                      </select>
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