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
3. [miniconda](https://docs.conda.io/en/latest/miniconda.html) for generating test data. Any [Python](https://www.python.org/downloads/) 3.7 installation will do
4. [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/), your local Kubernetes cluster for testing
5. [skaffold](https://skaffold.dev/), think of ```npm serve``` for K8S development

Public docker images are created using [GitHuab actions](https://github.com/features/actions) workflow.

## Create dev database and populate with seed data
Install minikube, skaffold and MongoDB server and tools first. Then

```
# populate test data 
mongoimport --db=dev --collection=accounts < benchmark/seed.json

# start account svc
cd accounts
skaffold run

# should see some output like 
Deploy complete in 248.007266ms
You can also run [skaffold run --tail] to get the logs

cd ../benchmark
skaffold run
# similar output to above

minikube service locust --url
# should display something like
http://192.168.39.32:3233

# open browser and point to the URL above, start new test, enter 1 for number of users and 1 as hatch rate, click on Start Swarming


```

## Run this with Docker without Kubernetes
The docker public images below should work for this purpose.

* [ngms-account-svc](https://cloud.docker.com/u/sloppycoder/repository/docker/sloppycoder/ngms-account-svc) 
* [ngms-locust](https://cloud.docker.com/u/sloppycoder/repository/docker/sloppycoder/ngms-locust) 

To be written. contributors welcome

```
