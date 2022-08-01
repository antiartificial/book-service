### gRPC Book Service

Cobbled together from an example project to better understand gRPC.
Fire up an ephemeral MongoDB and execute the client and server:
```console
$ docker run -d -p 27017:27017 --name example-mongo mongo:latest
$ go run cmd/server/main.go
$ go run cmd/client/main.go
```

To persist data inside the Docker filesystem. command creates a new Docker volume called mongo-data and mounts it into the container. The volume will be managed by Docker; you can see it by running docker volumes ls.
```console
docker run -d 
    -p 27017:27017 
    --name example-mongo 
    -v mongo-data:/data/db 
    mongo:latest
```

To setup a MongoDB user/pass without exposing the PW, just reference in a file:
```console
docker run -d 
    -p 27017:27017 
    --name example-mongo 
    -v mongo-data:/data/db 
    -e MONGODB_INITDB_ROOT_USERNAME=example-user 
    -e MONGODB_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongo-root-pw 
    mongo:latest
```