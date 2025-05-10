Welcome to the Microservices Project! This repository contains a suite of microservices built with Go, leveraging gRPC and RESTful APIs for seamless communication. The services include:

- **Authentication**: Secure user login and token management.
- **User Management**: Handle user profiles and permissions.
- **Records**: Manage and store application data.
- **Diagnostics**: Monitor and log system health.
- **Scheduling**: Manage tasks and events.
- **Notifications**: Send alerts and updates.

Explore the project structure and prerequisites below to get started!

This project is a collection of microservices built with Go, designed to work together using gRPC and RESTful APIs. It includes services for authentication, user management, records, diagnostics, scheduling, and notifications.

## Project Structure

The project is organized into the following directories:

- **auth/**: Contains the authentication microservice for secure user login and token management.
- **user/**: Handles user profiles, roles, and permissions.
- **records/**: Manages application data storage and retrieval.
- **diagnostics/**: Monitors system health and logs diagnostic information.
- **scheduling/**: Manages task scheduling and event handling.
- **notifications/**: Sends alerts and updates to users.
- **shared/**: Includes shared utilities and common code used across microservices.

Each directory represents a standalone microservice or shared component, designed to work together seamlessly.

Each folder represents a microservice or shared utilities.
---
## Prerequisites

Before you begin, ensure the following tools are installed on your system:

- **[Go](https://golang.org/dl/):** Version 1.19 or later is required to build and run the microservices.
- **[Docker](https://www.docker.com/):** Used for containerizing and running the services.
- **[Docker Compose](https://docs.docker.com/compose/):** Simplifies multi-container Docker applications.
- **[Protobuf](https://developers.google.com/protocol-buffers):** Required for defining service interfaces and generating code.
- **[gRPC](https://grpc.io/):** Used for efficient communication between microservices.
- **[Postman](https://www.postman.com/):** Optional, but recommended for testing RESTful APIs.
- **[Kafka](https://kafka.apache.org/):** Optional, but recommended for event-driven communication between services.

## Running the Project

### 1. Clone the Repository

Start by cloning the repository to your local machine:

```bash
git clone <repository-url>
cd <repository-folder>
```

### 2. Build and Run Services

Use Docker Compose to build and run all the microservices:

```bash
docker-compose up --build
```

This command will build the Docker images and start all the services defined in the `docker-compose.yml` file.

### 3. Access the Services

- **Authentication Service**: Accessible at `http://localhost:<auth-port>`
- **User Management Service**: Accessible at `http://localhost:<user-port>`
- **Records Service**: Accessible at `http://localhost:<records-port>`
- **Diagnostics Service**: Accessible at `http://localhost:<diagnostics-port>`
- **Scheduling Service**: Accessible at `http://localhost:<scheduling-port>`
- **Notifications Service**: Accessible at `http://localhost:<notifications-port>`

Replace `<auth-port>`, `<user-port>`, etc., with the respective ports configured in the `docker-compose.yml` file.

### 4. Test the APIs

Use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test the RESTful APIs. For gRPC services, you can use [Evans](https://github.com/ktr0731/evans) or any gRPC client of your choice.

### 5. Monitor Logs and Diagnostics

Monitor the logs and system health using the Diagnostics service. Logs can also be viewed in the Docker containers using:

```bash
docker logs <container-name>
```

### 6. Stop the Services

To stop all running services, use:

```bash
docker-compose down
```

