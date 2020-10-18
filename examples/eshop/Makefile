WORK_PATH=$(shell pwd)
DEBUG=0
CONF=$(WORK_PATH)/config.yaml

GoPkg=github.com/v1990/protoc-gen-goku/examples/eshop/proto/out_go

gen:
	cd ../../ && go install
	rm -rf proto/out_go proto/out_tmp handlers
	mkdir -p proto/out_go proto/out_tmp handlers
	make gen_files files=proto/common/*.proto
	make gen_files files=proto/item/*.proto
	make gen_files files=proto/order/*.proto
	mv proto/out_tmp/$(GoPkg) proto
	rm -rf proto/out_tmp


files=
gen_files:
	protoc -I proto \
		--go_out=proto/out_tmp \
		--micro_out=proto/out_tmp \
		--goku_out=debug=$(DEBUG),conf=$(CONF),workPath=$(WORK_PATH):. \
		$(files)


