## gRPC client in Python

### install tools
```
 conda install -c conda-forge grpcio-tools
 conda install -c conda-forge googleapis-common-protocs-grpc

```

### run python gRPC client
```
python -m grpc_tools.protoc -I../accounts/api/protos -I . --python_out=. --grpc_python_out=. ../accounts/api/protos/account.proto

python account_client.py

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

```



