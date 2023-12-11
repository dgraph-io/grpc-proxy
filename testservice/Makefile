
all: test_go

test_go: test.proto
	PATH="${GOPATH}/bin:${PATH}" protoc \
	  -I. \
		--go-grpc_out=. \
		--go_out=. \
		test.proto


