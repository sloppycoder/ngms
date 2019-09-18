## benchmark the API using locust

Use [Locust.io](http://locust.io) to benchmark API.

### install
```
 conda install -c conda-forge locust
 conda install -c conda-forge googleapis-common-protocs-grpc
 conda install -c conda-forge grpcio-tools

```

### run python gRPC client
```
python -m grpc_tools.protoc -I../accounts/api/protos -I . --python_out=. --grpc_python_out=. ../accounts/api/protos/account.proto

python client.py

# should print out something like below

ccount_id: "100-1234-5577-891"
prod_code: "1122"
prod_name: "Super Saver Account"
balances {
  amount: 10100.0
}
balances {
  amount: 9000.0
  type: AVAILABLE
}

### run locust
```
locust --host http://127.0.0.1:3000

# for bigger load use master slave mode
locust --host http://127.0.0.1:3000 --master

# open a new shell for each slave
locust --slave

# open browser to http://localhost:8089 for Locust UI

```



