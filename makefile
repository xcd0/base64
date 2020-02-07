MAKEFILE_DIR := $(shell pwd)
DST=$(MAKEFILE_DIR)/release
GOARCH=amd64

all: release

.PHONY: clean
clean:
	rm -rf $(DST)

.PHONY: release
release: clean
	mkdir -p $(DST)
	#
	# build
	#
	make release-sub BIN=base64
	make release-sub BIN=base64-decode-file
	make release-sub BIN=base64-decode-pipe
	make release-sub BIN=base64-encode-file
	make release-sub BIN=base64-encode-pipe
	#
	# archive
	#
	make release-archive BIN=base64 OS=windows
	make release-archive BIN=base64 OS=macOS
	make release-archive BIN=base64 OS=linux
	make release-archive BIN=base64 OS=freeBSD
	#
	# finish
	#

release-sub:
	#
	# $(BIN)
	#
	make release-build GOOS=windows OS=windows EXT=.exe FLAGS='-ldflags="-H windowsgui"'
	make release-build GOOS=darwin  OS=macOS
	make release-build GOOS=linux   OS=linux
	make release-build GOOS=freebsd OS=freeBSD

release-build:
	mkdir -p $(DST)/$(OS)
	cd $(BIN) && GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(DST)/$(OS)/$(BIN)$(EXT) $(FLAGS)

release-archive:
	cd $(DST)/$(OS) && zip ../$(BIN)_binary_$(OS)_$(GOARCH).zip *
	rm -rf $(DST)/$(OS)
