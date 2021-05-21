# geofinder

# BUILD
```
$ docker build -t geofinder .
```

# RUN
```
$ docker run --publish 80:80 geofinder /app/geofinder
```

# BUILD

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

# API documents
```
http://localhost:80/swagger/index.html
```

# Kubernetes
```
$ cd k8s
$ kubectl apply -k ./ -v6
```
