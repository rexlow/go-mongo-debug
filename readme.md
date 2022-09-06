## Readme

Ignore this repo please, created to debug some networking issues I have at work :)

## Run

```
MONGO_USER=MONGO_USER \
MONGO_PASS=MONGO_PASS \
MONGO_HOST=MONGO_HOST \
MONGO_AUTH=MONGO_AUTH \
MONGO_REPLICA=MONGO_REPLICA go run *.go
```

## Docker

```
docker build -t rexlow/go-mongo-debug -f Dockerfile --platform linux/amd64 .

docker push rexlow/go-mongo-debug:latest

docker run \
    -e MONGO_USER=${MONGO_USER} \
    -e MONGO_USER=${MONGO_PASS} \
    -e MONGO_USER=${MONGO_HOST} \
    -e MONGO_USER=${MONGO_AUTH} \
    -e MONGO_USER=${MONGO_REPLICA} \
    --name go-mongo-debug go-mongo-debug:latest
```