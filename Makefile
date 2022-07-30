.PHONY: gen
gen:
	@goa gen github.com/morning-night-dream/play-cicd/design

tidy:
	@go mod tidy

lint:
	@golangci-lint run --fix

dev:
	@air
