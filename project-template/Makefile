################################################################################
# Directorys
################################################################################
BUILDDIR:= ./docs
BOILERPLATE:=./boilerplate
ASSETS:=./assets
BMPS:=./bmps

################################################################################
# File-Lists
################################################################################
BOILERFILES:=$(BUILDDIR)/  $(BUILDDIR)/index.html  $(BUILDDIR)/.htaccess 

# Find all JSFILES IN BOILERPLATE folder
JSFILESSRC:=$(foreach dir,$(BOILERPLATE), $(wildcard $(dir)/*.js))
JSFILES:=$(subst $(BOILERPLATE),$(BUILDDIR),$(JSFILESSRC))

PNGFILESRC:=$(wildcard $(ASSETS)/*.png)
PNGTARGETS:=$(patsubst $(ASSETS)/%.png, $(BMPS)/bmp.%.go, $(PNGFILESRC))

################################################################################
# Internal
################################################################################
ASSETPREFIX := $(subst ./,,$(ASSETS))
BMPSPREFIX  := $(subst ./,,$(BMPS))


################################################################################
# Recipes
################################################################################
$(BUILDDIR)/main.wasm: main.go $(BOILERFILES) $(JSFILES) $(PNGTARGETS)
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o $@ .

$(BMPS)/bmp.%.go: $(ASSETS)/%.png
	$(eval RESSOURCENAME:= $(shell echo $(patsubst $(ASSETPREFIX)/%.png,%,$^) | sed -e 's/[^a-zA-Z0-9_]//g'))
	go run ./.tools/png2gowasbmp.go "$^" "$@" $(BMPSPREFIX) $(RESSOURCENAME)

$(BUILDDIR)/.htaccess: $(BOILERPLATE)/.htaccess
	cp $^ $@ 

$(BUILDDIR)/%.html: $(BOILERPLATE)/%.html
	cp $^ $@ 

$(BUILDDIR)/%.js: $(BOILERPLATE)/%.js
	cp $^ $@ 

setup: go.sum
	@echo "setup done"

go.sum: go.mod
	GOPRIVATE="github.com/rocco-gossmann" go mod tidy
	
start: run

.phony: echo run stop open dev clean remake $(BUILDDIR)/

echo:
	@echo "Boilerplate files: "$(BOILERFILES)
	@echo
	@echo "PNG Sources: "$(PNGFILESRC)
	@echo "PNG Targets: "$(PNGTARGETS)
	@echo
	@echo "JS Sources: "$(JSFILESSRC)
	@echo "JS Targets: "$(JSFILES)

open: 
	open http://localhost:7353

run:
	./.tools/startserver.zsh

stop:
	./.tools/stopserver.zsh

dev:
	./.tools/entr.zsh

remake: 
	rm -f $(BUILDDIR)/main.wasm
	make $(BUILDDIR)/main.wasm

$(BUILDDIR)/:
	mkdir $(BUILDDIR)

clean:
	rm -rf $(BUILDDIR) 

