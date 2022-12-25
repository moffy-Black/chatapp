# chat app

k8s の練習プロジェクト。

![demo](https://user-images.githubusercontent.com/58592807/209473332-d8d3465f-0fe2-4dca-a4ed-39c7e2e02461.gif)

## Setup

### minikube

```
minikube start
kubectl apply -f k8s/chat-redis-leader.yml
kubectl apply -f k8s/chat-redis-follower.yml
kubectl apply -f k8s/chat-redis-pubsub.yml
kubectl apply -f k8s/chat-app.yml
```

Publish Service

```
minikube service chat-app-service --url
```

Cleaning up

```
kubectl delete deployment -l app=redis
kubectl delete service -l app=redis
kubectl delete deployment chat-app
kubectl delete service chat-app-service
```

### chatjs

```
cd chatjs
npm i
npm start
```

### chatgo

```
cd chatgo
docker build -t chat-app .
docker tag chat-app moffyblack/chat-app:1.0.1
docker push moffyblack/chat-app:1.0.1
```
