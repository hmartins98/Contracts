# Go
cd ~\go\src
protoc --go_out=plugins=grpc:. github.com/hmartins98/Contracts/Auth/Auth.proto
protoc --go_out=plugins=grpc:. github.com/hmartins98/Contracts/CustomTypes/CustomTypes.proto
protoc --go_out=plugins=grpc:. github.com/hmartins98/Contracts/CMS/CMS.proto