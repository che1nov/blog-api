# Blogging Platform API

This project is a simple RESTful API for a personal blogging platform built with Go. It demonstrates basic CRUD operations for blog posts, including creating, reading, updating, and deleting blog posts.

## Goals

The goals of this project are to help you:
- Understand what RESTful APIs are, including best practices and conventions.
- Learn how to create a RESTful API.
- Learn about common HTTP methods like GET, POST, PUT, PATCH, DELETE.
- Learn about status codes and error handling in APIs.
- Learn how to perform CRUD operations using an API.
- Learn how to work with databases.

## Requirements

The API allows users to perform the following operations:
- Create a new blog post.
- Update an existing blog post.
- Delete an existing blog post.
- Get a single blog post.
- Get all blog posts.
- Filter blog posts by a search term.

## API Endpoints

### Create Blog Post

Create a new blog post using the POST method.

```
POST /posts
{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}
```

Each blog post should have the following fields:

```
{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}
```

The endpoint should validate the request body and return a 201 Created status code with the newly created blog post, or a 400 Bad Request status code with error messages in case of validation errors.

### Update Blog Post

Update an existing blog post using the PUT method.

```
PUT /posts/{id}
{
  "title": "My Updated Blog Post",
  "content": "This is the updated content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}
```

The endpoint should validate the request body and return a 200 OK status code with the updated blog post, or a 400 Bad Request status code with error messages in case of validation errors. It should return a 404 Not Found status code if the blog post was not found.

### Delete Blog Post

Delete an existing blog post using the DELETE method.

```
DELETE /posts/{id}
```

The endpoint should return a 204 No Content status code if the blog post was successfully deleted, or a 404 Not Found status code if the blog post was not found.

### Get Blog Post

Get a single blog post using the GET method.

```
GET /posts/{id}
```

The endpoint should return a 200 OK status code with the blog post, or a 404 Not Found status code if the blog post was not found.

### Get All Blog Posts

Get all blog posts using the GET method.

```
GET /posts
```

The endpoint should return a 200 OK status code with an array of blog posts. You don’t have to implement pagination, authentication, or authorization for this project. You can focus on the core functionality of the API.

### Filter Blog Posts

Filter blog posts by a search term using the GET method.

```
GET /posts?term={searchTerm}
```

This should return all blog posts that have the term in their title, content, or category. You can use a simple SQL query if you are using a SQL database or a similar query for a NoSQL database.

## Installation

1. Clone the repository:
```
git clone https://github.com/your-username/blogging-api.git
```

2. Change to the project directory:
```
cd blogging--api
```

3. Install the dependencies:
```
go mod tidy
```

4. Start the server:
```
go run main.go
```

## Usage

Once the server is running, you can use an API client like Postman to test the endpoints.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

https://roadmap.sh/projects/blogging-platform-api