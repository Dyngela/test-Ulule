## Introduction

This repo is made for a test at Ulule. It consists as a REST API with eight endpoints.  
Since the dataset given was quit large it is not present on this repo. Although you should be able to reproduce it, tables are described in src/model folder.


The overall architecture of this project follow rules of architecture of MVC.  
Request come in controller which are associated to a service who treat the request need according to its parameter, calling a repository function who do the sql call needed.

## Installation

The app is available directly with docker:
```shell
git clone 
```

Or you can download source code directly:
```shell
git clone https://github.com/Dyngela/test-Ulule.git
```
The go version used is 1.18.  
Before launching the app you need to connect to a database of your own, we use postgres 14 for our development. For that go in script/.env and modify
environment variable as needed.

## Test and documentation

The project is documented with comments in code for back-end developer who wish to continue.  
For front-end developer a swagger page is available. Just launch the app with for example port 8080 and go on this link:

```shell
http://localhost:8080/swagger/index.html
```

You can also import a postman collection if you want to test query already set with argument+: 

```shell
https://www.getpostman.com/collections/450f206508a93fc85185
```

## Libraries used

#### Gin: 
Gin is a library used for routing
```shell
https://github.com/gin-gonic/gin
```
#### PGX: 
PGX is a driver for PostgreSQL, allowing us to connect to a given database.
```shell
Pgx: https://github.com/jackc/pgx
```

#### Scany
Scany is a sql tool designed to bring less boiler code than normal standard sql query
```shell
Scanny: https://github.com/georgysavva/scany
```

#### Swagger
Swagger is a vastly used library for documenting API. Here we took the version for go
```shell
swagger: https://github.com/swaggo/swag
swagger-ui: https://github.com/swaggo/http-swagger
swagger-template: https://github.com/alecthomas/template
```

## License

This is Free Software, released under the MIT License.