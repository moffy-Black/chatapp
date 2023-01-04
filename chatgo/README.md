# chatapp/chatgo

websocket でメッセージを redis に pub/sub する Restful API Server

## docker

### redis

```
docker run --name redis --rm -d -p 6379:6379 redis:6-alpine3.17
```

### go

```
docker build -t moffyblack/chat-app:1.0.1 .
```
