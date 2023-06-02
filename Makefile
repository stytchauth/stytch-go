.PHONY: help
help:
	@echo "Available commands"
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort 

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: test
test:
	go test ./stytch/...

.PHONY: tests
tests: test # A useful alias
