# goldis
A server application which imitates Redis. It uses an in memory map to store key-value pairs.

### Install Go
This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)

### Run server
```bash
go run main.go
```

### Commands
Allowed commands are `DUMP`, `SET`, `RENAME`, `HELP` and used like below :
```bash 
DUMP <key>
SET <key> <value>
RENAME <key> <key>
HELP
```