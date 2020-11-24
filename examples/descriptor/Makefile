WORK_PATH=$(shell pwd)
DEBUG=1
CONF=$(WORK_PATH)/config.yaml
ROOT=$(realpath ../../)

OUT_PKG_NAME=descriptors

# PS: 先输出在测试目录，没问题再转移到正式目录
TMP_PATH=$(WORK_PATH)/out
# 正式目录
DESCRIPTOR_PATH=$(ROOT)/descriptors
install:
	cd $(ROOT) && go install
gen: install
	@echo "ROOT: $(ROOT)"
	@echo "OUT_PATH: $(TMP_PATH)"
	rm -rf $(TMP_PATH)/*.pb.go
	mkdir -p $(TMP_PATH) $(DESCRIPTOR_PATH)
	cp -f $(DESCRIPTOR_PATH)/descriptors.go $(TMP_PATH)/descriptors.go
	protoc -I proto --goku_out=debug=$(DEBUG),conf=$(CONF),workPath=$(WORK_PATH),outPath=$(TMP_PATH),outPkgName=$(OUT_PKG_NAME):. proto/google/protobuf/descriptor.proto
	#cd $(TMP_PATH) && go vet ./...
	cd $(ROOT) && go vet ./...
	#cat $(TMP_PATH)/FieldDescriptorProto.goku.pb.go
usage: install
	protoc -I proto --goku_out=debug=$(DEBUG),conf=usage.config.yaml,workPath=$(WORK_PATH),outPath=$(TMP_PATH),outPkgName=$(OUT_PKG_NAME):. proto/google/protobuf/descriptor.proto
gen_descriptors:
	make gen
	cd $(TMP_PATH) \
			&& go vet ./... \
			&& rm -f $(DESCRIPTOR_PATH)/*.pb.go \
			&& mkdir -p $(DESCRIPTOR_PATH) \
			&& cp -f $(TMP_PATH)/*.pb.go $(DESCRIPTOR_PATH)
	cd $(ROOT) && go vet ./...
