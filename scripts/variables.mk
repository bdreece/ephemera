MAKEFLAGS       += "-l -j $(shell nproc)"
SHELL           = /usr/bin/env bash

TMPDIR          := $(CURDIR)/tmp
BINNAME         ?= ephemera
APPNAME         ?= @ephemera/app

WEBDIR          := $(CURDIR)/web
APPDIR          := $(WEBDIR)/app
DISTDIR         := $(TMPDIR)/dist
STATICDIR       := $(WEBDIR)/static
NPMMODULEDIR    := $(CURDIR)/node_modules

GIT_SHA         = $(shell git rev-parse HEAD)
GIT_COMMIT      = $(shell git rev-parse --short HEAD)
GIT_TAG         = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY       = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

PKG             := ./...
TAGS            :=
TESTS           := .
TESTFLAGS       :=
LDFLAGS         := -w -s

GO		= $(addsuffix /bin/go,$(shell go env GOROOT))
GOFLAGS         := -v

NPM		= $(shell which npm)
NPMFLAGS        := -w $(APPNAME)

PARALLEL	:= $(shell which parallel)
PARALLELFLAGS   := --lb --tag --halt now,fail=1

ifdef VERSION
        BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

# Only set version if building a tag or VERSION is set
ifneq ($(BINARY_VERSION),)
        LDFLAGS += -X github.com/bdreece/ephemera/internal/version.version=${BINARY_VERSION}
endif

VERSION_METADATA = unreleased
# Clear the "unreleased" string in metadata
ifneq ($(GIT_TAG),)
        VERSION_METADATA =
endif

LDFLAGS         += -X github.com/bdreece/ephemera/internal/version.metadata=${VERSION_METADATA}
LDFLAGS         += -X github.com/bdreece/ephemera/internal/version.gitTag=$(GIT_TAG)
LDFLAGS         += -X github.com/bdreece/ephemera/internal/version.gitSha=$(GIT_SHA)
LDFLAGS         += -X github.com/bdreece/ephemera/internal/version.gitCommit=$(GIT_COMMIT)

