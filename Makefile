MOCKERY_DOCKER_IMAGE = vektra/mockery:v2.12.2
GOLANGCI_LINT_VERSION = v1.45.2

test:
	go test -race -v --cover ./...

fix-lint-docker:
	docker run --rm \
		-v $(PWD):/core \
		-w /core \
		golangci/golangci-lint:$(GOLANGCI_LINT_VERSION) \
		golangci-lint run --fix

docker-compose-up:
	docker compose \
		--project-directory ./deploy/build \
		build

	docker compose \
		--project-directory ./deploy/build \
		up

generate-go-mocks:
	rm ./internal/core/mocks/*.go || true
	docker run --rm \
		--volume $(PWD):/project \
		--workdir /project \
		$(MOCKERY_DOCKER_IMAGE) \
		--dir=./internal/core \
		--output=./internal/core/mocks \
		--all \
		--case snake

generate-http-go-mocks:
	rm ./internal/core/mocks/*.go || true
	docker run --rm \
		--volume $(PWD):/project \
		--workdir /project \
		$(MOCKERY_DOCKER_IMAGE) \
		--dir=./internal/infra/http/ \
		--output=./internal/infra/http/mocks \
		--all \
		--case snake