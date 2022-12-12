# chatapp/chatjs

## docker

### js

build

```
docker buildx build --platform linux/amd64 -t moffyblack/chat-js:1.0.0 .
```

run

```
docker run -p 3000:3000 --rm moffyblack/chat-js:1.0.0
```
