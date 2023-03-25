all:
	cd cmd/gim && go build && mv gim ../..

run:
	make all
	./gim

release-no-tag:
	goreleaser release --snapshot --clean

#generate:
#	find proto -name *.go | xargs rm -rf 
#	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/greeter/greeter.proto

#install-grpc:
#	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
#	brew install protobuf