# next generation micro service project (ngms)

POC for next generation cloud native micro service using [gRPC](https://grpc.io). [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) is used to translate between REST to gRPC when needed.

The benefits of the architecture:

* standardized interface spec
* more express and powerful than REST
* better performance with binary based protocol
* cloud native, integrate seamlessly into K8S

## Components
* [account service](./accounts) that supports both gRPC and REST, written in Go. Scaffolding by [grapi](https://github.com/izumin5210/grapi)
* gRPC [client](./clients/README.md) for account service
* [test data](./testdata/README.md) generator using Python library [mimesis](https://github.com/lk-geimfari/mimesis) 
* [benchmark](./benchmark/REAMDE.md) program using [Locust](http://locust.io) and [Boomer](https://github.com/myzhan/boomer) slave written in Go
