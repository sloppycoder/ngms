# next generation micro service project (ngms)

POC for next generation cloud native micro service using gRPC. grpc-gateway is used to translate between REST to gRPC when needed.

The benefits of the architecture:

* standardized interface spec
* more powerful than REST
* better performance using gRPC
* cloud native, integrate seamelssly into K8S

## Interface
1. protobuf interface spec
2. generate golang skeleton
3. generate python client skeleton

## Data
1. create Mongo database
2. generate test data (Ruby)

## Golang micro service
1. create gRPC microservice
	* framework [grapi](https://github.com/izumin5210/grapi)
2. use grapi's builtin [REST/gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
4. test locally with minikube/skaffold

## Load Test
1. create load generator [locust](http://locust.io) with Golang load generator slave [Boomer](https://github.com/myzhan/boomer)
2. test with gRPC and REST, see comparison
3. (?) chao monkey
