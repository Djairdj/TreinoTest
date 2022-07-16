install-proto-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

proto:
	#protoc proto/*.proto --go_out=/pb --go-grpc_out=/pb
	#windows
	#protoc proto .\proto\*.proto --go_out=.\pb --go-grpc_out=.\pb
	protoc protoc --go_out=.\ --go-grpc_out=. proto\*.proto
	# ou ainda no windows por definição:
	protoc --go_out=.\pkg --go-grpc_out=.\pkg proto\*.proto


clean:
	rm -rf pkg/pb/*

run:
	go run cmd/main.go

.PHONY: proto