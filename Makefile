.PHONY: gen run

gen:
	go run github.com/99designs/gqlgen generate

run: gen
	go run ./...

build: gen
	rm -rf bin
	CGO_ENABLED=0 \
	go build -v \
		-ldflags="-s -w" \
		-o bin/ \
		./...


.PHONY: build-docker

build-docker:
	go mod vendor
	docker build \
		-t $(IMAGE_TAG) \
		.

.PHONY: up deploy

up:
	docker compose up --detach --build