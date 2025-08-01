$(BIN): $(GOSRC) | $(GOGEN) $(TMPDIR)
	$(GO) build $(GOFLAGS) -o $(TMPDIR) $(PKG)

$(DIST): $(JSSRC) | $(JSLOCK)
	$(NPM) run $(NPMFLAGS) build

$(GOGEN): | $(GOLOCK)
	$(GO) generate $(GOFLAGS) $(PKG)

$(GOLOCK):
	$(GO) mod download
	$(GO) mod verify
	@touch -m $@

$(JSLOCK): 
	$(NPM) ci

$(TMPDIR):
	@mkdir -p $@

