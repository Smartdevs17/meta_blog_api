# Meta Blog API

Meta Blog API is a robust and scalable backend application for managing user authentication, blogs, and users. Built with Go and the Gin framework, it provides essential endpoints for user registration, authentication, blog management, and more.

## Features

- **User Authentication**
  - Register new users
  - User login with token-based authentication
  - Password reset functionality
  - Token validation

- **Blog Management**
  - Create, read, update, and delete blogs
  - Fetch all blogs or search by keywords
  - Retrieve single blogs or blogs by a specific user

- **User Management**
  - Retrieve all registered users (admin access)

## Technologies Used

- **Programming Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL (or compatible database)
- **Environment Variables**: Managed with `godotenv`

## Project Structure

```
meta_blog_api/
├── controllers/       # Contains all route handler functions
├── initializers/     # Handles environment variables, database connection, and migrations
├── middleware/       # Contains middleware functions for authentication
├── models/           # Defines database models
├── routes/           # (Optional) For route grouping, if implemented
├── main.go          # Entry point for the application
```

## Prerequisites

- Go (1.18 or later)
- PostgreSQL database

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/meta_blog_api.git
   cd meta_blog_api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables:

   Create a `.env` file in the root directory and define the following variables:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname
   JWT_SECRET=yourjwtsecret
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

## Endpoints

### Authentication

| Method | Endpoint                 | Description            |
|--------|--------------------------|------------------------|
| POST   | `/api/auth/register`     | Register a new user    |
| POST   | `/api/auth/login`        | User login             |
| GET    | `/api/auth/validate`     | Validate user token    |
| POST   | `/api/auth/resetpassword`| Reset user password    |

### Blog Management

| Method | Endpoint                 | Description                              |
|--------|--------------------------|------------------------------------------|
| POST   | `/api/blogs`             | Create a new blog (auth required)        |
| GET    | `/api/blogs`             | Retrieve all blogs                       |
| GET    | `/api/blogs/search`      | Search blogs by keywords                 |
| GET    | `/api/blogs/single/:id`  | Retrieve a specific blog by ID           |
| GET    | `/api/blogs/user/:id`    | Retrieve blogs by a specific user (auth) |
| PUT    | `/api/blogs/:id`         | Update a blog (auth required)            |
| DELETE | `/api/blogs/:id`         | Delete a blog (auth required)            |

### User Management

| Method | Endpoint     | Description          |
|--------|--------------|----------------------|
| GET    | `/api/users` | Retrieve all users   |

## Running Tests

To run tests, use:

```bash
go test ./...
```

## Deployment

1. Build the application:

   ```bash
   go build -o meta_blog_api
   ```

2. Run the binary:

   ```bash
   ./meta_blog_api
   ```

3. Optionally, deploy using Docker or any cloud platform like AWS, GCP, or Heroku.

## Contributing

1. Fork the repository.
2. Create a new feature branch:

   ```bash
   git checkout -b feature/your-feature
   ```

3. Commit your changes:

   ```bash
   git commit -m "Add your feature"
   ```

4. Push to the branch:

   ```bash
   git push origin feature/your-feature
   ```

5. Open a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

