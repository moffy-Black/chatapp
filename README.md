# chat app

k8s の練習プロジェクト。

## Setup

### minikube

start

```
minikube start
```

apply manifest

```
kubectl apply -f k8s/chat-redis-pubsub.yml
kubectl apply -f k8s/chat-app.yml

kubectl apply -f k8s/user-postgres.yml
kubectl apply -f k8s/user-app.yml

kubectl apply -f k8s/service-ingress.yml
```

Publish Service

simple service

```
minikube service chat-js-service --url
```

ingress

```
minikube tunnel
```

Cleaning up

```
kubectl delete deployment -l app=redis
kubectl delete service -l app=redis
kubectl delete deployment chat-app
kubectl delete service chat-app-service

kubectl delete ConfigMap postgres-config
kubectl delete deployment -l app=postgres
kubectl delete service -l app=postgres
kubectl delete PersistentVolumeClaim postgres-pv-claim
kubectl delete PersistentVolume postgres-pv-volume

kubectl delete deployment user-app
kubectl delete service user-app-service
kubectl delete ingress chatapp-ingress
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

## log

```
kubectl logs {podname}
```

## exec

```
kubectl exec --stdin --tty {podname} -- /bin/bash
```
