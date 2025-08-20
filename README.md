# online-subs-api - Subscription Management Service

A simple **subscription management backend** written in Go.
This project allows you to create and manage subscriptions (e.g., Netflix, Spotify, YouTube Premium, etc.) with PostgreSQL as storage (GORM).

---

## 🚀 Features
- Create new subscriptions
- List all subscriptions
- Retrieve subscription by ID
- Delete subscriptions
- Built with **Go + net/http**
- Uses **PostgreSQL** (GORM) for persistence
- JSON-based API

---

# Project Structure

```
online-subs-api
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── handlers
│   └── subsHandler.go
├── models
│   └── subsModel.go
├── repo
│   ├── db.go
│   └── subsRepo.go
├── router
│   └── routes.go
├── services
│   └── subs.go
├── utils
│   ├── logger.go
│   └── uuid.go
├── .env
├── docker-compose.yaml
├── dockerfile
├── go.mod
├── go.sum
├── logs.txt
├── main.go
└── README.md
```

---

## ⚙️ Setup & Installation

### 1. Clone the repo
```bash
git clone git@github.com:baqdauletd/online-subs-api.git
cd online-subs-api
````

### 2. Install dependencies

```bash
go mod tidy
```

#### API Documentation using Swagger
Swagger docs are available in the docs/ folder - used `swaggo/swag` package for them
To generate/update Swagger docs:
```bash
swag init -g main.go
```

Once the app is running, Swagger UI is available at:

```bash
http://localhost:8080/swagger/index.html
```

### 3. Setup environment variables

Set up your `.env` file (or environment variables):

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=subscriptions
```

### 4. Run the server using docker

```bash
docker-compose up --build
```

Server will start at:

```
http://localhost:8080
```

---

## 🔑 API Endpoints

### Create Subscription

`POST /subs/create`

**Request Body (Examples):**

```json
{
    "service_name": "Bagamol Podcast",
    "price": 200000000,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0bbb",
    "start_date": "08-2025",
    "end_date": "09-2025"
}

{
    "service_name": "Netflix",
    "price": 4500,
    "user_id": "6a8b6fc1-71f2-4a2e-a1f7-f7de93bfbac3",
    "start_date": "10-2025",
    "end_date": "11-2025"
}

{
    "service_name": "Spotify Premium",
    "price": 1500,
    "user_id": "d7a9bdb2-47c0-48cd-9a25-83b9c5a123ab",
    "start_date": "03-2025",
    "end_date": "09-2025"
}

{
    "service_name": "YouTube Premium",
    "price": 2200,
    "user_id": "a17d35f4-19c7-4d80-91d2-3b5673d82e45",
    "start_date": "04-2025",
    "end_date": "07-2025"
}

{
    "service_name": "GitHub Copilot",
    "price": 10000,
    "user_id": "ba8c2ddc-48c9-40d3-a80f-48236e1f78ef",
    "start_date": "05-2025",
    "end_date": "08-2025"
}

{
    "service_name": "Netflix",
    "price": 4500,
    "user_id": "6a8b6fc1-71f2-4a2e-a1f7-f7de93bfbac3",
    "start_date": "12-2025",
    "end_date": "01-2026"
}
```

**Response Example:**

```json
{
    "id": "d14a028c-1234-5678-9abc-4f57dcd3d29b",
    "service_name": "Netflix",
    "price": 4500,
    "user_id": "6a8b6fc1-71f2-4a2e-a1f7-f7de93bfbac3",
    "start_date": "2025-10-01T00:00:00Z",
    "end_date": "2025-11-01T00:00:00Z"
}
```

---

### Get All Subscriptions

`GET /subs/listAll`

**Response Examples:**

```json
[
    {
        "id": "8c2f39eb-177d-4046-9071-3808a1169a7c",
        "service_name": "Bagamol Podcast",
        "price": 200000000,
        "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0bbb",
        "start_date": "2025-08-01T00:00:00Z",
        "end_date": "2025-08-01T00:00:00Z"
    },
    {
        "id": "3da289b1-899e-4822-97cb-c3725530d2d6",
        "service_name": "Netflix",
        "price": 4500,
        "user_id": "6a8b6fc1-71f2-4a2e-a1f7-f7de93bfbac3",
        "start_date": "2025-10-01T00:00:00Z",
        "end_date": "2025-11-01T00:00:00Z"
    },
    {
        "id": "4e25b5a1-645f-4a18-aeb7-49584ea87975",
        "service_name": "Spotify Premium",
        "price": 1500,
        "user_id": "d7a9bdb2-47c0-48cd-9a25-83b9c5a123ab",
        "start_date": "2025-03-01T00:00:00Z",
        "end_date": "2025-09-01T00:00:00Z"
    },
    {
        "id": "665627a9-6844-4ed4-bb16-f27c86a0fb71",
        "service_name": "YouTube Premium",
        "price": 2200,
        "user_id": "a17d35f4-19c7-4d80-91d2-3b5673d82e45",
        "start_date": "2025-04-01T00:00:00Z",
        "end_date": "2025-07-01T00:00:00Z"
    },
    {
        "id": "f04f080d-60c6-4ba0-b396-704778aaf57d",
        "service_name": "GitHub Copilot",
        "price": 10000,
        "user_id": "ba8c2ddc-48c9-40d3-a80f-48236e1f78ef",
        "start_date": "2025-05-01T00:00:00Z",
        "end_date": "2025-08-01T00:00:00Z"
    },
    {
        "id": "d84b5293-2158-495e-86ca-fdd082d4c1bc",
        "service_name": "Netflix",
        "price": 4500,
        "user_id": "6a8b6fc1-71f2-4a2e-a1f7-f7de93bfbac3",
        "start_date": "2025-12-01T00:00:00Z",
        "end_date": "2026-01-01T00:00:00Z"
    }
]
```

---

### Get Subscription by ID

`GET /subs/getById?id=xxxxxxxxx`

**Response Example:**

```json
{
    "id": "8c2f39eb-177d-4046-9071-3808a1169a7c",
    "service_name": "Bagamol Podcast",
    "price": 200000000,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0bbb",
    "start_date": "2025-08-01T00:00:00Z",
    "end_date": "2025-08-01T00:00:00Z"
}
```

---

### Update Subscription by ID

`GET /subs/update?id=xxxxxxxxx`

**Request:**
Also id in the query
For Update - provide all of the fields - not only the ones that are changing

```json
{
    "service_name": "Changed Just Adobe Creative 111111",
    "price": 25000,
    "user_id": "c12e7b11-6d83-44b1-95b0-918cb7e9a9f1",
    "start_date": "06-2025",
    "end_date": "09-2025"
}
```

---

### Delete Subscription

`DELETE /subs/delete?id=xxxxxxxx`

---

### Get Total Cost

`DELETE /subs/total-cost?start=2025-01-01&end=2025-02-01&user_id=<uuid>&service_name=Netflix`

---

## 🛠️ Tech Stack

* **Language:** Go
* **Framework:** net/http
* **Database:** PostgreSQL
* **ORM:** GORM
* **Containerization:** Docker

---

## 📜 License

MIT License. Feel free to use and modify.

---

## ✨ Author

👤 [Bakdaulet Dauletov] (https://github.com/baqdauletd)
