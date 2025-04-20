# Gin Gorm Sample

## Description

This is a sample of Go (Gin) BE application that integrates with MySQL (via GORM), Redis, and local file handling. This app also showcases how to generate files in 2 formats, spreadsheets and PDFs. This app also has a JWT implementation which demonstrates how to validate user auth. This project will be run using a docker container.

## Highlight

1. Auth implementation with JWT (as well as middleware handling)
2. Redis integration
3. DB CRUD operations
4. File generation (spreadsheets and PDFs)
5. Project structuring

## Project Structure

The project is organized as follows:

```
├── assets/ # Contain sample files for fs implementation
├── dtos/ # Defines data transfer objects
├── handlers/ # Functions to process HTTP requests and responses
├── helpers/ # Utility functions and general purpose code used across the app
├── initializers/ # Initialize components like DB connection
├── middlewares/ # Functions to process requests before they reach handlers
├── migrate/ # Functions that will maintain DB schemas
├── models/ # Describes DB schemas
├── repositories/ # Contains logic for querying and persisting data
├── routers/ # Routes setup and handler mappings
├── tests/ # Contains unit tests
├── usecases/ # Specifies application business logic
├── .dockerignore # Excluded files and directories from Docker build context
├── .env.backup # Example environment variables file
├── .gitignore # Files and directories to ignore in Git
├── docker-compose.yml # Docker container orchestration
├── Dockerfile # Docker image build specifications
├── erd.puml # ERD visualization
├── go.mod # Go modules file
├── go.sum # Go checksum file
├── LICENSE # App license info
├── main.go # App entry point
├── makefile # Command shortcuts
└── README.md # Project documentation
```

## Prerequisites

Before running this project, you need to have the following installed:

- **Go 1.23 or higher**: For running the Go application
- **MySQL**: DB used by the app
- **Redis**: Redis data storage
- **Docker**: To run the project as this project implements containerization
- **Make (Optional)**: To simplify common tasks by creating shortcuts
- **PlantUML VSCode Extension (Optional)**: To view erd.puml file defined in this project

## How to Run the Project

1. **Clone the repository**:

```
git clone https://github.com/ertantorizkyf/gin-gorm-sample
```

2. **Navigate into the project folder**:

```
cd gin-gorm-sample
```

3. **Adjust environment variables**:

- Adjust .env.example to fit your local machine configuration. It will be copied to the docker container

4. **Run the application**:

- Using docker:

```
docker-compose up --build -d
```

- Using make:

```
make up
```

5. **The application should now be running (at port 3000 by default)**:

- Hit the `/ping` endpoint to test if the app runs properly

## Documentation

1. DB schemas can be viewed under the `erd.puml` file

## License

This project is licensed under the MIT License.
