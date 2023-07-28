# Kpop Album List

A webserver app that allows you to add, delete, edit, and view Kpop albums in your collection.
Uses Go, Fiber web framework, and MongoDB.

![alt text](https://i.imgur.com/IsQ7HC3.png)
![alt text](https://i.imgur.com/2EqJtfp.png)

## Run locally

```shell
go run .
```

## Run on Docker

```shell
./bash_scripts/docker-build-run.sh
```

You can now test by going on http://localhost:3000

## Kill and remove container

This must be done so that when kubernetes deployment is run, there won't be port conflicts

```shell
docker kill $(docker ps -q --filter ancestor=thomas3212/go-client-server)
docker container prune
```

# Deploy Kubernetes Manifests

```shell
bash bash_scripts/k8s-deployment.sh
```

# Testing and port-forward

```shell
bash bash_scripts/k8s-test.sh
```

# Running

You can now run by going on http://localhost:45287