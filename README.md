# grpc-try

```
protoc notification.proto --go_out=plugins=grpc:.
protoc --rust_out=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_rust_plugin` notification.proto
```

#### TROUBLESOOT:

Make sure package names match between rust and go proto files or else you face this

```
rpc error: code = Unimplemented desc =
```
