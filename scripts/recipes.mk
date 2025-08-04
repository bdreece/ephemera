$(go_bin) : go_src += $(sql_gen) $(stringer_gen)
$(go_bin) : $(go_src) | $(tmpbindir)
	${GO_BUILD} -o $| $(PKG)

$(sql_gen)                  : $(sql_src)
$(stringer_gen)             : $(stringer_src)
$(sql_gen) $(stringer_gen) &: $(sql_src) $(stringer_src) | $(GOSTAMP)
	${GO_GENERATE} $(GENPKG)

$(GOSTAMP) : $(GOMOD)
	${GO} mod download
	${GO} mod verify
	@touch -m $@

$(vite_dist) : $(vite_src) | $(NPMPACKAGELOCK) $(tmpdistdir)
	${NPM_RUN_SCRIPT} build

$(NPMPACKAGELOCK) : NPMFLAGS += --include-workspace-root
$(NPMPACKAGELOCK) : $(NPMPACKAGEJSON)
	${NPM_INSTALL}

$(tmpbindir) $(tmpdistdir) :
	@mkdir -p $@

