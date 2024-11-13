# Makefile
protoc:
	cd proto && protoc --go_opt=paths=source_relative \
	--go_out=../protogen/golang \
 	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang \
	./**/*.proto
