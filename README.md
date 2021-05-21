# geofinder

# BUILD
```
$ docker build -t geofinder .
```

## swag
Install the swaggo
```
$ go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```

swag
```
$ swag init -g cmd/geofinder/main.go -o docsapi
```

# RUN
```
$ docker run --publish 80:80 geofinder /app/geofinder
```

# Test
```
$ curl -X POST localhost:80/v1.0/finds -d '{"address":"81.2.69.142","countries":["GB"]}'
{"is_listed":true}
```

# API documents
```
http://localhost:80/swagger/index.html
```

# Kubernetes
```
$ cd k8s
$ kubectl apply -k ./ -v6
```
