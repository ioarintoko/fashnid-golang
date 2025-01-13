# TODO API Documentation

This document outlines the API for the Todo application, providing details on available endpoints, request parameters, and response structures.

## Base URL

```
http://localhost:8087/api/todo
```

---

## Endpoints

### 1. Get All Todos

**Description:** Fetches all todo tasks with optional query parameters.

**Endpoint:**

```
GET /api/todo
```

**Query Parameters:**

- `status` (optional): Filter tasks by status (e.g., `1` for completed, `0` for pending).

**Response:**

```json
[
  {
    "IDTask": 1,
    "Description": "Design Tampilan Mobile",
    "Status": 1,
    "CreateDate": "2018-03-12T00:00:00Z",
    "DueDate": "2018-03-12T00:00:00Z",
    "ExpireDate": "2018-03-12T00:00:00Z"
  },
  {
    "IDTask": 2,
    "Description": "Setup Backend API",
    "Status": 0,
    "CreateDate": "2018-03-13T00:00:00Z",
    "DueDate": "2018-03-15T00:00:00Z",
    "ExpireDate": "2018-03-20T00:00:00Z"
  }
]
```

---

### 2. Get Todo by ID

**Description:** Fetches a single todo task by its ID.

**Endpoint:**

```
GET /api/todo/{id}
```

**Path Parameters:**

- `id` (required): The ID of the todo task.

**Response:**

```json
{
  "IDTask": 1,
  "Description": "Design Tampilan Mobile",
  "Status": 1,
  "CreateDate": "2018-03-12T00:00:00Z",
  "DueDate": "2018-03-12T00:00:00Z",
  "ExpireDate": "2018-03-12T00:00:00Z"
}
```

---

### 3. Create a New Todo

**Description:** Creates a new todo task.

**Endpoint:**

```
POST /api/todo
```

**Request Body:**

```json
{
  "Description": "Create UI Mockups",
  "Status": 0,
  "CreateDate": "2025-01-12T00:00:00Z",
  "DueDate": "2025-01-15T00:00:00Z",
  "ExpireDate": "2025-01-20T00:00:00Z"
}
```

**Response:**

Status: `200 OK`

```json
{
  "message": "OK"
}
```

---

### 4. Update an Existing Todo

**Description:** Updates an existing todo task by its ID.

**Endpoint:**

```
PUT /api/todo/{id}
```

**Path Parameters:**

- `id` (required): The ID of the todo task to update.

**Request Body:**

```json
{
  "Description": "Update Backend Logic",
  "Status": 1
}
```

**Response:**

Status: `200 OK`

```json
{
  "message": "OK"
}
```

---

### 5. Delete a Todo

**Description:** Deletes a todo task by its ID.

**Endpoint:**

```
DELETE /api/todo/{id}
```

**Path Parameters:**

- `id` (required): The ID of the todo task to delete.

**Response:**

Status: `200 OK`

```json
{
  "message": "OK"
}
```

---

### Error Responses

**400 Bad Request:**

```json
{
  "error": "Invalid input or missing parameters."
}
```

**404 Not Found:**

```json
{
  "error": "Todo not found."
}
```

**500 Internal Server Error:**

```json
{
  "error": "An internal server error occurred."
}
```

