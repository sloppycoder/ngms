## benchmark the API using locust

Use [Locust.io](http://locust.io) to benchmark API.

### install
```
 pip install locust

```

### run locust
```
locust --host http://127.0.0.1:3000

# for bigger load use master slave mode
locust --host http://127.0.0.1:3000 --master

# open a new shell for each slave
locust --slave

# open browser to http://localhost:8089 for Locust UI

```

### run locust master with boomer load generator
```
# build slave
go build slave.go

# start locust master and Web UI
locust --master

# test gRPC API
./slave -proto grpc

# test with REST API
./slave -proto rest

```


