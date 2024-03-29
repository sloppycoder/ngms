# next generation micro service project (ngms)

![badge](https://github.com/sloppycoder/ngms/workflows/publish%20docker%20images/badge.svg)

POC for next generation cloud native micro service using [gRPC](https://grpc.io). [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/) is used to translate between REST to gRPC when needed.

The benefits of the architecture:

* standardized interface spec
* more express and powerful than REST
* better performance with binary based protocol
* cloud native, integrates seamlessly into K8S

## Components
* [account service](./accounts) that supports both gRPC and REST, written in Go. Scaffolding by [grapi](https://github.com/izumin5210/grapi)
* gRPC [client](./clients/README.md) for account service
* [test data](./testdata/README.md) generator using Python library [mimesis](https://github.com/lk-geimfari/mimesis) 
* [benchmark](./benchmark/README.md) program using [Locust](http://locust.io) and [Boomer](https://github.com/myzhan/boomer) slave written in Go


## Tools used in this repo
1. [Go 1.13](https://golang.org/dl/)
2. [MongoDB server 4.2](https://www.mongodb.com/download-center/community)
test data. Any [Python](https://www.python.org/downloads/) 3.7 installation will do
3. [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/), your local Kubernetes cluster for testing
4. [skaffold](https://skaffold.dev/), think of ```npm serve``` for K8S development
5. [GitHuab actions](https://github.com/features/actions) workflow is used to build and publish docker images on [Github Package Registry](https://github.com/features/package-registry).

## Create dev database and populate with seed data
Install [MongoDB server 4.2](https://www.mongodb.com/download-center/community) then 

```
# populate test data 
# this will import 1 json document into dev db, accounts collection
cd benchmark
mongoimport -u dev -p dev --db=dev --collection=accounts < benchmark/seed.json
```

## Quick Start
To quickly try the setup including Prometheus and Grafana but without installing Go and compile and etc,

1. [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)
2. make sure your /etc/hosts contain an entry called vmhost.localm which is used by db.yaml
3. Run ```cd k8s; ./all```
4. Open browser to grafana and Locust by getting url with
```
minikube service grafana --url

minikube service locust --url

``` 

## Run this with Docker without Kubernetes
The docker public images below should work for this purpose.

* [ngms/account-svc](https://github.com/sloppycoder/ngms/packages/account-svc/26706) 
* [ngms/benchmark](https://github.com/sloppycoder/ngms/packages/benchmark/26707) 

To be written. contributors welcome

```
# these steps are not yet tesed
# populate test data.. 
mongoimport -u dev -p dev --db=dev --collection=accounts < benchmark/seed.json

docker run -n account-svc docker.pkg.github.com/sloppycoder/ngms/account-svc:latest
docker run -n locust docker.pkg.github.com/sloppycoder/ngms/benchmark:latest


```
