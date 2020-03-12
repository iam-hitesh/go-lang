# GO PROGRAMMING LANGUAGE
Learn Go Programming language, A repository which aims to explain and implement in Golang.

### Tech Stack used
- Golang
- GIN framework 
- GORM
- MySQL
- JWT

### Projects
1. REST API for To - Do application
2. GIN user authentication RESTful API

# Installation & Run
```
# Download mysql driver for connecting with MySQL DB
go get github.com/go-sql-driver/mysql

# ORM
go get github.com/jinzhu/gorm

# GIN framework
go get github.com/gin-gonic/gin

# For password encryption
go get golang.org/x/crypto/bcrypt

# For JWT implementation
go get github.com/dgrijalva/jwt-go
```

# DB Setup
1. Go to Config/Database.go
2. Update DBName, User, Password and Host, Port according to your database configuration

# Routes
```
Check Routes/Routes.go file
```