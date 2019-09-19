## gRPC client in Python

### install tools
```
 conda install -c conda-forge grpcio-tools
 conda install -c conda-forge googleapis-common-protocs-grpc

```

### run python gRPC client
```
python -m grpc_tools.protoc -I../accounts/api/protos -I . --python_out=. --grpc_python_out=. ../accounts/api/protos/account.proto

python account_client.py <optional_account_id> 

# should print out something like below

...
retrieving account info for 4548759848
result ====>
id: "5d836a5d02565e5a8db60c1c"
account_id: "4548759848"
prod_code: "1002"
prod_name: "Salary Plus Account"
balances {
  amount: 148612.8812
}
balances {
  amount: 148612.8812
}


```



