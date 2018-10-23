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

### Installation

1. Use `go get` to install dependencies
	- `go get github.com/joho/godotenv`
	- `go get github.com/gorilla/mux`

2. Create the `.env` file, or remove the `.example` suffix from the example file
	- Set the parameters to your liking
	- I have set the example values to those I use in my local development environment (except any secrets, I will leave them blank)

### Adding Routes

1. Add route handlers in `src/api`
	- The structure I have chosen is `src/api/{root_endpoint_name}/*.go`
		- I name the root file the same as the root endpoint
		- The other (nested) endpoints, I name those accordingly
		- For example `root/root.go` is `/api/root/`, `root/nested.go` is `/api/root/{root_id}/nested`, and so on
	- Once you have added the appropriate route handlers, the route definitions go at the bottom of the file
		- For each endpoint that has nested routes, those are assigned to the `ChildRoutes` field
		- Nested route definitions are defined in the nested source files

### Starting the Server

To start the server, run `go run ./src/main.go` from the project's root directory.

---

## Things I Am Unsure Of

* The error handler / wrapper I implemented (`src/router/error_wrapper.go`) is the best I could do, without knowing more about Go and Mux. If you know a better way to do it, feel free to create a Pull Request!

## Future Goals

* I am currently researching how to efficiently manage database connections in Go. Once I figure it out, or I at least have a more comfortable grasp on it, I will be implementing some database functionality.