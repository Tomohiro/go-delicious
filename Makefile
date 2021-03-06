# Project information
PACKAGE=delicious

# Tasks
help:
	@echo "Please type: make [target]"
	@echo "  test         Run tests"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  help         Show this help messages"

test: deps
	@echo "===> Running tests..."
	go test -v ./${PACKAGE} -cover

deps:
	@echo "===> Installing runtime dependencies..."
	@go get github.com/govend/govend
	govend -v

updatedeps:
	@echo "===> Updating runtime dependencies..."
	govend -v -u

.PHONY: help test deps updatedeps
