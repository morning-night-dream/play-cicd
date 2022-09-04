GOLANG_VERSION = 1.19.0
DEV_CONTAINER_NAME = golangcicd

.PHONY: goa
goa:
	@goa gen github.com/morning-night-dream/play-cicd/design

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: tidy-docker
tidy-docker:
	@docker run --rm -it\
			-v `pwd`/.go:/go \
			-v `pwd`:/go/src \
			-w /go/src \
			--name $(DEV_CONTAINER_NAME) \
			golang:$(GOLANG_VERSION)-bullseye \
			go mod tidy

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: lint-docker
lint-docker:
	@docker run --rm -it\
			-v `pwd`/.go:/go \
			-v `pwd`:/go/src \
			-w /go/src \
			--name $(DEV_CONTAINER_NAME) \
			golang:$(GOLANG_VERSION)-bullseye \
			golangci-lint run --fix

dev:
	@air
