# E-Commerce Platform Backend (Golang + PostgreSQL + JWT)

This is the backend service for an **E-Commerce Platform**, built with **Golang**, **Gin**, **PostgreSQL**, and **JWT authentication**.  
It provides APIs for user authentication, product management, and order handling.  

The project is still under development, with the goal of creating a scalable backend that can support user accounts, browsing products, managing carts, placing orders, and handling payments.
  

---

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Setup & Running](#setup--running)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

---

## Features
- 🔑 User registration & login
- 🔒 JWT-based authentication
- 📦 PostgreSQL for persistence
- 🐳 Dockerized setup
- 📑 RESTful APIs with Gin

---

## Tech Stack
- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Containerization:** Docker & Docker Compose

---

## Getting Started

### Prerequisites
- [Go 1.22+](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Setup & Running

### Step 1: Create Environment File
Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_USER=youruser
DB_PASS=yourpass
DB_NAME=authdb
DB_SSLMODE=disable
DB_PORT=5432

APP_PORT=8000
```
---

### Step 2: Run with Docker Compose 
```bash
docker compose up --build

### Step 3: Run Without Docker
```
go mod tidy
go run main.go
```

---

⚡ This intro makes it clear the backend is **for an e-commerce platform** and still a **work in progress**.  

Do you also want me to include a **sample `docker-compose.yml`** in the README so that new contributors can copy-paste and spin up the whole local environment instantly?
