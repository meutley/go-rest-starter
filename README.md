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

---

## Things I Am Unsure Of

* The error handler / wrapper I implemented (`src/router/error_wrapper.go`) is the best I could do, without knowing more about Go and Mux. If you know a better way to do it, feel free to create a Pull Request!

## Future Goals

* I am currently researching how to efficiently manage database connections in Go. Once I figure it out, or I at least have a more comfortable grasp on it, I will be implementing some database functionality.