gen_from_proto:
	protoc --proto_path=proto \
  --go_out=grpc --go_opt=module=github.com/IdentityStolen/RiskOS \
  --go-grpc_out=grpc --go-grpc_opt=module=github.com/IdentityStolen/RiskOS \
  proto/transaction.proto