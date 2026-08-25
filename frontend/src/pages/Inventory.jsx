import { useState, useEffect } from 'react';
import API from '../api/axios';
import { Package, Plus, Trash2, Home, DollarSign, CheckSquare, FileText, LogOut } from 'lucide-react';
import { Link, useNavigate } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

export default function Inventory() {
  const [products, setProducts] = useState([]);
  const [name, setName] = useState('');
  const [price, setPrice] = useState('');
  const [stock, setStock] = useState('');
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    try {
      const res = await API.get('/products');
      setProducts(res.data || []);
    } catch (err) {
      console.error('Failed to fetch products', err);
    } finally {
      setLoading(false);
    }
  };

  const handleAddProduct = async (e) => {
    e.preventDefault();
    try {
      await API.post('/products', {
        name,
        price: parseFloat(price),
        stock: parseInt(stock),
      });
      setName('');
      setPrice('');
      setStock('');
      fetchProducts();
    } catch (err) {
      alert('Failed to add product');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Delete this product?')) return;
    try {
      await API.delete(`/products/${id}`);
      fetchProducts();
    } catch (err) {
      alert('Failed to delete product');
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex">
      {/* Sidebar */}
      <Sidebar />

      {/* Content */}
      <main className="flex-1 p-8">
        <h2 className="text-3xl font-bold mb-6">Inventory Management</h2>

        {/* Add Product Form */}
        <form onSubmit={handleAddProduct} className="bg-gray-800 p-6 rounded-xl border border-gray-700 mb-8 flex gap-4 items-end">
          <div className="flex-1">
            <label className="block text-sm mb-1 text-gray-400">Product Name</label>
            <input
              type="text"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className="w-32">
            <label className="block text-sm mb-1 text-gray-400">Price (৳)</label>
            <input
              type="number"
              required
              step="0.01"
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={price}
              onChange={(e) => setPrice(e.target.value)}
            />
          </div>
          <div className="w-32">
            <label className="block text-sm mb-1 text-gray-400">Stock Qty</label>
            <input
              type="number"
              required
              className="w-full p-2.5 bg-gray-700 rounded border border-gray-600 focus:outline-none"
              value={stock}
              onChange={(e) => setStock(e.target.value)}
            />
          </div>
          <button type="submit" className="bg-green-600 hover:bg-green-700 text-white px-5 py-2.5 rounded flex items-center gap-2 font-medium transition">
            <Plus size={18} /> Add Product
          </button>
        </form>

        {/* Products Table */}
        <div className="bg-gray-800 rounded-xl border border-gray-700 overflow-hidden">
          <table className="w-full text-left">
            <thead className="bg-gray-750 text-gray-400 text-sm border-b border-gray-700">
              <tr>
                <th className="p-4">ID</th>
                <th className="p-4">Product Name</th>
                <th className="p-4">Price</th>
                <th className="p-4">Stock Status</th>
                <th className="p-4 text-right">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-700">
              {loading ? (
                <tr><td colSpan="5" className="p-4 text-center text-gray-400">Loading products...</td></tr>
              ) : products.length === 0 ? (
                <tr><td colSpan="5" className="p-4 text-center text-gray-400">No products found. Add one above!</td></tr>
              ) : (
                products.map((p) => (
                  <tr key={p.id} className="hover:bg-gray-750">
                    <td className="p-4 text-gray-400">#{p.id}</td>
                    <td className="p-4 font-medium">{p.name}</td>
                    <td className="p-4 text-green-400">৳{p.price}</td>
                    <td className="p-4">
                      <span className={`px-2.5 py-1 rounded-full text-xs font-semibold ${p.stock > 5 ? 'bg-green-500/10 text-green-400' : 'bg-red-500/10 text-red-400'}`}>
                        {p.stock} units
                      </span>
                    </td>
                    <td className="p-4 text-right">
                      <button onClick={() => handleDelete(p.id)} className="text-red-400 hover:text-red-300 p-1">
                        <Trash2 size={18} />
                      </button>
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