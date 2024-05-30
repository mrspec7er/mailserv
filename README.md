# Project Setup Instructions

## Prerequisites

1. Install Docker and Docker Compose on your development environment.

## Setup

2. Clone the repository:

````sh
     git clone https://github.com/mrspec7er/mailserv
     ```

3. Configure environment variables as needed.

4. Ensure the following ports are free:
- 5432
- 8080
- 1025

## Building and Running the Application

5. Build the Docker containers:

```sh
     docker compose build
     ```

6. Run the server:
- For production:
````

docker-compose up -d

```
- For development:
```

docker-compose up watch

```

## Accessing the Application

The default port is 8080. You can access the endpoints on your local machine at:

- GET `http://localhost:8080/emails`
- POST `http://localhost:8080/emails`
- SMTP `localhost:1025`

```
