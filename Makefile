# =============================================
#            _
#   ___ _ __| |_  ___ _ __  ___ _ _ __ _
#  / -_) '_ \ ' \/ -_) '  \/ -_) '_/ _` |
# \___| .__/_||_\___|_|_|_\___|_| \__,_|
#    |_|
#
# MIT License
# Copyright (c) 2025 Brian Reece
# =============================================

include scripts/variables.mk

default: build

## help: print this help message
.PHONY: help
help:
	@cat scripts/banner.txt
	@sed -n 's/^## \(.*\): \(.*\)/\1|\2/p' ${MAKEFILE_LIST} \
		| column -t -s '|' -N TARGET,DESCRIPTION

## version: display version information
.PHONY: version
version:
	@echo 'ephemera version $(or $(BINARY_VERSION),<none>) $(if $(VERSION_METADATA),($(VERSION_METADATA)))'
	@echo 'git tag:          $(or $(GIT_TAG),<none>)'
	@echo 'git commit:       $(GIT_SHA)'
	@echo 'git commit sha:   $(GIT_COMMIT)'
	@echo 'git working tree: $(GIT_DIRTY)'
	@git status --porcelain

## clean: clean build artifacts
.PHONY: clean
clean:
	@rm -rf $(TMPDIR) $(GOGEN)

## fullclean: clean build artifacts and remove dependencies
.PHONY: fullclean
fullclean: clean
	@rm -rf $(NPMMODULEDIR) $(GOLOCK)

## generate: generate Go source code
.PHONY: generate
generate: $(GOGEN)

## restore: restore application dependencies
.PHONY: restore restore/go restore/npm
restore: restore/go restore/npm

## restore/go: restore Go module cache
restore/go: $(GOLOCK)

## restore/npm: restore NPM node_modules/ directory
restore/npm: $(NPMLOCK)

## build: build the application
.PHONY: build build/go build/solidjs
build: build/go build/solidjs

## build/go: compile the Go backend
build/go: restore/go $(BIN) 

## build/solidjs: transpile the SolidJS frontend.
build/solidjs: $(DIST)

## watch: launch applications with hot-reload capabilities
.PHONY: watch watch/go watch/solidjs
watch: | $(TMPDIR)
	@parallel $(PARALLELFLAGS) $(MAKE) ::: "watch/go" "watch/js"

## watch/go: watch the Go backend with Air
watch/go:
	$(GO) tool air -c configs/air.toml

## watch/vite: serve the SolidJS frontend.
watch/solidjs:
	$(NPM) run $(NPMFLAGS) dev

## test: run all testing suites
.PHONY: test test/go test/vitest
test: test/go test/vitest

## test/go: run Go testing suite
test/go:
	$(GO) test $(GOFLAGS) $(PKG)

## test/vitest: run Vitest testing suite
test/vitest:
	$(NPM) run $(NPMFLAGS) test

include scripts/recipes.mk

