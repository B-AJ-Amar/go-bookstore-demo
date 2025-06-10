# 📚 go-bookstore-demo

A simple microservices-based bookstore application built with **Go**, for learning purposes :

- **Chi** for the RESTful public API
- Two services:
  - 📘 Bookstore Service (REST API + gRPC client)
  - 💰 Payment Service (gRPC server)
- **gRPC** for internal service communication

---

## 🧱 Architecture

```bash
[ Client ] <- REST (Chi) -> [ Bookstore Service ] <- gRPC -> [ Payment Service ]
```

---

## 🔧 Features

### ✅ Bookstore Service
- Add, list, update, and delete books
- Purchase a book (calls Payment Service via gRPC)

### 💰 Payment Service
- Exposes a `ProcessPayment` gRPC method
- Simulates or logs payment handling
- Sends back payment status

---

## 📁 Project Structure

```bash
go-bookstore-demo/
├── bookstore/
│   ├── main.go            # REST API entry
│   ├── router.go          # Chi router setup
│   ├── handlers/          # HTTP route logic
│   ├── models/            # Book, PurchaseRequest types
│   ├── repository/        # Book repository logic
│   ├── rpc/               # gRPC client logic
│   │   └── client.go
├── payment/
│   ├── main.go            # gRPC server entry
│   ├── service.go         # gRPC service implementation
│   ├── models/            # Shared payment models
├── proto/
│   └── payment.proto      # gRPC Protobuf definition
└── README.md
```


---
<!--
## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-bookstore-demo.git
cd go-bookstore-demo
```

### 2. Generate gRPC Code

Make sure you have `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` installed.

```bash
protoc --go_out=. --go-grpc_out=. proto/payment.proto
```

### 3. Run the Services

In separate terminals:

```bash
# Terminal 1: Payment Service (gRPC server)
cd payment
go run main.go
```

```bash
# Terminal 2: Bookstore Service (REST + gRPC client)
cd bookstore
go run main.go
```

---

## 🧪 Sample Requests

```bash
# Add a book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "The Go Programming Language", "price": 35.0}'

# Buy a book (calls Payment Service via gRPC)
curl -X POST http://localhost:8080/buy/{book_id}
```

---


