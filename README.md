# url-shortener

## 🧩 Go URL Shortener
A simple URL shortening service written in Go. It allows you to convert long URLs into short, redirectable links using an in-memory store.

## 🚀 Features
* Shorten long URLs to short, unique codes
* Redirect to the original URL using the short code
* Simple RESTful API
* In-memory storage (no database required)

## 📦 Project Structure

url-shortener/
├── main.go                  # Entry point
├── handler/                 # HTTP handlers
│   └── url.go
├── model/                   # Request and response models
│   └── url.go
├── storage/                 # In-memory storage logic
│   └── memory.go
└── utils/                   # Utility functions (e.g., short code generator)
    └── shortener.go


## 🛠️ Setup
### 1. Clone the repo

git clone https://github.com/your-username/url-shortener.git
cd url-shortener

### 2. Initialize Go modules
go mod init url-shortener

## 3. Run the server
go run main.go
Server will start on http://localhost:8080

## 🧪 Usage
### 1. Shorten a URL
POST /shorten

Request Body:
{
  "url": "https://example.com"
}
Response:
{
  "short_url": "http://localhost:8080/abc123"
}
## 2. Redirect
### GET /:code
Redirects to the original long URL.

Example:

curl -L http://localhost:8080/abc123

## 🔒 Notes
This is an in-memory version; all data is lost on restart.

Ideal for learning or as a base to build on (e.g., MongoDB support, Redis, expiration, rate limiting, etc.)

## 🛠️ Future Improvements
Add persistent storage (e.g., MongoDB or PostgreSQL)

Expiration time for short URLs

Analytics (click tracking)

Frontend UI

Authentication


