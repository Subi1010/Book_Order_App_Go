# JWT Authentication Implementation

This document describes the JWT authentication system implemented in the Book Order App.

## Features

- User registration with username, password, and role
- User login with JWT token generation (includes role in token)
- Password hashing using bcrypt
- JWT token-based authentication and authorization
- Role-based access control (admin/user)
- Protected routes using authentication middleware
- User profile endpoint

## API Endpoints

### 1. Register User
**POST** `/api/v1/users/register`

Request body:
```json
{
  "username": "johndoe",
  "password": "password123",
  "role": "user"
}
```

Response (201 Created):
```json
{
  "user": {
    "id": 1,
    "username": "johndoe",
    "role": "user",
    "created_at": "2026-01-01T18:00:00Z",
    "updated_at": "2026-01-01T18:00:00Z"
  }
}
```

**Note:** Registration only creates the user. No token is generated at this stage.

### 2. Login User
**POST** `/api/v1/users/login`

Request body:
```json
{
  "username": "johndoe",
  "password": "password123"
}
```

Response (200 OK):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "johndoe",
    "created_at": "2026-01-01T18:00:00Z",
    "updated_at": "2026-01-01T18:00:00Z"
  }
}
```

### 3. Get User Profile (Protected)
**GET** `/api/v1/users/profile`

Headers:
```
Authorization: Bearer <your-jwt-token>
```

Response (200 OK):
```json
{
  "id": 1,
  "username": "johndoe",
  "created_at": "2026-01-01T18:00:00Z",
  "updated_at": "2026-01-01T18:00:00Z"
}
```

## Testing with cURL

### Register a new user:
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"johndoe","password":"password123","role":"user"}'
```

### Register an admin user:
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123","role":"admin"}'
```

### Login:
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"johndoe","password":"password123"}'
```

### Get profile (replace TOKEN with actual token):
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer TOKEN"
```

## Implementation Details

### Components Created:

1. **models/user.go** - User model with password hashing
   - User struct with username and password fields
   - Password hashing using bcrypt
   - Login and Register request/response models

2. **middleware/auth.go** - JWT authentication middleware
   - Token generation function
   - Authentication middleware for protected routes
   - Token validation and claims extraction

3. **services/user_service.go** - User business logic
   - User registration with duplicate check
   - User login with password verification
   - User retrieval by username or ID

4. **controllers/user_controller.go** - HTTP handlers
   - RegisterUser - handles user registration
   - LoginUser - handles user authentication
   - GetProfile - returns authenticated user's profile

5. **routers/user_routes.go** - Route definitions
   - Public routes: /login, /register
   - Protected routes: /profile (requires JWT token)

6. **migrations/000003_create_users_table.up.sql** - Database migration
   - Creates users table with proper indexes

### Security Notes:

- Passwords are hashed using bcrypt before storage
- JWT tokens expire after 24 hours
- The JWT secret key should be moved to environment variables in production
- Password field is excluded from JSON responses using `json:"-"` tag

### Database Schema:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user'
);
```

## Role-Based Authorization

The system supports two roles:
- **user**: Regular user with standard permissions
- **admin**: Administrator with elevated permissions

### Using Authentication in Other Routes

To protect routes with authentication only:

```go
// In your route file
import "book_order_app/middleware"

// Protected route - requires valid JWT token
books.GET("/", middleware.AuthMiddleware(), bookController.GetAllBooks)
```

To protect routes with role-based authorization:

```go
// Admin-only route
books.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), bookController.DeleteBook)

// Multiple roles allowed
books.POST("/", middleware.AuthMiddleware(), middleware.RequireRole("admin", "user"), bookController.CreateBook)
```

### Accessing User Information in Controllers

To get the authenticated user information in your controller:

```go
// Get user ID
userID, exists := c.Get("user_id")
if !exists {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
    return
}
// Use userID.(uint) to get the actual ID

// Get username
username, _ := c.Get("username")
usernameStr := username.(string)

// Get role
role, _ := c.Get("role")
roleStr := role.(string)
```

### JWT Token Structure

The JWT token includes the following claims:
- `user_id`: User's unique identifier
- `username`: User's username
- `role`: User's role (admin/user)
- `exp`: Token expiration time (24 hours from issuance)
- `iat`: Token issued at time
- `nbf`: Token not valid before time
