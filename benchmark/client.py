import grpc

import account_pb2
import account_pb2_grpc

def run():
    with grpc.insecure_channel('[::]:3001') as channel:
    	stub = account_pb2_grpc.AccountServiceStub(channel)
    	response = stub.GetAccount(account_pb2.GetAccountRequest(account_id="100-1234-5577-891"))
    	if response != None :
    		print(response)	
    	else :
    		print("None")

if __name__ == '__main__':
    run()
