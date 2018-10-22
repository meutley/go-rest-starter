# Go REST Starter

Go REST Starter is my own implementation of a REST API boilerplate / starter project using Go and gorilla/mux.

## Dependencies

1. Go v1.11.1 ([https://golang.org](https://golang.org))
2. Gorilla Mux ([http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux))
3. GoDotEnv ([https://github.com/joho/godotenv](https://github.com/joho/godotenv))

## My Development Environment

1. Mac OS X 10.13.3
2. Visual Studio Code v1.25.1 ([https://code.visualstudio.com/](https://code.visualstudio.com/))

---

## Getting Started

### Starting the Server

To start the server, run `go run ./src/main.go` from the project's root directory.

### Adding Routes

Add route handlers in `src/api`, then add the appropriate routes in `src/router/routes.go`.