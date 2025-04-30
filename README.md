# Job Recommendation System

Sistem rekomendasi pekerjaan berbasis microservices (Go, gRPC, REST, NGINX, Docker Compose).

## Fitur

- **Auth Service**: Otentikasi JWT, login/register, gRPC & REST API
- **CV Service**: Upload & ekstraksi data CV, rekomendasi pekerjaan
- **API Gateway**: NGINX sebagai reverse proxy & CORS handler
- **Database**: MySQL
- **Dev Tools**: Hot reload, PHPMyAdmin, TypeScript check, GitHub Actions CI

---

## Prasyarat

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go 1.24+](https://go.dev/doc/install)

---

## Cara Install & Jalankan (Development)

1. **Clone repository**

   ```bash
   git clone https://github.com/agusheryanto182/job-recommendations.git
   cd job-recommendations
   ```

2. **Copy file environment**

   ```bash
   cp .env.example .env
   # Edit .env sesuai kebutuhan (MySQL, JWT, dsb)
   ```

3. **Jalankan semua service**

   ```bash
   cd docker
   docker-compose up --build
   ```

4. **Akses service**
   - API Gateway: [http://localhost:8000](http://localhost:8000)
   - Auth Service: [http://localhost:8000/auth/health](http://localhost:8000/auth/health)
   - CV Service: [http://localhost:8000/cv/health](http://localhost:8000/cv/health)
   - PHPMyAdmin: [http://localhost:8080](http://localhost:8080)
   - Frontend: [http://localhost:3000](http://localhost:3000)

---

## Struktur Project

```
backend/
  ├── auth-service/      # Service otentikasi (Go, gRPC, REST)
  ├── cv-service/        # Service CV & rekomendasi (Go, gRPC, REST)
  └── proto/             # Shared proto files (gRPC)
docker/
  ├── docker-compose.yml # Orkestrasi semua service
  └── nginx/             # Konfigurasi NGINX (API Gateway)
frontend/                # Frontend React/Next.js
scrapper/                # Scrape data set dari Linkedin
model/                   # Model rekomendasi pekerjaan menggunakan content-based filtering
```

---

## Endpoints

### Auth Service

- `GET /user/health`
- `GET /user/google/login`
- `GET /user/google/callback`
- `GET /user/profile`
- `GET /user/refresh`
- `GET /user/logout`

### CV Service

- `POST /cv/extract` (upload CV, extract data)
- `GET /cv/history` (riwayat rekomendasi)
- `GET /cv/health`

---

## Testing

- **TypeScript & Lint**:  
  Otomatis di-check via GitHub Actions setiap push/PR.
- **Go Build & Test**:  
  Otomatis di-check via GitHub Actions setiap push/PR.

---

## Tips Pengembangan

- Untuk development, gunakan `Dockerfile.dev` (hot reload, volume mounting)
- Untuk production, gunakan `Dockerfile.prod` (multi-stage, image kecil)
- Semua service saling komunikasi via gRPC/REST, proto file di-share di `backend/proto`
- Gunakan NGINX sebagai satu-satunya entry point (port 8000)

---

## Kontribusi

1. Fork repo ini
2. Buat branch baru (`git checkout -b fitur-baru`)
3. Commit perubahan (`git commit -am 'Tambah fitur baru'`)
4. Push ke branch (`git push origin fitur-baru`)
5. Buat Pull Request
6. Request review ke agusheryanto182

---

## Lisensi

MIT License

---

## Troubleshooting

- **Build context terlalu besar?**  
  Pastikan `.dockerignore` sudah benar di root project.
- **gRPC error import proto?**  
  Pastikan replace path di `go.mod` sesuai dengan environment (Docker/CI).
- **Service tidak bisa diakses?**  
  Cek network di `docker-compose.yml` dan pastikan semua service join network yang sama.

---

## Author

- [Agus Heryanto](https://github.com/agusheryanto182)
