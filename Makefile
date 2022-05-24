VERSION=$(shell git describe --tags --always)

# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o bin/ ./...

.PHONY: docker-build
# docker-build
docker-build:
	 docker build -t gendatasrv:$(VERSION) --build-arg version=$(VERSION)  .

.PHONY: generate
# generate
generate:
	docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf generate

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
