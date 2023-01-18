# chat app

k8s の練習プロジェクト。

![demo](https://user-images.githubusercontent.com/58592807/209473332-d8d3465f-0fe2-4dca-a4ed-39c7e2e02461.gif)

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
docker run -p 3000:3000 --rm moffyblack/chat-js:1.0.2
```

## log

```
kubectl logs {podname}
```

## exec

```
kubectl exec --stdin --tty {podname} -- /bin/bash
```
