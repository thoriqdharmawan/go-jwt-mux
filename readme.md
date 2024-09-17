# 🛡️ Golang JWT Authentication with Gorilla Mux

This project is a simple REST API built with Golang, using JWT (JSON Web Token) for authentication and Gorilla Mux as the HTTP router. 

## 🚀 Features

- 🔒 **User Authentication (Login & Register)** with JWT
- 🍪 **JWT Token stored in Cookies** for secure access
- 📦 **Dummy Products API** for demonstration purposes
- 🛡️ **JWT Middleware** to protect certain routes

## 🔧 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/thoriqdharmawan/go-jwt-mux.git
   cd go-jwt-mux
   ```

2. Install the required dependencies:
   ```bash
   go mod tidy
   ```

3. Set up your database connection in the `models` package.

4. Run the project:
   ```bash
   go run main.go
   ```

## 🛠️ Endpoints

| Method | Endpoint      | Description                 | Auth Required? |
|--------|---------------|-----------------------------|----------------|
| POST   | `/login`       | Login user and generate JWT | ❌              |
| POST   | `/register`    | Register new user           | ❌              |
| GET    | `/logout`      | Logout user                 | ❌              |
| GET    | `/api/products`| Get dummy product data      | ✅              |

## 🛡️ JWT Middleware

The `/api/products` route is protected by the JWT middleware, ensuring that only authenticated users can access it.

## 📦 Dummy Product Data

The `productscontroller` currently serves static, dummy product data for demonstration purposes. This can be replaced with real data from a database in the future.

## 🔐 Authentication Flow

1. **Register**: Create a new user by sending a `POST` request to `/register` with a `username` and `password`.
2. **Login**: After registering, login via `/login`. This will return a JWT token in a cookie.
3. **Access Protected Routes**: Use the JWT token to access protected routes such as `/api/products`.
4. **Logout**: Clear the token by hitting `/logout`.

## 🔑 Technologies Used

- **Golang** 🐹
- **Gorilla Mux** 🦍
- **JWT** 🔑
- **GORM** 🗄️

## ✨ Future Improvements

- Integrate with a real product database.
- Improve error handling.
- Add more user roles and permissions.