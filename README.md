<p align="center"><a href="https://mikti.id" target="_blank"><img src="https://mikti.id/assets/images/resources/logo-1.png" width="400" alt="Mikti Logo"></a></p>

# 🎫 DEPUBLIC TICKETING: CAPSTONE PROJECT MIKTI STUDI INDEPENDEN

## 🌻 Deskripsi
Depublic merupakan platform pembelian tiket. Melalui depublic user bisa membeli tiket event secara online tanpa perlu mengantri.
Depublic juga sudah terintegrasi dengan payment gateway sehingga mempermudah pembayaran dengan banyak metode.

## ⚡ Teknologi

- [Go v1.22](https://go.dev/) : Bahasa Pemrograman
- [Go Echo](https://echo.labstack.com/) : Framework
- [Gorm](https://gorm.io/docs/query.html) : ORM
- [Go Migrate](https://github.com/golang-migrate/migrate) : Migration
- [Gocraft](https://github.com/gocraft/work) : Background Job
- [PostgreSQL](https://www.postgresql.org/) : Database
- [Redis](https://redis.io/) : Cache
- [Midtrans](https://midtrans.com/) : Payment Gateway

## 🚩 Cara Menjalankan
1. Clone repository ini dengan perintah

```git
git clone https://github.com/bloomingbug/depublic.git
```
2. Masuk ke direktori aplikasi dengan perintah

```
cd depublic
```
3. Salin file .env.example menjadi .env dengan perintah

```
cp .env.example .env
```
4. Sesuaikan konfigurasi aplikasi, database, smtp, payment gateway, dll.  pada file .env sesuai dengan environment yang akan digunakan
5. Run Migration

```
make migrate
```

6. Jalankan database seeder

```
go run .\cmd\seeder\main.go
```
7. Jalankan background proccess

```
go run .\cmd\background\main.go
```
8. Jalankan aplikasi utama

```
go run .\cmd\server\main.go
```

## 👨‍💻 Kontributor

- [Tarmuji (Team Lead)](https://www.linkedin.com/in/tarmuji-tarmuji/)
- Hildiah Khairuniza (Scrum Master)
- Reza Bintang Suherman (Dokumentator)
- Ardy Surya Pratama (API Dokumentator)
- Dhie Ajeng Mia Anjari (Code Reviewer & QA)
