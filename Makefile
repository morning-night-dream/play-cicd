GOLANG_VERSION = 1.19.0
DEV_CONTAINER_NAME = golangcicd

.PHONY: goa
goa:
	@goa gen github.com/morning-night-dream/play-cicd/design

.PHONY: goa-docker
goa-docker:
	@$(call _docker_,goa gen github.com/morning-night-dream/play-cicd/design)

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: tidy-docker
tidy-docker:
	@$(call _docker_,go mod tidy)

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: lint-docker
lint-docker:
	@$(call _docker_,golangci-lint run --fix)

bash:
	@$(call _docker_,bash)

tool-docker:
	@$(call _docker_,go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0) && \
	$(call _docker_,go install goa.design/goa/v3/cmd/goa@v3) && \
	$(call _docker_,go install google.golang.org/protobuf/cmd/protoc-gen-go@latest) && \
	$(call _docker_,go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest)

dev:
	@air

define _docker_
	docker run --rm -it\
			-v `pwd`/.go:/go \
			-v `pwd`:/go/src \
			-w /go/src \
			--name $(DEV_CONTAINER_NAME) \
			golang:$(GOLANG_VERSION)-bullseye \
			$1
endef
