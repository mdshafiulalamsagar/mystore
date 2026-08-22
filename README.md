# myStore - Cloud-Based Store Management System API

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14%2B-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![JWT Auth](https://img.shields.io/badge/JWT-Protected-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white)](https://jwt.io/)

**myStore** is a robust, production-ready, RESTful SaaS Backend API designed for small-to-medium retail businesses and shop owners. It enables shop managers to digitize inventory tracking, income/expense management, pending customer orders, and shop dues in real-time.

---

## Key Features

- **Authentication & Security:** Secure Signup & Login with Bcrypt password hashing and JWT token authentication.
- **Inventory Management:** Product catalog tracking with stock counts, unit prices, and low-stock threshold monitoring.
- **Financial Tracking:** Real-time logging of Income and Expenses with automated Net Profit calculations.
- **Task & Order Management:** Order deadline tracking and customer request status management.
- **Dues Management:** Rent, utility bill, and vendor account balance ledger tracking.
- **Clean Architecture:** Scalable monorepo layout separating domain models, database abstractions, controllers, and middlewares.

---

## Tech Stack

- **Backend Language:** Go (Golang)
- **Database:** PostgreSQL
- **Security:** JWT (JSON Web Tokens) & Bcrypt
- **Architecture:** Monorepo Layout (Go HTTP standard library with modular handlers)

---

## Project Architecture

```text
mystore/
├── backend/
│   ├── config/         # Environment setup
│   ├── controllers/    # API Request Logic (Auth, Inventory, Financials, Tasks, Dues)
│   ├── database/       # DB Pool initialization & SQL Schema scripts
│   ├── middleware/     # JWT Authentication Middleware
│   ├── models/         # Go Structural Data Mappings
│   ├── routes/         # REST API Endpoint Mappings
│   ├── utils/          # JWT generation & Bcrypt Hashing utilities
│   ├── .env            # Environment Variables (Ignored in Git)
│   ├── go.mod          # Go Module dependencies
│   └── main.go         # Application Entry Point
└── frontend/           # Planned React.js / Web Application
```

---

## Getting Started

### Prerequisites

- Go (v1.20 or later)
- PostgreSQL

### Setup Instructions

1. Clone the Repository:
   git clone https://github.com/mdshafiulalamsagar/mystore.git
   cd mystore/backend

2. Configure Environment Variables:
   Create a .env file inside the backend directory:
   PORT=8080
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_postgres_password
   DB_NAME=mystore_db

3. Set Up PostgreSQL Database Schema:
   sudo -u postgres psql -d mystore_db -f database/schema.sql

4. Run the Application:
   go run main.go

The backend server will run on http://localhost:8080.

---

## API Endpoints Documentation

| Category | Method | Endpoint | Auth Required | Description |
| :--- | :--- | :--- | :---: | :--- |
| **Auth** | `POST` | `/api/signup` | ❌ | Register a new shop owner account |
| **Auth** | `POST` | `/api/login` | ❌ | Authenticate user & receive JWT token |
| **Inventory** | `POST` | `/api/inventory/add` | ✅ | Add a new stock item |
| **Inventory** | `GET` | `/api/inventory/list` | ✅ | List all inventory items |
| **Financials** | `POST` | `/api/transactions/add` | ✅ | Record income or expense |
| **Financials** | `GET` | `/api/transactions/list` | ✅ | List all transactions |
| **Financials** | `GET` | `/api/transactions/summary` | ✅ | Fetch total income, expense & net profit |
| **Tasks** | `POST` | `/api/tasks/add` | ✅ | Add customer order or shop task |
| **Tasks** | `GET` | `/api/tasks/list` | ✅ | Fetch all tasks |
| **Dues** | `POST` | `/api/dues/add` | ✅ | Log shop rent, bills, or vendor dues |
| **Dues** | `GET` | `/api/dues/list` | ✅ | Fetch all recorded dues |

---

## Developer & Author

**Md Shafiul Alam Sagar**
- GitHub: https://github.com/mdshafiulalamsagar

---
*Built with passion and Go!*