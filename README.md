# 📌 TECHNICAL TEST - [BTS.ID](http://BTS.ID)

## **To-Do List API Documentation**

API ini memungkinkan pengguna untuk mengelola daftar tugas (To-Do List) dengan autentikasi JWT.

To-Do List ini dibangun menggunakan **Golang, Fiber, Gorm, Bcrypt,** dan **JWT (Json Web Token)** untuk autentikasi.

---

## 🔑 **Autentikasi**

### **Register User**

- **Endpoint:** `POST /users/register`
- **Request Body:**

```json
{
  "nama": "Jhon",
  "username": "brian",
  "password": "password123"
}
```

### **Login User**

- **Endpoint:** `POST /users/login`
- **Request Body:**

```json
{
  "username": "brian",
  "password": "password123"
}
```

---

## 📋 **To-Do List API**

```
Authorization: Bearer <your_jwt_token>
```

### **Buat To-Do List**

- **Endpoint:** `POST /api/todo/list`
- **Request Body:**

```json
{
  "title": "Belajar Golang",
  "description": "Menyelesaikan project dengan Fiber"
}
```

### **Dapatkan Semua To-Do List**
```
- **Endpoint:** `GET /api/todo/lists`
```

### **Detail Checklist (Berisi Item-Item To-Do)**
```
- **Endpoint:** `GET /api/todo/list/:id`

```


## ✅ **To-Do API**

```
Authorization: Bearer <your_jwt_token>
```

### **Tambah To-Do ke dalam To-Do List**

- **Endpoint:** `POST /api/todos`
- **Request Body:**

```json
{
  "todo_list_id": 1,
  "task": "Baca dokumentasi Fiber",
  "completed": false
}
```

### **Detail Item (To-Do Item)**

- **Endpoint:** `GET /api/todo/:id`

### **Edit To-Do**

- **Endpoint:** `PUT /api/todos/:id`
- **Request Body:**

```json
{
  "task": "Baca dokumentasi Fiber lebih mendalam",
  "completed": false
}
```

### **Hapus To-Do**

- **Endpoint:** `DELETE /api/todos/:id`

### **Tandai To-Do Sebagai Selesai**

- **Endpoint:** `PATCH /api/todos/:id/complete`

