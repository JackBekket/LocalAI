# LocalAI API

This package implements a basic REST API with two endpoints:

1. `/health`: Returns a 200 OK status to indicate the server is healthy.
2. `/version`: Returns the API version.

The server listens on port 8080 and uses Gorilla Mux for routing. It also includes graceful shutdown functionality to allow the server to shut down cleanly when a SIGINT or SIGTERM signal is received.

## File structure

- main.go

## Code summary

The `main.go` file contains the main function, which initializes a logger, creates a new router, registers API endpoints, starts the server, and handles graceful shutdown.

The `healthHandler` function returns a 200 OK status to indicate the server is healthy.

The `versionHandler` function returns the API version.

## Edge cases

The application can be launched by running the `main.go` file using the `go run` command.

## Unclear places

None.

## Dead code

None.

## Relations between code entities

The `main` function initializes the logger, creates the router, registers the API endpoints, starts the server, and handles graceful shutdown. The `healthHandler` and `versionHandler` functions are responsible for handling the respective API endpoints.

