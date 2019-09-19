# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: account.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='account.proto',
  package='ngms.accounts',
  syntax='proto3',
  serialized_options=_b('Z/github.com/sloppycoder/ngms/accounts/api;api_pb'),
  serialized_pb=_b('\n\raccount.proto\x12\rngms.accounts\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\"y\n\x07\x41\x63\x63ount\x12\n\n\x02id\x18\x01 \x01(\t\x12\x12\n\naccount_id\x18\x02 \x01(\t\x12\x11\n\tprod_code\x18\x03 \x01(\t\x12\x11\n\tprod_name\x18\x04 \x01(\t\x12(\n\x08\x62\x61lances\x18\x05 \x03(\x0b\x32\x16.ngms.accounts.Balance\"\xaf\x01\n\x07\x42\x61lance\x12\x0e\n\x06\x61mount\x18\x01 \x01(\x01\x12)\n\x04type\x18\x02 \x01(\x0e\x32\x1b.ngms.accounts.Balance.Type\x12\x13\n\x0b\x63redit_flag\x18\x03 \x01(\x08\x12\x30\n\x0clast_updated\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\"\n\x04Type\x12\x0b\n\x07\x43URRENT\x10\x00\x12\r\n\tAVAILABLE\x10\x01\"\'\n\x11GetAccountRequest\x12\x12\n\naccount_id\x18\x01 \x01(\t2x\n\x0e\x41\x63\x63ountService\x12\x66\n\nGetAccount\x12 .ngms.accounts.GetAccountRequest\x1a\x16.ngms.accounts.Account\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/accounts/{account_id}B1Z/github.com/sloppycoder/ngms/accounts/api;api_pbb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])



_BALANCE_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='ngms.accounts.Balance.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CURRENT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AVAILABLE', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=389,
  serialized_end=423,
)
_sym_db.RegisterEnumDescriptor(_BALANCE_TYPE)


_ACCOUNT = _descriptor.Descriptor(
  name='Account',
  full_name='ngms.accounts.Account',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ngms.accounts.Account.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='ngms.accounts.Account.account_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='prod_code', full_name='ngms.accounts.Account.prod_code', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='prod_name', full_name='ngms.accounts.Account.prod_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='balances', full_name='ngms.accounts.Account.balances', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=124,
  serialized_end=245,
)


_BALANCE = _descriptor.Descriptor(
  name='Balance',
  full_name='ngms.accounts.Balance',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='amount', full_name='ngms.accounts.Balance.amount', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='ngms.accounts.Balance.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='credit_flag', full_name='ngms.accounts.Balance.credit_flag', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='last_updated', full_name='ngms.accounts.Balance.last_updated', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _BALANCE_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=248,
  serialized_end=423,
)


_GETACCOUNTREQUEST = _descriptor.Descriptor(
  name='GetAccountRequest',
  full_name='ngms.accounts.GetAccountRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='ngms.accounts.GetAccountRequest.account_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=425,
  serialized_end=464,
)

_ACCOUNT.fields_by_name['balances'].message_type = _BALANCE
_BALANCE.fields_by_name['type'].enum_type = _BALANCE_TYPE
_BALANCE.fields_by_name['last_updated'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_BALANCE_TYPE.containing_type = _BALANCE
DESCRIPTOR.message_types_by_name['Account'] = _ACCOUNT
DESCRIPTOR.message_types_by_name['Balance'] = _BALANCE
DESCRIPTOR.message_types_by_name['GetAccountRequest'] = _GETACCOUNTREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Account = _reflection.GeneratedProtocolMessageType('Account', (_message.Message,), dict(
  DESCRIPTOR = _ACCOUNT,
  __module__ = 'account_pb2'
  # @@protoc_insertion_point(class_scope:ngms.accounts.Account)
  ))
_sym_db.RegisterMessage(Account)

Balance = _reflection.GeneratedProtocolMessageType('Balance', (_message.Message,), dict(
  DESCRIPTOR = _BALANCE,
  __module__ = 'account_pb2'
  # @@protoc_insertion_point(class_scope:ngms.accounts.Balance)
  ))
_sym_db.RegisterMessage(Balance)

GetAccountRequest = _reflection.GeneratedProtocolMessageType('GetAccountRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETACCOUNTREQUEST,
  __module__ = 'account_pb2'
  # @@protoc_insertion_point(class_scope:ngms.accounts.GetAccountRequest)
  ))
_sym_db.RegisterMessage(GetAccountRequest)


DESCRIPTOR._options = None

_ACCOUNTSERVICE = _descriptor.ServiceDescriptor(
  name='AccountService',
  full_name='ngms.accounts.AccountService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=466,
  serialized_end=586,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetAccount',
    full_name='ngms.accounts.AccountService.GetAccount',
    index=0,
    containing_service=None,
    input_type=_GETACCOUNTREQUEST,
    output_type=_ACCOUNT,
    serialized_options=_b('\202\323\344\223\002\030\022\026/accounts/{account_id}'),
  ),
])
_sym_db.RegisterServiceDescriptor(_ACCOUNTSERVICE)

DESCRIPTOR.services_by_name['AccountService'] = _ACCOUNTSERVICE

# @@protoc_insertion_point(module_scope)