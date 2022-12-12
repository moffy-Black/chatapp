# chatapp/authgo/postgresql

## docker image

build

```
docker build -t moffyblack/chat-postgres:1.0.0 .
```

run

```
docker run --rm --name chat-postgres -d -p 5432:5432 moffyblack/chat-postgres:1.0.0
```

connect test

```
go test
```

## reference

https://stackoverflow.com/questions/26598738/how-to-create-user-database-in-script-for-docker-postgres

Redis.uptrace.dev/guide/go-redis-pubsub.html
