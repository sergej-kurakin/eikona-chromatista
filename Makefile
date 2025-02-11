main_package_path = github.com/sergej-kurakin/eikona-chromatista
binary_name = eico

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## fmt: format .go files
.PHONY: fmt
fmt: deps
	go fmt ./...

## deps: tidy modfiles and fill vendors
.PHONY: deps
deps:
	go mod tidy -v
	go mod vendor

## build: build the application
.PHONY: build
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=./bin/${binary_name} ${main_package_path}

## run: run the  application
.PHONY: run
run: build
	./bin/${binary_name}
