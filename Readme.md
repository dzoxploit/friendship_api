Sure, here's an example `README.md` for your project:

````markdown
# Friendship Management API

This is a simple friendship management API built using Golang and MySQL. It supports creating friend requests, accepting/rejecting friend requests, listing friend requests, listing friends, retrieving common friends, and blocking users.

## Features

- Create a friend request
- Accept or reject a friend request
- List friend requests
- List friends
- Retrieve common friends
- Block a user

## Requirements

- Docker
- Docker Compose

## Setup

### 1. Clone the repository

```sh
git clone https://github.com/yourusername/friendship-api.git
cd friendship-api
```
````

### 2. Create the `.env` file

Create a `.env` file in the project root with the following content:

```env
MYSQL_DSN=root:password@tcp(db:3306)/friendship?charset=utf8mb4&parseTime=True&loc=Local
```

### 3. Build and run the Docker containers

```sh
docker-compose up --build
```

This command will build the Docker images and start the containers.

## API Endpoints

### 1. Create a Friend Request

- **Endpoint**: `POST /friend-request`
- **Request Body**:
  ```json
  {
    "requestor": "andy@example.com",
    "to": "john@example.com"
  }
  ```
- **Response**:
  ```json
  {
    "success": true
  }
  ```

### 2. Accept or Reject a Friend Request

- **Endpoint**: `POST /friend-request/respond`
- **Request Body**:
  ```json
  {
    "requestor": "andy@example.com",
    "to": "john@example.com",
    "action": "accept" // or "reject"
  }
  ```
- **Response**:
  ```json
  {
    "success": true
  }
  ```

### 3. List Friend Requests

- **Endpoint**: `GET /friend-requests`
- **Request Body**:
  ```json
  {
    "email": "john@example.com"
  }
  ```
- **Response**:
  ```json
  {
    "requests": [
      { "requestor": "andy@example.com", "status": "accepted" },
      { "requestor": "joe@example.com", "status": "rejected" },
      { "requestor": "grace@example.com", "status": "pending" }
    ]
  }
  ```

### 4. List Friends

- **Endpoint**: `GET /friends`
- **Request Body**:
  ```json
  {
    "email": "andy@example.com"
  }
  ```
- **Response**:
  ```json
  {
    "friends": ["john@example.com", "joe@example.com"]
  }
  ```

### 5. Retrieve Common Friends

- **Endpoint**: `GET /common-friends`
- **Request Body**:
  ```json
  {
    "friends": ["andy@example.com", "john@example.com"]
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "friends": ["frank@example.com"],
    "count": 1
  }
  ```

### 6. Block a User

- **Endpoint**: `POST /block-user`
- **Request Body**:
  ```json
  {
    "requestor": "andy@example.com",
    "block": "john@example.com"
  }
  ```
- **Response**:
  ```json
  {
    "success": true
  }
  ```

## Testing the API

You can use `curl` or any API client like Postman to test the API endpoints. Below are some examples using `curl`:

### Create a Friend Request

```sh
curl -X POST http://localhost:8081/friend-request -H "Content-Type: application/json" -d '{
  "requestor": "andy@example.com",
  "to": "john@example.com"
}'
```

### Accept a Friend Request

```sh
curl -X POST http://localhost:8081/friend-request/respond -H "Content-Type: application/json" -d '{
  "requestor": "andy@example.com",
  "to": "john@example.com",
  "action": "accept"
}'
```

### List Friend Requests

```sh
curl -X GET http://localhost:8081/friend-requests -H "Content-Type: application/json" -d '{
  "email": "john@example.com"
}'
```

### List Friends

```sh
curl -X GET http://localhost:8081/friends -H "Content-Type: application/json" -d '{
  "email": "andy@example.com"
}'
```

### Retrieve Common Friends

```sh
curl -X GET http://localhost:8081/common-friends -H "Content-Type: application/json" -d '{
  "friends": ["andy@example.com", "john@example.com"]
}'
```

### Block a User

```sh
curl -X POST http://localhost:8081/block-user -H "Content-Type: application/json" -d '{
  "requestor": "andy@example.com",
  "block": "john@example.com"
}'
```

## Technical Choices

- **Golang**: Chosen for its performance and simplicity.
- **Gin**: A lightweight and fast web framework for Golang.
- **GORM**: An ORM library for Golang, simplifying database operations.
- **MySQL**: A reliable relational database management system.
- **Docker**: To containerize the application for consistent environments across different stages of development and deployment.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

```

### Explanation

This `README.md` provides a comprehensive guide to set up, run, and test the Friendship Management API. It includes:

- **Features**: A list of features the API provides.
- **Requirements**: Dependencies required to run the project.
- **Setup**: Instructions to clone the repository, set up environment variables, and build/run the Docker containers.
- **API Endpoints**: Detailed descriptions of each API endpoint, including request and response formats.
- **Testing the API**: Example `curl` commands to test the API endpoints.
- **Technical Choices**: A brief explanation of the technologies used in the project.
- **License**: Information about the project's license.

This should help users understand how to use and test the API effectively.
```
