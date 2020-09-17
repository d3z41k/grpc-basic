# GRPC-basic

Basic GRPC server-client.

## Generating client and server code

```
protoc --proto_path=proto --go_out=plugins=grpc:proto service.proto
```