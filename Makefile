.PHONY: gen
gen:
	@goa gen todo/design

lint:
	@golangci-lint run --fix
