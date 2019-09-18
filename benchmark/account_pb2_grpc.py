# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import account_pb2 as account__pb2


class AccountServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetAccount = channel.unary_unary(
        '/ngms.accounts.AccountService/GetAccount',
        request_serializer=account__pb2.GetAccountRequest.SerializeToString,
        response_deserializer=account__pb2.Account.FromString,
        )


class AccountServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetAccount(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AccountServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetAccount': grpc.unary_unary_rpc_method_handler(
          servicer.GetAccount,
          request_deserializer=account__pb2.GetAccountRequest.FromString,
          response_serializer=account__pb2.Account.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'ngms.accounts.AccountService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
