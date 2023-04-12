.PHONY: deadlock go-channel go-routine mutex race-condition waitgroup worker-pattern  

deadlock:
	go run -race ./cmd/deadlock

go-channel:
	go run -race ./cmd/go-channel

go-routine:
	go run -race ./cmd/go-routine

mutex:
	go run -race ./cmd/mutex

race-condition:
	go run -race ./cmd/race-condition

waitgroup:
	go run -race ./cmd/waitgroup

worker-pattern:
	go run -race ./cmd/worker-pattern

$(GOBIN)/golangci-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

lint: | $(GOBIN)/golangci-lint
	@echo Linting...
	@golangci-lint  -v --concurrency=3 --config=.golangci.yml --issues-exit-code=1 run \
	--out-format=colored-line-number 
