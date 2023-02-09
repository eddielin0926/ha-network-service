# High Availability Network Service

## Perquisites

- Go
- Docker

## Start 

Ensure the `.env` is in the project root folder and variables are set.

### Checkout Configuration

```shell
docker-compose config
```

### Start Docker Compose

```shell
docker-compose up -d --build
```

### Stop Docker Compose 

```shell
docker-compose down
```

## Kubernetes

### Docker 

### Start

```shell
kubectl apply -f deploy
```

### Stop

```shell
kubectl delete -f deploy
```