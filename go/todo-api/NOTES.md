
```bash
go run todo.go

curl localhost:8080

curl localhost:8080/hello?name=Jason
```

- There is a web framework for golang called [Gin](https://gin-gonic.com/)
- But since I'm learning Golang idiom, starting from scratch is better

- All we need is the `fmt` package for displaying the server output and the
 `net/http` for the server implementations
 - So far it's similar to express.js/node.js in terms of function and parameter

```bash
mux = http.NewServeMux()
# create a server instance

mux.HandleFunc("GET /", homeHandler)
# from that server instance attach route and controller
```

Can read up more on the package methods docs [net/http](https://pkg.go.dev/net/http)


- Instead of using ORM, industry standard for Go are learning towards [SQLC](https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html)

```bash
go mod init todo-api
# create go.mod in your . The go.mod file tracks your project's properties and 3rd party package just like package.json

go get modernc.org/sqlite
# install package and add it to dependency, like npm install
```

Next: 
1. Configure sqlc.yaml
2. Write Schema and Queries
3. Compile SQL into Go code