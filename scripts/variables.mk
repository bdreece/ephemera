MAKEFLAGS       += "-l -j $(shell nproc)"
SHELL            = /usr/bin/env bash
.SUFFIXES: .go .tmpl .sql .ts .tsx .js .jsx .css .json

prefix           = /usr/local
exec_prefix     := $(prefix)
bindir          := $(prefix)/bin

datarootdir     := $(prefix)/share
datadir         := $(datarootdir)

sysconfdir      := $(prefix)/etc

localstatedir   := $(prefix)/var
runstatedir     := $(localstatedir)/run

srcdir          := $(CURDIR)
nodemoduledir   := $(srcdir)/node_modules
pkgdir          := $(srcdir)/pkg
scriptdir       := $(srcdir)/scripts
webdir          := $(srcdir)/web

tmpdir          := $(srcdir)/tmp
tmpbindir       := $(tmpdir)/bin
tmpdistdir      := $(tmpdir)/dist

program         ?= ephemera
webapp          ?= @ephemera/web

DESTDIR         := $(prefix)
WEBDESTDIR      := $(DESTDIR)/dist

GIT_SHA          = $(shell git rev-parse HEAD)
GIT_COMMIT       = $(shell git rev-parse --short HEAD)
GIT_TAG          = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY        = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

PKG             := ./...
GENPKG          := $(PKG)
GENFLAGS        :=
TAGS            :=
TESTS           := .
TESTFLAGS       :=
LDFLAGS         := -w -s

FIND            := find
FINDFLAGS       := -type f
FIND_GO          = ${FIND} $(srcdir) $(FINDFLAGS) -name "*.go"
FIND_SQL         = ${FIND} $(srcdir)/pkg/database $(FINDFLAGS) -name "*.sql"

find_vite        = ${FIND} $(webdir)/$1 $(FINDFLAGS)
FIND_VITE_PUBLIC = $(call find_vite,public)
FIND_VITE_SRC    = $(call find_vite,src)

GO	        := go
GOFLAGS         := -v
GOPACKAGE       := github.com/bdreece/$(program)
GOMOD           := $(srcdir)/go.mod
GOSTAMP         := $(srcdir)/.go.stamp
GO_BUILD         = ${GO} build $(GOFLAGS) -ldflags "$(LDFLAGS)" -tags "$(TAGS)"
GO_GENERATE      = ${GO} generate $(GOFLAGS) $(GENFLAGS)
GO_TEST          = ${GO} test $(GOFLAGS) $(TESTFLAGS)

AIR              = ${GO} tool air
AIRFLAGS        := -c configs/air.toml

go_src          := $(shell ${FIND_GO})
go_bin          := $(tmpbindir)/$(program)

sql_src         := $(shell ${FIND_SQL})
sql_gen         := $(addsuffix .go,$(sql_src))

stringer_src	:= $(shell ${FIND_GO} -exec grep -Hl 'go tool stringer' {} \;)
stringer_gen	:= $(shell ${FIND_GO} -exec grep -H 'go tool stringer' {} \;    \
                         | $(scriptdir)/extract_stringer.sed                    \
                         | $(scriptdir)/resolve_stringer.awk)

NPM              = npm
NPMFLAGS        := --workspace $(webapp)
NPMPACKAGEJSON  := $(srcdir)/package.json
NPMPACKAGELOCK  := $(nodemoduledir)/.package-lock.json
NPM_INSTALL      = ${NPM} install $(NPMFLAGS)
NPM_RUN_SCRIPT   = ${NPM} run-script $(NPMFLAGS)

vite_cfg        := tsconfig.json        \
                   tsconfig.app.json    \
                   tsconfig.node.json    \
                   vite.config.ts

vite_src         = $(shell $(FIND_VITE_PUBLIC)) \
                   $(shell $(FIND_VITE_SRC))    \
                   $(addprefix $(webdir)/,index.html $(vite_cfg))
vite_dist        = $(tmpdistdir)/index.html

PARALLEL	:= parallel
PARALLELFLAGS   := --lb
PARALLELFLAGS   += --tag 
PARALLELFLAGS   += --halt now,fail=1
PARALLEL_MAKE    = ${PARALLEL} $(PARALLELFLAGS) $(MAKE)

ifdef VERSION
        BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

ifneq ($(BINARY_VERSION),)
        LDFLAGS += -X $(GOPACKAGE).version=${BINARY_VERSION}
endif

VERSION_METADATA = unreleased
ifneq ($(GIT_TAG),)
        VERSION_METADATA =
endif

LDFLAGS         += -X $(GOPACKAGE).metadata=${VERSION_METADATA}
LDFLAGS         += -X $(GOPACKAGE).gitTag=$(GIT_TAG)
LDFLAGS         += -X $(GOPACKAGE).gitSha=$(GIT_SHA)
LDFLAGS         += -X $(GOPACKAGE).gitCommit=$(GIT_COMMIT)
