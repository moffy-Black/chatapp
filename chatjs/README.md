# chatapp/chatjs

## docker

### js

build
mac

```
docker build -t moffyblack/chat-js:1.0.2 .
```

linux

```
docker buildx build --platform linux/amd64 -t moffyblack/chat-js:1.0.2 .
```

run

```
docker run -p 3000:3000 --rm moffyblack/chat-js:1.0.2
```

## Reference

- [Complete guide to authentication with React Router v6](https://blog.logrocket.com/complete-guide-authentication-with-react-router-v6/)
