# myStore — Store & Financial Management System

A lightweight, full-stack store management and financial tracking web application built with **Go (Golang)**, **React.js (Vite)**, and **PostgreSQL**. Designed to help small-to-medium business owners effortlessly manage inventory, transactions, orders, and customer dues in real-time.

---

## Features

- **Financial Overview:** Real-time dashboard showing Total Income, Expenses, and Net Profit metrics.
- **Inventory Management:** Add, track, and update product stocks and prices dynamically.
- **Transaction Ledger:** Categorized logging for shop income and operational expenses.
- **Task & Order Tracker:** Manage customer orders and update status efficiently.
- **Dues Tracker (বাকি হিসাব):** Track customer debts, calculate remaining dues, process partial/full payments, and auto-update status (`Unpaid`, `Partial`, `Paid`).
- **JWT Authentication:** Secure registration and login flow with protected routes and state management.

---

## Tech Stack

### **Frontend**
- **Framework:** React.js (Vite)
- **Styling:** Tailwind CSS
- **Icons:** Lucide React
- **HTTP Client:** Axios
- **Deployment:** Vercel

### **Backend**
- **Language:** Go (Golang)
- **Database Driver:** `[github.com/lib/pq](https://github.com/lib/pq)`
- **Authentication:** JWT (JSON Web Tokens) & bcrypt hashing
- **Deployment:** Render

### **Database**
- **Database System:** Cloud PostgreSQL (Neon.tech)

---

## Project Structure

```text
mystore/
├── backend/
│   ├── config/           # App configurations
│   ├── controllers/      # API handlers (Auth, Inventory, Dues, etc.)
│   ├── database/         # Database connection & pooling
│   ├── middleware/       # JWT Auth & CORS handling
│   ├── models/           # Data models (Structs)
│   ├── routes/           # Endpoint definitions
│   └── main.go           # Application entry point
│
└── frontend/
    ├── src/
    │   ├── api/          # Axios instance & interceptors
    │   ├── components/   # Reusable UI components (Sidebar, Protected Routes)
    │   ├── context/      # React Auth Context
    │   ├── pages/        # App pages (Dashboard, Inventory, Dues, etc.)
    │   ├── App.jsx       # Route setups
    │   └── main.jsx      # React entry point
    ├── index.html
    └── vite.config.js
```

---

## Getting Started Locally

### **Prerequisites**
- [Go 1.20+](https://go.dev/dl/) installed
- [Node.js 18+](https://nodejs.org/) installed
- PostgreSQL database instance running locally or via Cloud (Neon)

### **1. Backend Setup**
```bash
cd backend

# Install dependencies
go mod tidy

# Set environment variables (Optional fallback handles defaults)
export DATABASE_URL="postgres://user:password@localhost:5432/mystore_db?sslmode=disable"
export PORT=8080

# Run the server
go run main.go
```
The backend server will run at `http://localhost:8080`.

### **2. Frontend Setup**
```bash
cd frontend

# Install dependencies
npm install

# Start Vite dev server
npm run dev
```
The frontend application will run at `http://localhost:5173`.

---

## Live Deployment Architecture

| Tier | Provider | Description |
| :--- | :--- | :--- |
| **Frontend** | Vercel | Global CDN deployment for ultra-fast React client delivery. |
| **Backend** | Render | Dockerless Go environment handling API endpoints. |
| **Database** | Neon.tech | Serverless Cloud PostgreSQL instance. |

---

## License

This project is open-source and available under the [MIT License](LICENSE).