# Go HTTP API — Homework Project

This is a simple HTTP service built with **Go (Golang)**.  
It demonstrates how to work with routes, JSON, query parameters, middleware, and HTTP status codes.

---

## 🚀 How to Run

1. Make sure you have **Go** installed:
   ```bash
   go version
# Go HTTP API — Homework Project

This is a simple HTTP service built with **Go (Golang)**.  
It demonstrates how to work with routes, JSON, query parameters, middleware, and HTTP status codes.

---

## 🚀 How to Run

1. Make sure you have **Go** installed:
   ```bash
   go version
Run the server:

bash
go run main.go
The server will start on:

arduino
http://localhost:8080

🔐 Middleware

All routes are protected by a simple middleware that:

Logs each request method and path (e.g. "GET /user")

Checks the header X-API-Key

Only allows requests with X-API-Key: secret123

Returns 401 Unauthorized for missing or invalid keys

Example unauthorized response:

{"error": "unauthorized"}

🧠 API Endpoints
1️⃣ GET /user

Description: Get user by ID (query parameter)

Example:

curl -i -H "X-API-Key: secret123" "http://localhost:8080/user?id=42"


✅ Success (200):

{"user_id": 42}


❌ Invalid ID (400):

{"error": "invalid id"}

2️⃣ POST /user

Description: Create a new user

Request body (JSON):

{"name": "Alice"}


Example:

curl -i -X POST -H "X-API-Key: secret123" -H "Content-Type: application/json" -d '{"name":"Alice"}' "http://localhost:8080/user"


✅ Success (201):

{"created": "Alice"}


❌ Invalid name (400):

{"error": "invalid name"}

🧩 Project Structure
Practice Go1/
│
├── main.go          # Main server file
└── README.md        # Project documentation

👨‍💻 Author

Batyrkhan Kadyrbay
Student project for Golang Application Development course.

✅ Definition of Done

go run ./main.go starts the server on port 8080

Both /user routes are protected by the middleware

All success and error paths return proper JSON with Content-Type: application/json

Middleware denies unauthorized requests and logs requests