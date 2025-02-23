# Golang Backend Project Structure

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)![Postgresql](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

## 1. Struktur Direktori dan Penjelasannya

#### **Root Directory**
- Berisi direktori utama yang mengorganisir kode berdasarkan fungsionalitasnya.
- File konfigurasi seperti .env dan .air.toml.

#### **cmd/**
- Berisi entry point untuk menjalankan aplikasi dan perintah tambahan.
- **api/**: Direktori ini digunakan untuk menjalankan aplikasi API utama.
- **seeders/**: Berisi script untuk memasukkan data awal ke dalam database.

#### **config/**
- Berisi file konfigurasi yang mengatur berbagai aspek aplikasi seperti koneksi database, environment variables, dll.

#### **internal/**
Berisi seluruh implementasi internal aplikasi yang terdiri dari beberapa sub-direktori:
  
  - **constant/**: Berisi konstanta yang digunakan oleh aplikasi.
  - **database/**
    - **migrations/**: Berisi file migrasi database menggunakan Goose.
    - **seeders/**: Berisi script untuk memasukkan data awal ke database.

  - **middleware/**
    Berisi middleware yang digunakan dalam aplikasi, seperti autentikasi, logging, dan lainnya.

  - **modules/**
    Berisi implementasi setiap fitur aplikasi yang dipisahkan berdasarkan modul. 
  - **(nama modul)**
    
    - **http/**
      - **handler/**: Berisi fungsi handler untuk menangani request, memanggil funsi di service layer, dan mengembalikan response.
      - **route/**: Berisi definisi rute atau endpoint.
    - **model/**: Berisi definisi model, entity, dto (request dan response).
    - **service/**: Berisi logika bisnis (business logic).
    - **repository/**: Berisi fungsi untuk berinteraksi dengan database, seperti operasi Create, Read, Update, Delete, atau operasi DML lainnya.


  - **routes/**
    Berisi daftar rute utama dari seluruh modul aplikasi.

#### **pkg/**
Berisi kode yang dapat digunakan kembali di berbagai bagian aplikasi.
  - **error/**: Berisi definisi error custom.
  - **logger/**: Berisi konfigurasi logging.
  - **success/**: Berisi format respons sukses.
  - **utils/**: Berisi fungsi-fungsi utilitas.
  - **validation/**: Berisi fungsi validasi input.

---

## 2. Cara Menjalankan Project dengan Air

[Air](https://github.com/air-verse/air) adalah live reloading tool untuk Go. Dengan Air, setiap perubahan dalam kode akan langsung diterapkan tanpa perlu restart manual.

### **Instalasi Air**
Jika belum terinstal, gunakan perintah berikut:
```sh
go install github.com/air-verse/air@latest
```

### **Menjalankan Project dengan Air**
1. Pastikan berada di root direktori project.
2. Jalankan perintah berikut:
   ```sh
   air
   ```
3. Aplikasi akan berjalan dengan live reload setiap kali ada perubahan kode.

---

## 3. Goose untuk Database Migration

[Goose](https://github.com/pressly/goose) digunakan untuk mengelola skema database melalui migrasi.

### **Instalasi Goose**
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### **Membuat File Migrasi**
Untuk membuat file migrasi baru:
```sh
goose create nama_migrasi sql
```
Ini akan menghasilkan file di direktori `internal/database/migrations/` dengan format:
```
20250222120000_nama_migrasi.sql
```

### **Menjalankan Migrasi**
Menjalankan semua migrasi yang belum diterapkan:
```sh
goose up
```

### **Rollback Migrasi**
Untuk membatalkan migrasi terakhir:
```sh
goose down
```

### **Melihat Status Migrasi**
Untuk melihat daftar migrasi yang telah dijalankan:
```sh
goose status
```

Dengan struktur yang terorganisir dan penggunaan Air serta Goose, pengembangan backend menjadi lebih efisien dan terstruktur.

