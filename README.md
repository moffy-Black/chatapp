# chat app

k8s の練習プロジェクト。

## Setup

### minikube

```
minikube start
kubectl apply -f k8s/go-redis-leader.yml
kubectl apply -f k8s/go-redis-follower.yml
kubectl apply -f k8s/chat-app.yml
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
docker tag chat-app moffyblack/chat-app:1.0.0
docker push moffyblack/chat-app:1.0.0
```
