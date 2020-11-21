.PHONY: proto

DEBUG=1
CONF=./config.yaml

PROTO_GO_PKG=github.com/v1990/protoc-gen-goku/examples/blog/proto_go
proto:
	rm -rf ./$(PROTO_GO_PKG) ./proto_go
	protoc -I proto --go_out=. --micro_out=. --goku_out=debug=$(DEBUG),conf=$(CONF):./out proto/posts/posts.proto
	mv ./$(PROTO_GO_PKG) ./proto_go
	rm -rf ./$(PROTO_GO_PKG) ./github.com
handler: install proto
	rm -rf ./out
	mkdir -p ./out
	protoc -I proto --goku_out=debug=$(DEBUG),conf=$(CONF):./out proto/posts/*.proto
	go vet ./out/...
install:
	cd ../../ && go install
