PWD=$(shell pwd)
GOROOT=$(shell go env GOROOT)
PKG=github.com/v1990/protoc-gen-goku
godoc:
#	rm -rf $(GOROOT)/$(PKG)
#	mkdir -p $(shell dirname $(GOROOT)/$(PKG))
#	ln -s $(PWD) $(GOROOT)/$(PKG)
	godoc -http=:6060
