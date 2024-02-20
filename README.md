# Travel Project - Go Language Refactored Version
This is a web backend project implemented in Go language, which is used to display travel-related information. This project uses MySQL and MongoDB databases to store data.

## Installation and Startup
1. Clone or download the project code:
```
git clone https://github.com/hyh315/ltt-gc.git
```
2. Install dependencies:
```go
go mod tidy
```
3. Create and configure the database:
This project requires the use of MySQL and MongoDB databases. You need to install and start these two databases, and configure the database connection information in the config/config.yaml file.
4. Start the project:
```go
go run main.go
```
Or start the project using the compiled executable file:
```go
go build main.go
./main
```

## Dependency Versions
+ Go version: 1.16 or higher
+ MySQL version: 5.7 or higher
+ MongoDB version: 4.2 or higher

## Technology Stack
+ Go
+ Gin
+ Gorm
+ MySQL
+ MongoDB

## Author
+ hyh-design
+ Email: 2856400640@qq.com

If you have any questions or issues, please do not hesitate to contact me at 2856400640@qq.com.
