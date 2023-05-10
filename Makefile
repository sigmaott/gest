# SHELL=/bin/bash -e -o pipefail
# PWD = $(shell pwd)
base_path?=i18n-example/src/module
#module_name?= noName
fmt: ## Formats all code with go fmt
	@go fmt ./...

run_payment: fmt ## Run a controller from your host
	@go run ./payment/src/main.go
swag: fmt ## Run a controller from your host
	@ ./cmd/swagger init -dir /payment/src
gen_module:
	@ cookiecutter   https://github.com/gestgo/gest.git --directory template/module  name=$(name) base_path=$(base_path) --no-input --output-dir ./src/module
