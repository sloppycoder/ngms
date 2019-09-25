## gRPC client in Python

### install tools
```
 pip install grpcio grpcio-tools googleapis-common-protos

```

### run python gRPC client
```
python -m grpc_tools.protoc -I../accounts/api/protos -I . --python_out=. --grpc_python_out=. ../accounts/api/protos/account.proto

python account_client.py <optional_account_id> 

# should print out something like below

retrieving account info for 111
result ====>
id: "5d8b11a9495b11be00f1a225"
account_id: "111"
nickname: "Heath"
prod_code: "1001"
prod_name: "Super Saver Account"
currency: "THB"
servicer: "6631"
status_last_updated {
  seconds: 1547249339
}
balances {
  amount: 4777342.7365
  last_updated {
    seconds: 1547249339
  }
}
balances {
  amount: 4777342.7365
  last_updated {
    seconds: 1547249339
  }
}

```



