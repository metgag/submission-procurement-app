# Procurement System (Sistem Pengadaan Barang)

Aplikasi sistem pengadaan barang yang digunakan untuk mencatat transaksi pembelian barang dari supplier, dengan perhitungan harga dan stok yang dikontrol penuh oleh Backend.

## Table of Contents

* [Tech Stack](#tech-stack)
* [Features](#features)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Configuration](#configuration)
* [Running the Application](#running-the-application)
* [API Documentation](#api-documentation)
* [Database Schema](#database-schema)
* [Role-based Access](#role-based-access)
* [Makefile Instructions](#makefile-instructions)
* [Troubleshooting](#troubleshooting)

---

## Tech Stack

### Backend

* **Language**: Go 1.24.11 atau lebih
* **Framework**: Go Fiber
* **ORM**: GORM
* **Database**: PostgreSQL 16.11 atau lebih
* **Authentication**: JWT

### Frontend

* **Library**: jQuery
* **Styling**: Tailwind CSS (CDN)

---

## Features

### Core Features

* **Authentication**: Register & Login dengan JWT Token
* **Master Data**: CRUD Items & Suppliers (Admin only)
* **Purchase Transaction**: Transaksi pembelian dengan perhitungan otomatis di backend
* **Inventory Management**: Tracking stok barang real-time
* **Role-based Access Control**: User & Admin roles

### Bonus Features

* **Database Transaction (ACID)**: Atomic operations untuk data consistency
* **Event Delegation**: Dynamic cart management
* **Reusable AJAX Wrapper**: Centralized API calls dengan auto-authentication
* **Robust Error Handling**: Toast notifications untuk user feedback

---

## Prerequisites

Pastikan sudah terinstall:

* Go 1.24 atau lebih
* Docker untuk running PostgreSQL atau PostgreSQL Client
* Git

---

## Installation

### 1. Clone Repository

```bash
git clone https://github.com/metgag/submission-procurement-app
cd submission-procurement-app
```

### 2. Setup Database (via Docker)

```bash
docker run --name postgres-procurement \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=procurement \
  -p 5436:5432 \
  -d postgres:16.11
```

---

## Configuration

### 1. Create Environment File

Buat file `.env` di root directory project:

```env
PG_DSN="host=localhost user=postgres password=postgres dbname=procurement port=5436 sslmode=disable TimeZone=Asia/Jakarta"
JWT_SECRET="a-string-secret-at-least-256-bits-long"
```

### 2. Environment Variables

| Variable     | Description                             | Example                                                                                                             |
| ------------ | --------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `PG_DSN`     | PostgreSQL connection string            | `host=localhost user=postgres password=postgres dbname=procurement port=5436 sslmode=disable TimeZone=Asia/Jakarta` |
| `JWT_SECRET` | Secret key untuk JWT (minimal 256 bits) | `a-string-secret-at-least-256-bits-long`                                                                            |

---

## Running the Application

### 1. Install Dependencies

```bash
make install
```

### 2. Jalankan PostgreSQL

```bash
make docker-up
```

### 3. Jalankan Migration

```bash
make migrate
```

* Membuat tabel-tabel yang dibutuhkan.
* Jika tabel sudah ada, migrasi akan menyesuaikan struktur tabel sesuai model GORM.

### 4. Seed Database (Optional)

```bash
make db-seed
```

Menambahkan sample data ke database untuk testing.

### 5. Run Backend Server

```bash
make run
```

Server akan berjalan di `http://localhost:3080`

**Note**:

* Table database akan dibuat secara otomatis saat pertama kali running
* Data masih kosong, perlu registrasi user terlebih dahulu

### 6. Register User

**Endpoint**: `POST http://localhost:3080/auth/register`

**Request Body**:

```json
{
  "username": "admin",
  "password": "admin123"
}
```

### 7. Access Frontend

* Login Page: `http://localhost:3080/`
* Dashboard: `http://localhost:3080/dashboard` (setelah login)
* Purchase: `http://localhost:3080/purchase` (setelah login)

---

## API Documentation

### Authentication

#### Register

```http
POST /auth/register
Content-Type: application/json

{
  "username": "user1",
  "password": "secret"
}
```

### Master Data (Protected Routes)

#### Get Items

```http
GET /items
Authorization: Bearer <token>
```

#### Get Suppliers

```http
GET /suppliers
Authorization: Bearer <token>
```

#### Get Supplier Items

```http
GET /supplier-items
Authorization: Bearer <token>
```

### Purchase Transaction

#### Create Purchase

```http
POST /purchase
Authorization: Bearer <token>
Content-Type: application/json

{
  "supplier_id": 1,
  "items": [
    {
      "supplier_item_id": 1,
      "quantity": 2
    }
  ]
}
```

**Validation Rules**:

* `username` > 3 karakter
* `password` > 8 karakter
* `items` tidak boleh kosong
* `quantity` > 0
* `stock` >= `quantity`
* `supplier_item_id` harus valid

---

## Database Schema

### Users

* `id` (Primary Key)
* `username` (Unique)
* `password` (Hashed)
* `role` (user/admin)

### Suppliers

* `id` (Primary Key)
* `name`
* `email`
* `address`

### Items

* `id` (Primary Key)
* `name`
* `created_at`

### Supplier Items

* `id` (Primary Key)
* `supplier_id` (Foreign Key)
* `item_id` (Foreign Key)
* `price`
* `stock`
* `created_at`

### Purchases

* `id` (Primary Key)
* `date`
* `supplier_id` (Foreign Key)
* `user_id` (Foreign Key)
* `grand_total`
* `created_at`

### Purchasing Details

* `id` (Primary Key)
* `purchasing_id` (Foreign Key)
* `supplier_item_id` (Foreign Key)
* `quantity`
* `subtotal`

---

## Role-based Access

### User Role

* Login & Logout
* View Dashboard (Inventory)
* Create Purchase Transaction

### Admin Role

* Semua akses User
* CRUD Items
* CRUD Suppliers

---

## Makefile Instructions

### Available Commands

| Command          | Description                                                       |
| ---------------- | ----------------------------------------------------------------- |
| `make install`   | Install semua dependencies Go (`go mod tidy` & `go mod download`) |
| `make run`       | Jalankan server Go (`go run ./cmd/server`)                        |
| `make docker-up` | Jalankan PostgreSQL container via Docker                          |
| `make migrate`   | Jalankan migration untuk membuat atau update table di database    |
| `make db-seed`   | Seed database dengan sample data                                  |
| `make clean`     | Membersihkan build cache dan folder `bin/`                        |

---

## Troubleshooting

### Dependencies Error / Module Not Found

```bash
go mod tidy
go mod download
```

### Database Connection Failed

* Pastikan PostgreSQL container sudah running: `docker ps`
* Cek konfigurasi port di `.env` sesuai dengan container
* Restart container: `docker restart postgres-procurement`

### JWT Token Invalid

* Pastikan `JWT_SECRET` di `.env` minimal 256 bits
* Clear browser localStorage dan login ulang

### Table Not Created

* Stop server dan jalankan ulang: `make migrate`
* Cek log error di terminal

### Build Error

```bash
go clean -cache
go mod tidy
go build ./cmd/server
```
