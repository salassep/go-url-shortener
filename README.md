# URL Shortener

This is a simple URL shortener written in **Go** using **Redis**. The service provides endpoints to create and retrieve shortened URLs.  

---

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/salassep/go-url-shortener
   cd go-url-shortener
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run Redis Server**:
   Ensure Redis is running on your machine

4. **Run the Application**:
   ```bash
   go run main.go
   ```

---

## API Endpoints

### 1. **Create a Shortened URL**
   **Endpoint**: `POST /create-short-url`  
   **Description**: Creates a shortened URL for the given original URL.  
   **Request Body**:  
   ```json
   {
     "originalUrl": "string (required)",
     "userId": "string (required)"
   }
   ```
   **Response**:  
   ```json
  {
    "message":"Create short url success",
    "shortUrl":"http://localhost:8080/your-short-url"
  }
   ```
   **Example**:
   ```bash
   curl -X POST http://localhost:8080/create-short-url \
   -H "Content-Type: application/json" \
   -d '{"originalUrl": "https://example.com", "userId": "user123"}'
   ```

---

### 2. **Redirect to the Original URL**
   **Endpoint**: `GET /:shortUrl`  
   **Description**: Redirects to the original URL associated with the shortened URL.  
   **Example**:
   ```bash
   curl http://localhost:8080/abc123
   ```

---

If you'd like to understand the code in detail, please read this article:
https://www.eddywm.com/lets-build-a-url-shortener-in-go/