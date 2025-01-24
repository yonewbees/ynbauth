# JWT Authentication with PostgreSQL

YNBAUTH is a simple Go-based JWT authentication system that connects to a PostgreSQL database. It supports login with email and password, generates both access and refresh tokens, and provides endpoints for refreshing tokens.

## Features

- JWT-based authentication
- Access and refresh tokens
- PostgreSQL integration
- Secure token generation with environment variables

## Prerequisites

Before running the project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (1.18 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)
- A `.env` file with the necessary configurations checkout env.example 

---

## Setup Guide

Follow these steps to set up and run the project:

### 1. Clone the Repository

HTTPS

```bash
git clone https://github.com/yonewbees/ynbauth

```
SSH

```bash
git@github.com:yonewbees/ynbauth.git

```

### 2. Install Required dependencies

Navigate to ynbauth directory and install the required dependencies

```bash
cd ynbauth && go get

```

### 3. Database Configuration

Install and configure PostgreSQL on your machine then start PostgreSQL and create a database:

```bash
CREATE DATABASE testdb;

```
Create a users table:

```bash
CREATE TABLE users (
    id SERIAL PRIMARY KEY,                  -- Unique ID for each user
    username VARCHAR(255) UNIQUE NOT NULL, -- Unique username
    full_name VARCHAR(255) NOT NULL,       -- Full name of the user
    email VARCHAR(255) UNIQUE NOT NULL,    -- Unique email address
    password VARCHAR(255) NOT NULL,        -- Hashed password
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for record creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for updates
);

```

### 4. Configure Environment Variables

Create a .env file in the root directory with the following content:

```bash
DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
JWT_SECRET_KEY=your_secret_key_here

```

Replace username, password, localhost, 5432, and dbname with your PostgreSQL credentials and database name. Set JWT_SECRET_KEY to a secure random string.

### 5. Run the app

To run this app

```bash
go run ynbauth

```
Alternatively

```bash
go build -o ynbauth

```
then 

```bash
./ynbauth

```

### 6. Test the API Endpoints

Use curl or a tool like Postman to test the endpoints:

####  Obtain Access and Refresh Tokens

```bash
curl -X POST http://localhost:8080/api/auth-token-obtain \
-H "Content-Type: application/json" \
-d '{"email": "abc@test.com", "password": "mypassword"}'

```
Response:

```bash
{
    "access": "access_token_here",
    "refresh": "refresh_token_here"
}

```

#### Read the docs for more endpoints

Happy coding!