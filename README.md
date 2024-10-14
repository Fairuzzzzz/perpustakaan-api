### Perpustakaan API

Perpustakaan API adalah RESTful API sederhana yang menyediakan fungsionalitas untuk mengelola data perpustakaan,
seperti buku, anggota, dan transaksi peminjaman.

#### Fitur
- Mengelola data buku (tambah, lihat, update, hapus)
- Mengelola data anggota perpustakaan
- Mengelola peminjaman dan pengembalian buku

#### Instalasi Proyek
1. Clone repository
```bash
git clone https://github.com/Fairuzzzzz/perpustakaan-api.git
cd perpustakaan-api
```

2. Jalankan secara lokal
```bash
go run cmd/main.go
```

3. Instalasi Docker
```bash
docker-compose up -d
```

4. Database Migration
```bash
make migrate-create
make migrate-up
```
