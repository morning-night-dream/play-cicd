.PHONY: gen
gen:
	@goa gen github.com/morning-night-dream/play-cicd/design

lint:
	@golangci-lint run --fix
