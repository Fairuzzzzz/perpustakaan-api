## Perpustakaan API

Perpustakaan API adalah RESTful API sederhana yang menyediakan fungsionalitas untuk mengelola data perpustakaan,
seperti buku, anggota, dan transaksi peminjaman.

#### Fitur
- Mengelola data buku (tambah, lihat, update, hapus)
- Mengelola data anggota perpustakaan
- Mengelola peminjaman dan pengembalian buku

### Konfigurasi
Konfigurasi aplikasi dapat ditemukan `perpustakaan_api/internal/configs/config.yaml`.
Berikut contoh konfigurasi:

```yaml
service:
  port: ":8888"
  secretJWT: "fairuztest"

database:
  dataSourceName: "root:testpassword@tcp(127.0.0.1:3306)/perpustakaan?parseTime=true"
```

#### Instalasi Proyek
1. Clone repository
```bash
git clone https://github.com/Fairuzzzzz/perpustakaan-api.git
cd perpustakaan-api
```

2. Instalasi Docker
```bash
docker-compose up -d
```

3. Database Migration
```bash
make migrate-up
```

4. Jalankan secara lokal
```bash
go run cmd/main.go
```

### API Endpoints
Berikut adalah daftar endpoint yang tersedia dalam Perpustakaan API :

#### Buku

| HTTP Method | Endpoint                  | Deskripsi                                | Autentikasi | Otorisasi |
|-------------|----------------------------|------------------------------------------|-------------|-----------|
| POST        | `/books/add-book`          | Menambahkan buku baru                    | Ya          | Admin     |
| DELETE      | `/books/delete-book`       | Menghapus buku                           | Ya          | Admin     |
| PUT         | `/books/update-book`       | Memperbarui informasi buku               | Ya          | Admin     |
| GET         | `/books/`                  | Mendapatkan daftar semua buku            | Ya          | Admin     |
| POST        | `/books/borrow-book`       | Meminjam buku                            | Ya          | Anggota   |
| POST        | `/books/return-book`       | Mengembalikan buku                       | Ya          | Anggota   |
| GET         | `/books/borrowed-book`     | Mendapatkan daftar buku yang dipinjam    | Ya          | Admin     |

#### Keanggotaan

| HTTP Method | Endpoint                    | Deskripsi                                  | Autentikasi | Otorisasi |
|-------------|------------------------------|--------------------------------------------|-------------|-----------|
| POST        | `/membership/sign-up`        | Mendaftar anggota baru                     | Tidak       | -         |
| POST        | `/membership/login`          | Login anggota                              | Tidak       | -         |
| DELETE      | `/membership/delete-user`    | Menghapus anggota                          | Ya          | Admin     |
| GET         | `/membership/`               | Mendapatkan daftar semua anggota           | Ya          | Admin     |
| GET         | `/membership/:userID/borrow-history`| Mendapatkan riwayat peminjaman buku | Ya          | Admin     |

### Contoh Request dan Response

#### Menambahkan Buku
**Request**
```http
POST /books/add-book
Content-Type: application/json
Authorization: Bearer <token>

{
  "title": "Billy Bat",
  "author": "Naoki Urasawa",
  "category": ["Mystery", "Sci-fi"],
  "publicationYear": "2008-10-16",
  "totalCopies": 50
}
```

**Response**
```http
HTTP/1.1 201 Created
```

#### Login Anggota
**Request**
```http
POST /membership/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response**
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "accessToken": "eyJHasdUasdfijasdfJ..."
}
```
