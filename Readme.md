## commands to generate the .pb files from .proto

Firstly we need to have the go extensions location (protoc-gen-go) in the Path variable to be able to generate the grpc code from .proto
```
export GO_PATH=~/go
export PATH=$PATH:$GO_PATH/bin
or 
export PATH="$PATH:$(go env GOPATH)/bin"

vi ~/.zsh_profile //  add above two exports and save

source ~/.zsh_profile // to reload the Path

```

Now the command to generate the pb files,
```
protoc -I=${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto

protoc -I=grpc/proto --go_opt=module=grpcLearning --go_out=. --go-grpc_opt=module=grpcLearning --go-grpc_out=. grpc/proto/*.proto // it uses the go_package mentioned in proto file to generate output

protoc -I=calculator/proto --go_opt=module=grpcLearning --go_out=. --go-grpc_opt=module=grpcLearning --go-grpc_out=. calculator/proto/*.proto
```

To generate the necessary certificates for SSL, use the ssl.sh script in ssl folder. It has the necessary commands to generate the csr, private key and signed cert
```
chmod +x ssl.sh
./ssl.sh
```

Tools or options we have similar to postman for grpc are 
1. evans-cli https://github.com/ktr0731/evans
2. bloom-rpc

Note: Most of this code is inspired from https://github.com/Clement-Jean/grpc-go-course 