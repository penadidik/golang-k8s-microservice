Berikut adalah file **README.md** untuk menjalankan **User Service** dalam arsitektur Clean Architecture dengan **Golang, Kubernetes, dan Docker**.  

---

## **User Service - Clean Architecture**
User Service adalah microservice untuk mengelola pengguna dalam sistem. Proyek ini dibangun dengan **Golang**, **Gin Framework**, dan menggunakan **PostgreSQL** sebagai database.

---

## **📌 Fitur**
- **Registrasi Pengguna** (`POST /users/register`)
- **Clean Architecture** sesuai dengan prinsip **SOLID**
- **Kubernetes-ready** dengan deployment manifest
- **Dockerized** untuk memudahkan deployment

---

## **🚀 Cara Menjalankan**

### **1️⃣ Clone Repository**
```sh
git clone https://github.com/your-repo/user-service.git
cd user-service
```

### **2️⃣ Setup Environment**
Buat file `.env` di root project:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=user_db
DB_PORT=5432
APP_PORT=8080
```

---

### **3️⃣ Jalankan Database PostgreSQL (Docker)**
Pastikan Anda memiliki **Docker** terinstal, lalu jalankan perintah berikut:
```sh
docker run --name user-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=user_db -p 5432:5432 -d postgres
```

---

### **4️⃣ Jalankan Aplikasi Secara Lokal**
Pastikan **Go** telah terinstal, lalu jalankan:
```sh
go mod tidy
go run cmd/main.go
```
Aplikasi akan berjalan di `http://localhost:8080`.

---

### **5️⃣ Testing API**
Gunakan `cURL` atau **Postman** untuk menguji API.

#### 🔹 **Registrasi User**
```sh
curl -X POST http://localhost:8080/users/register -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role_id": 1
}'
```
**Respon Sukses**
```json
{
  "message": "User created successfully!"
}
```

---

## **📦 Jalankan dengan Docker**
### **1️⃣ Build Docker Image**
```sh
docker build -t user-service .
```
### **2️⃣ Jalankan Container**
```sh
docker run -p 8080:8080 --env-file .env user-service
```

---

## **📌 Deploy ke Kubernetes**
Pastikan **kubectl** dan **minikube** sudah diinstal.

### **1️⃣ Start Minikube**
```sh
minikube start
```

### **2️⃣ Deploy ke Kubernetes**
```sh
kubectl apply -f k8s/user-service.yaml
```

### **3️⃣ Cek Status Pod**
```sh
kubectl get pods
```

### **4️⃣ Akses Service**
```sh
kubectl port-forward svc/user-service 8080:80
```
API dapat diakses di `http://localhost:8080`.

---

## **📜 Struktur Proyek**
```
user-service/
  ├── cmd/
  │   ├── main.go           # Entry point
  ├── config/
  │   ├── db.go             # Database configuration
  ├── internal/
  │   ├── entity/           # Domain models
  │   ├── repository/       # Data access layer
  │   ├── usecase/          # Business logic
  │   ├── handler/          # API handlers
  ├── infrastructure/
  │   ├── router.go         # HTTP router setup
  ├── k8s/
  │   ├── user-service.yaml # Kubernetes deployment file
  ├── Dockerfile            # Docker image build file
  ├── .env                  # Environment variables
  ├── go.mod                # Go module dependencies
```

---

## **📌 Teknologi yang Digunakan**
- **Golang** (Gin Framework)
- **PostgreSQL** (Database)
- **Docker** (Containerization)
- **Kubernetes** (Orchestration)
- **Clean Architecture** (SOLID Principles)

---

## **📢 TODO**
✅ Implementasi **User Service**  
✅ Docker & Kubernetes Deployment  
🔜 Tambahkan **Role Service & Permission Service**  
🔜 Gunakan **gRPC untuk komunikasi antar microservices**  
🔜 Implementasi **JWT Authentication**  

---

## **📞 Kontak & Kontribusi**
Jika ada pertanyaan atau ingin berkontribusi, silakan buat **Pull Request** atau hubungi saya di `penadidik@gmail.com`.

---


