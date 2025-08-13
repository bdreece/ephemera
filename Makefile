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
	@sed -f $(scriptdir)/list_targets.sed ${MAKEFILE_LIST} \
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
	@rm -rf $(tmpdir) $(sql_gen) $(stringer_gen)

## fullclean: clean build artifacts and remove dependencies
.PHONY: fullclean
fullclean: clean
	@rm -rf $(nodemoduledir) $(GOSTAMP)

## generate: generate Go source code
.PHONY: generate
generate: $(sql_gen) $(stringer_gen)
	@echo $(stringer_src)
	@echo $(stringer_gen)

## restore: restore application dependencies
.PHONY: restore restore/go restore/js
restore: restore/go restore/js

## restore/go: restore Go module cache
restore/go: $(GOSTAMP)

## restore/js: restore NPM node_modules/ directory
restore/js: $(NPMPACKAGELOCK)

## build: build the application
.PHONY: build build/go build/js
build: build/go build/js

## build/go: compile the Go backend
build/go: $(go_bin)

## build/js: transpile the SolidJS frontend.
build/js: $(vite_dist)

## watch: launch applications with live-reload
.PHONY: watch watch/go watch/js
watch:
	@${PARALLEL_MAKE} ::: "watch/go" "watch/js"

## watch/go: watch the Go backend with Air
watch/go:
	${AIR} $(AIRFLAGS)

## watch/vite: serve the SolidJS frontend.
watch/js: restore/js
	${NPM_RUN_SCRIPT} dev

## test: run all testing suites
.PHONY: test test/go test/js
test: test/go test/js

## test/go: run Go testing suite
test/go: generate
	${GO_TEST} $(PKG)

## test/js: run Vitest testing suite
test/js: restore/js
	${NPM_RUN_SCRIPT} test

include scripts/recipes.mk

