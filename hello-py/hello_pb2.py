# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bhello.proto\x12\x05hello\"\x1c\n\x06MsgReq\x12\x12\n\nhello_name\x18\x01 \x01(\t\"\x1a\n\x06MsgRes\x12\x10\n\x08response\x18\x01 \x01(\t21\n\x06\x44\x38grpc\x12\'\n\x05Hello\x12\r.hello.MsgReq\x1a\r.hello.MsgRes\"\x00\x42.Z,github.com/driver8soft/examples/d8grpc/hellob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'hello_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/driver8soft/examples/d8grpc/hello'
  _globals['_MSGREQ']._serialized_start=22
  _globals['_MSGREQ']._serialized_end=50
  _globals['_MSGRES']._serialized_start=52
  _globals['_MSGRES']._serialized_end=78
  _globals['_D8GRPC']._serialized_start=80
  _globals['_D8GRPC']._serialized_end=129
# @@protoc_insertion_point(module_scope)
