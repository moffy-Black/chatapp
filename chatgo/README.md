# chatapp/chatgo

websocket でメッセージを redis に pub/sub する Restful API Server

## local

redis サービスを立ち上げた後

```
go run main.go
```

## docker

### redis

```
docker run --name redis --rm -d -p 6379:6379 redis:6-alpine3.17
```

### go

```
docker build -t chat-app .
docker tag chat-app moffyblack/chat-app:1.0.4
docker push moffyblack/chat-app:1.0.4
```
