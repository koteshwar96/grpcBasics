### we need to have a Mongo golang driver
```
go get go.mongodb.org/mongo-driver/mongo
```

command to generate the .pb files form .proto
```
protoc -I=blog/proto --go_opt=module=grpcLearning --go_out=. --go-grpc_opt=module=grpcLearning --go-grpc_out=. blog/proto/*.proto
```