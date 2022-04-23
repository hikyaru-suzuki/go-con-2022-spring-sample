MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
GO_VERSION := $(shell cat ${MAKEFILE_DIR}.tool-versions | grep golang | cut -d " " -f 2)

define docker-compose
	GO_VERSION=${GO_VERSION} \
	docker-compose \
		--file ${MAKEFILE_DIR}docker/docker-compose.yaml \
		--project-name go-con-2022-spring-sample \
		$1
endef

## docker-composeの操作 (usage: `make dc arg=ps` -> docker-compose ps）
.PHONY: dc
dc:
	$(call docker-compose, $(arg))

## protocを実行 (usage: `make protoc arg='protoc --proto_path=./proto --go_out=module=github.com/hikyaru-suzuki/go-con-2022-spring-sample:. ./proto/server/options/master/options.proto'`)
.PHONY: protoc
protoc:
	$(call docker-compose, run --rm protoc $(arg))

## 自動生成を実行
.PHONY: generate
generate:
	$(call docker-compose, run --rm --entrypoint sh generator ./script/protoc.sh)

## イメージをリビルド (usage: `make image-rebuild`, `make image-rebuild target=generator`)
.PHONY: image-rebuild
image-rebuild:
	$(call docker-compose, build $(target))
