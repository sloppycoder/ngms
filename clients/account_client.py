import grpc
import account_pb2
import account_pb2_grpc

class T:
    def __init__(self):
        self.channel = grpc.insecure_channel('[::]:3001')
        self.stub = account_pb2_grpc.AccountServiceStub(self.channel)


    def run(self):
        response = self.stub.GetAccount(account_pb2.GetAccountRequest(account_id="100-1234-5577-891"))
        if response != None :
            print(response)	
        else :
            print("None")

if __name__ == '__main__':
    t = T()
    t.run()
