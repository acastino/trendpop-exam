## Introduction

This backend service is coded in `Golang`, and uses `GOrm` for easier communications with the database and uses `Gorilla Mux` for its routing needs. The rest are native Golang features.

The `.env` file is only used during development. When deploying, set variables on the OS level, so that the program will choose those variables, instead of loading the `.env` file.

## Project setup (development)

Running the Server:

```
go run *.go
```

## Project setup (production)

Most platforms now have an automated way of serving Golang projects, prefer to use those defaults. However, don't forget to build the project, so that it runs as fast as possible, as originally intended. You can later then simply execute the binary file, without needing the `go` command, like so:

```
go build
./backend-golang
```
