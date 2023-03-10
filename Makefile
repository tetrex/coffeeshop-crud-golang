### DEV ###

.PHONY: dev
dev:
	~/go/bin/air -c .air.toml
	

.PHONY:run
run:
	go run cmd/server/main.go


### BUILD ###

.PHONY: build
build:
	go build -o build/prod cmd/server/main.go

### docker ###

.PHONY: dc-up
dc-up:
	docker-compose up -d

.PHONY: dc-down
dc-down:
	docker-compose down