# goldis
A server application which imitates Redis. It uses an in memory map to store key-value pairs.

### Install Go
This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)

```bash
go run server.go
```

Allowed commands are `DUMP`, `SET`, `RENAME`, `HELP` and used in following manner : 
DUMP <key>
SET <key> <value>
RENAME <key> <key>
HELP