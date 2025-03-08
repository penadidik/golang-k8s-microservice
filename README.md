Here's a **simple README** for your **Golang Microservices with Clean Architecture, SOLID Principle, JWT Authentication, and Kubernetes Deployment** 🚀.

---

## 🔥 Microservices Architecture with Golang, JWT & Kubernetes
This project is a **microservice-based architecture** built using **Golang** with **Clean Architecture** and **SOLID Principles**. It consists of the following services:

- **User Service** (Register, Login, Get Profile)
- JWT Authentication
- MongoDB as Database
- Docker & Kubernetes for Deployment

---

### 📌 Tech Stack
- Golang
- MongoDB
- JWT Authentication
- Docker
- Kubernetes
- SOLID Principles
- Clean Architecture

---

## 🏗️ Project Structure
```bash
user-service
├── cmd               # App Bootstrap
├── config            # Configuration Files
├── internal          # Business Logic
│   ├── entity       # Entities (Models)
│   ├── repository   # Data Access Layer
│   ├── usecase      # Business Logic Layer
│   └── delivery     # HTTP Handlers
├── infrastructure    # Database & External Services
└── Dockerfile        # Docker Image
```

---

## ⚙️ How to Run Locally

### 1. Clone Repository
```bash
git clone https://github.com/yourusername/microservices-golang.git
cd microservices-golang
```

---

### 2. Setup MongoDB
Use Docker to run MongoDB locally:
```bash
docker run -d -p 27017:27017 --name mongo mongo:latest
```

---

### 3. Run the Service
```bash
go mod tidy
go run cmd/main.go
```

---

### 4. API Endpoints
| Endpoint   | Method | Description        |
|-----------|--------|-------------------|
| `/register` | POST   | Register User     |
| `/login`    | POST   | Login User       |
| `/profile`  | GET    | Get Profile (JWT Required) |

---

## 🔑 Example Request
### Register User
```bash
POST /register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@mail.com",
  "password": "password123"
}
```

---

### Login User
```bash
POST /login
Content-Type: application/json

{
  "email": "john@mail.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "JWT_TOKEN_HERE"
}
```

---

### Get Profile
```bash
GET /profile
Authorization: Bearer JWT_TOKEN_HERE
```

---

## 🐳 Docker Commands
### Build Docker Image
```bash
docker build -t user-service .
```

---

### Run Docker
```bash
docker run -d -p 8081:8081 user-service
```

---

## ☸️ Deploy to Kubernetes
1. Apply Deployment & Service
```bash
kubectl apply -f deployment.yml
kubectl apply -f service.yml
```

2. Access the Service
```bash
http://localhost:30081/register
```

---

## 🎯 Environment Variables
| Variable      | Description       |
|--------------|------------------|
| `MONGO_URI`  | MongoDB Connection String |
| `JWT_SECRET` | JWT Secret Key    |

---

## 📄 License
MIT License

---

---

✅ Now your Microservices are ready for production 🚀🔥  

Do you need **Swagger API Documentation + Postman Collection**? 📄✨
