# chatapp/authgo

User のデータを管理する postgres とやりとりする Restful API server

## local

postgers サービスを立ち上げた後

```
go run main.go
```

## docker

### go

#### build

```
docker build -t moffyblack/user-app:1.0.7 .
```

#### run

```
docker run --name user-app --rm -d -p 8000:8000 moffyblack/user-app:1.0.7
```
