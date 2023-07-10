# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Show list of make targets and their description
	@grep -E '^[/%.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

dirBaseName=$(notdir $(CURDIR))

# 服务列表
TARGET_BIN_CONTENT_SVR='content_svr'
TARGET_BIN_SECURITY_SVR='security_svr'


.DEFAULT_GOAL:= help


# https://github.com/favadi/protoc-go-inject-tag
.PHONY: proto
proto:
	cd ..; ./${dirBaseName}/makefile_helper.sh
	protoc-go-inject-tag -input="./protobuf/*/*.pb.go"

.PHONY: build
build:
	make proto
	make build_content_svr
	make build_security_svr

## content_svr
.PHONY: build_content_svr
build_content_svr:
	rm -f build/${TARGET_BIN_CONTENT_SVR}/bin/*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/${TARGET_BIN_CONTENT_SVR}/bin/${TARGET_BIN_CONTENT_SVR} ./app/${TARGET_BIN_CONTENT_SVR}/main.go
	cd ./build/${TARGET_BIN_CONTENT_SVR}/; zip -r ./build.zip ./*

## security_svr
.PHONY: build_security_svr
build_security_svr:
	rm -f build/${TARGET_BIN_SECURITY_SVR}/bin/*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/${TARGET_BIN_SECURITY_SVR}/bin/${TARGET_BIN_SECURITY_SVR} ./app/${TARGET_BIN_SECURITY_SVR}/main.go
	cd ./build/${TARGET_BIN_SECURITY_SVR}/; zip -r ./build.zip ./*

