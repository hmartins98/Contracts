# Go
cd ~
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/hmartins98/Contracts/CustomTypes/CustomTypes.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/hmartins98/Contracts/Auth/Auth.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/hmartins98/Contracts/CMS/CMS.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/hmartins98/Contracts/API/API.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative github.com/hmartins98/Contracts/Products/Products.proto