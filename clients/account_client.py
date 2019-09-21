import sys
import grpc
import account_pb2
import account_pb2_grpc

class TestGrpcClient:
    def __init__(self):
        self.channel = grpc.insecure_channel('[::]:3001')
        self.stub = account_pb2_grpc.AccountServiceStub(self.channel)


    def run(self, id):
        return self.stub.GetAccount(account_pb2.GetAccountRequest(account_id=id))


if __name__ == '__main__':
    id = sys.argv[1] if len(sys.argv) > 1 else '111'
    print('retrieving account info for ' + id)
    print('result ====>')
    client = TestGrpcClient()
    resp = client.run(id)
    print(resp if resp != None else 'None')
