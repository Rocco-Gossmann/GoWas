################################################################################
# Directorys
################################################################################
BUILDDIR:= ./docs
BOILERPLATE:=./boilerplate
ASSETS:=./assets
BMPS:=./bmps
RESSOURCE:=./ressources

################################################################################
# File-Lists
################################################################################
BOILERFILES:=$(BUILDDIR)/ $(BUILDDIR)/index.html  $(BUILDDIR)/.htaccess 

# Find all JSFILES IN BOILERPLATE folder
JSFILESSRC:=$(foreach dir,$(BOILERPLATE), $(wildcard $(dir)/*.js))
JSFILES:=$(subst $(BOILERPLATE),$(BUILDDIR),$(JSFILESSRC))

PNGFILESRC:=$(wildcard $(ASSETS)/*.png)
PNGTARGETS:=$(patsubst $(ASSETS)/%.png, $(BMPS)/bmp.%.go, $(PNGFILESRC))

RESSOURCETARGETDIR:=$(BUILDDIR)/$(RESSOURCE)
RESSOURCEFILESRC:=$(foreach dir,$(RESSOURCE), $(wildcard $(dir)/*))
RESSOURCETARGETS:=$(patsubst $(RESSOURCE)/%, $(RESSOURCETARGETDIR)/%, $(RESSOURCEFILESRC))

################################################################################
# Internal
################################################################################
ASSETPREFIX := $(subst ./,,$(ASSETS))
BMPSPREFIX  := $(subst ./,,$(BMPS))

################################################################################
# Recipes
################################################################################
$(BUILDDIR)/main.wasm: main.go $(BOILERFILES) $(JSFILES) $(PNGTARGETS) $(RESSOURCETARGETS)
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

$(RESSOURCETARGETDIR)/%: $(RESSOURCE)/%
	cp $^ $@

setup: go.sum
	@echo "setup done"

go.sum: go.mod
	GOPRIVATE="github.com/rocco-gossmann" go mod tidy
	
start: run

run: $(BUILDDIR)/main.wasm
	go run ./.tools/server/server.go

.phony: dev open echo clean remake $(BUILDDIR)/ $(RESSOURCETARGETDIR)

open: 
	open http://localhost:7353

echo: 
	@echo "Boilerplate files: "$(BOILERFILES)
	@echo
	@echo "PNG Sources: "$(PNGFILESRC)
	@echo "PNG Targets: "$(PNGTARGETS)
	@echo
	@echo "JS Sources: "$(JSFILESSRC)
	@echo "JS Targets: "$(JSFILES)
	@echo
	@echo "Ressource Sources:" $(RESSOURCEFILESRC)
	@echo "Ressource Targets:" $(RESSOURCETARGETS)
	@echo
	@echo "Watchfiles:"
	@echo $(shell find .. -type f \( -path "*/ressources/*" -o -name "*.png" -o -name "*.js" -o -name "*.html" -o -name "*.go" \) -not -iregex '.*/docs/.*')

dev:
	find .. -type f \( -path "*/ressources/*" -o -name "*.png" -o -name "*.js" -o -name "*.html" -o -name "*.go" \) -not -iregex '.*/docs/.*' | entr make remake


remake: 
	rm -f $(BUILDDIR)/main.wasm
	make $(BUILDDIR)/main.wasm

$(BUILDDIR)/:
	mkdir $(BUILDDIR)
	mkdir -p $(RESSOURCETARGETDIR)

clean:
	rm -rf $(BUILDDIR) 
	rm -rf $(BMPS)/bmp.*.go 

