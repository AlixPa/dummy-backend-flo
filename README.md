# Small Go Backend for Florian

This repository contains the backend for a simple application that manages a list of user profiles.

## Getting Started

### Prerequisites

You'll need to have Go installed.

---

**Mac os**
```
brew install go
```

**Windows** Using the Official Installer

1. Visit the [official Go downloads page](https://go.dev/doc/install).
2. Download the .msi installer for Windows.
3. Run the installer and follow the setup steps (it sets PATH automatically).

---

Verify the installation (in bash):
```bash
go version
```

You should see output like:

![](src/images/go_version.png)
With your own architecture written on the second line.

### Environment Variables

1. Copy the example environment file as `.env`:

```bash
cp .env_example .env
```

2. Edit `.env` and set your own environment variables.

## Running the Backend

In the root directory of the project, run:

```bash
go run main.go
```

## API Overview

This backend exposes a couple of basic APIs.

### Error Responses

All endpoints may return the following error responses:

- 400 Bad Request: When the request body is invalid or missing required fields
- 404 Not Found: When the requested profile ID doesn't exist
- 500 Internal Server Error: When an unexpected server error occurs

**Error Response Example**:
```json
{
  "error": "Profile not found"
}
```

---

### Ping Endpoint

Base `URL: /api/ping`

---

URL: `GET /api/ping`

Health check point.

**Response Code**: 200 OK

**Response**:
```json
{
  "message": "pong"
}
```

---

### Profiles Endpoint

Base `URL: /api/profiles`

---

`GET /api/profiles`

Retrieves the list of all stored profiles.

**Response Code**: 200 OK

**Response Example**: 
```json
{
  "data": [
    {
      "name": "Alice",
      "age": 30
    },
    {
      "name": "Bob",
      "age": 25
    }
  ]
}
```

---

`GET /api/profiles/{id}`

Retrieves a specific profile by ID.

**Response Code**: 200 OK

**Response Example**:
```json
{
  "data": {
    "name": "Alice",
    "age": 30
  }
}
```

---

`POST /api/profiles`

Creates a new profile.

**Request Body Example**:
```json
{
  "name": "Charlie",
  "age": 22
}
```

**Response Code**: 201 Created
```

---

`PUT /api/profiles/{id}`

Updates an existing profile by ID.

**Request Body Example**:
```json
{
  "name": "Charlie Updated",
  "age": 23
}
```

**Response Code**: 204 No Content

---

`DELETE /api/profiles/{id}`

Deletes a profile by ID.

**Response Code**: 204 No Content

# Dummy Backend API

A simple REST API backend built with FastAPI.

## Setup

1. Create a virtual environment:
```bash
python -m venv venv
```

2. Activate the virtual environment:
- On macOS/Linux:
```bash
source venv/bin/activate
```
- On Windows:
```bash
.\venv\Scripts\activate
```

3. Install dependencies:
```bash
pip install -r requirements.txt
```

## Running the Application

To run the application, use the following command:

```bash
uvicorn src.main:app --reload
```

The API will be available at:
- API: http://localhost:8000
- API Documentation: http://localhost:8000/docs
- Alternative API Documentation: http://localhost:8000/redoc

## API Endpoints

- `GET /`: Welcome message
- `GET /health`: Health check endpoint

## Development

This project uses FastAPI for the backend API. The main application code is in `src/main.py`.