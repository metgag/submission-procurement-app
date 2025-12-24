# Procurement System

**(Sistem Pengadaan Barang)**

---

## 1. Tujuan Sistem

Aplikasi ini digunakan untuk mencatat transaksi pembelian barang dari supplier, dengan perhitungan harga dan stok yang dikontrol penuh oleh Backend.

---

## 2. Requirements

### 2.1 Core / Functional Requirements

#### Backend

##### Authentication

* Register & Login user
* Menggunakan **JWT Token**
* Middleware protected route

  * Hanya user login yang dapat membuat purchase
  * Role-based access control

##### Master Data

* CRUD **Items** (Admin)
* CRUD **Suppliers** (Admin)

##### Purchasing Transaction (Business Logic)

* Endpoint untuk membuat transaksi pembelian
* Backend **wajib menghitung**:

  * `subTotal`
  * `grandTotal`
* Harga item **tidak boleh dipercaya dari frontend**

---

#### Frontend

##### Login Page

* Form login terintegrasi API
* Token JWT disimpan di:

  * `localStorage` atau
  * `Cookie`

##### Dashboard & Inventory

* Menampilkan list item, dan nama suppliernya
* Menampilkan sisa stok
* Data diambil dari API

##### Create Purchase Page

1. User memilih **Supplier**
2. User memilih **Item** dan menginput **Qty**
3. Item masuk ke tabel sementara (**Cart**)
4. Data cart belum dikirim ke server
5. Tombol **Submit Order**:

   * Mengirim seluruh data cart ke backend
   * Dalam **1 request JSON**

---

### 2.2 Bonus / Non-Functional Requirements

#### Backend

##### Database Transaction (ACID)

Saat transaksi dibuat, backend harus menjalankan proses berikut secara **atomic**:

1. Insert Purchase Header
2. Insert Purchase Detail
3. Update / Potong Stock Item

Jika salah satu gagal, maka **seluruh proses di-rollback**.

##### External Integration (Webhook)

* Setelah transaksi sukses:

  * Kirim HTTP **POST (JSON)** ke URL dummy (Webhook.site / RequestBin)
  * Payload berisi detail order

---

#### Frontend

##### Event Delegation

* Tombol **Hapus Item** pada cart dibuat secara dinamis
* Harus tetap berfungsi dengan benar

##### Reusable AJAX

* Membuat wrapper AJAX
* Header `Authorization: Bearer <token>` tidak ditulis berulang

##### Robust Error Handling

* Menampilkan error yang informatif menggunakan:

  * Toast Notification
  * Modal Alert
* Contoh error:

  * Login gagal
  * Stok habis
  * Server error (500)
* **Dilarang menggunakan**:

  * `window.alert()`
  * `console.log()` sebagai feedback user

---

## 3. High-Level Architecture

```
Client (jQuery) → Server (Go Fiber API) → Database (PostgreSQL)
```

---

## 4. Tech Stack

### Backend

* Language: **Go**
* Framework: **Go Fiber**
* ORM: **GORM**
* Database: **PostgreSQL**
* Auth: **JWT**

### Frontend

* Library: **jQuery**
* Styling: **Tailwind CSS (CDN)**
* Constraint:

  * ❌ Tidak menggunakan framework JS modern (React, Vue, dll)

---

## 5. Detail Desain Backend

### 5.1 Server Application

* Go Fiber dengan environment configuration
* Password di-hash saat registrasi
* Validasi input basic
* JWT Payload:

  * Expired time
  * Username
  * Role

##### Role Authorization

* **User**

  * Create purchase
* **Admin**

  * CRUD Items
  * CRUD Suppliers

---

### 5.2 Database Schema

#### Users

* `id`
* `username`
* `password`
* `role`

#### Suppliers

* `id`
* `name`
* `email`
* `address`

#### Items

* `id`
* `name`
* `created_at`

#### Supplier Items

* `id`
* `supplier_id`
* `item_id`
* `price`
* `stock`
* `created_at`

#### Purchases

* `id`
* `date`
* `supplier_id`
* `user_id`
* `grand_total`
* `created_at`

#### Purchasing Details

* `id`
* `purchasing_id`
* `supplier_item_id`
* `quantity`
* `subtotal`

---

## 6. API Specification

### Authentication

#### POST `/auth/register`

Request:

```json
{
  "username": "user1",
  "password": "secret"
}
```

#### POST `/auth/login`

Response:

```json
{
  "message": "Login successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
  }
}
```

---

### Master Data

#### GET `/items`

* Menampilkan list item (protected)

#### GET `/suppliers`

* Menampilkan list supplier (protected)

#### GET `/supplier-items`

* Menampilkan list items dari supplier beserta stok dan harganya (protected)

---

### Purchasing Transaction

#### POST `/purchase`

Request:

```json
{
  "supplier_id": 1,
  "items": [
    {
      "supplier_item_id": 1,
      "quantity": 1
    }
  ]
}
```

Response:

```json
{
  "message": "Purchase created successfully",
  "data": {
    "id": 31,
    "grand_total": 370000,
    "items": [
      {
        "item_id": 2,
        "quantity": 1,
        "subtotal": 90000
      }
    ]
  }
}
```

---

### Validation Rules

* `username` > 3 karakter
* `password` > 8 karakter
* `items` tidak boleh kosong
* `quantity` > 0
* `stock >= quantity`
* `supplier_item_id` harus valid

---

## 7. Client Side Responsibilities

* Login via AJAX
* Menyimpan token
* Fetch data item & supplier
* Add / remove item dari cart tanpa refresh
* Menyusun JSON payload
* Submit ke API

---

## 8. Catatan Tambahan

* Fokus utama pada **logic dan clean code**
* UI sederhana, tidak perlu kompleks
* README wajib berisi:

  * Cara menjalankan backend
  * Env example
  * Cara setup database
