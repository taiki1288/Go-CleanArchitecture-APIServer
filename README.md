# Go-CleanArchitecture-APIServer
### Go(Echo)×GORM×CleanArchitectureでAPIServerを作成。
## language
<p align="center">
  <a href="https://golang.org/"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Go_Logo_Aqua.svg/1280px-Go_Logo_Aqua.svg.png" width="200px";
  </a>
  <a></a>
</p>
  
## Architecture
### CleanArchitecture
<p align="center">
  <img src="https://cdn-ak.f.st-hatena.com/images/fotolife/a/a_beco/20161211/20161211205919.jpg" width="600px" height="450px";>
</p>

## Startup
```
$ git@github.com:taiki1288/Go-CleanArchitecture-APIServer.git
$ cd Go-CleanArchitecture-APIServer
$ docker-compose up -d
$ go run server.go
```

## Routing
### POST
```
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"id": 1, "name": "test", "email": "test@example.com", "age": 20}' localhost:1323/users
```
### GET
```
$ curl -i -H 'Content-Type:application/json' localhost:1323/users
```

### GET
```
$ curl -i -H 'Content-Type:application/json' localhost:1323/users/1
```

### PUT
```
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"id": 1, "name": "testtest", "email": "test2@example.com", "age": 21}' localhost:1323/users/1
```

### DELETE
```
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE localhost:1323/users/1
```
