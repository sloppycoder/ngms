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
3. generate java client skeleton
4. (optional) generate swagger spec from spec

## Data
1. create Mongo database
2. generate test data (Ruby)

## Golang micro service
1. create gRPC microservice
	* framework go-kit vs echo
	* enable JWT or not
	* entitlement check before data access
2. envoy / other proxy REST to gRPC gateway
3. (??) additional REST transport in micro servie itself
4. test locally with minikube/skaffold

## Load Test
1. create load generator use jmeter/gatlin/locus/?
2. test with gRPC and REST, see comparison
3. (?) chao monkey
