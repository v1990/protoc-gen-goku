WORK_PATH=$(shell pwd)
DEBUG=0
CONF=$(WORK_PATH)/config.yaml

gen:
	rm -rf out
	mkdir -p out
	cd ../../ && go install
	protoc -I proto --goku_out=debug=$(DEBUG),conf=$(CONF),workPath=$(WORK_PATH):. proto/*.proto
	# 清除个人信息
	sed -i "" "s:$(WORK_PATH):\$$WORK_PATH:g;s:$$HOME:\$$HOME:g" out/*

