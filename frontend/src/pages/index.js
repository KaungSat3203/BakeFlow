export default function Home() {
  return (
    <div className="min-h-screen p-10 bg-gray-100">
      <h1 className="text-3xl font-bold">Bakeflow Dashboard</h1>
      <p className="text-gray-600 mt-2">Incoming orders will appear here.</p>

      <div className="mt-10 p-6 bg-white rounded-xl shadow">
        <h2 className="text-xl font-semibold mb-3">Orders</h2>
        <p>No orders yet... (waiting for backend)</p>
      </div>
    </div>
  );
}
