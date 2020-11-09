.PHONY: help
help:
	@echo "Available commands"
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort 

.PHONY: lint
lint:
	golangci-lint run --fix
