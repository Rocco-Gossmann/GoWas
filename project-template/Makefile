BUILDDIR:= ./docs
BOILERPLATE:=./boilerplate

BOILERFILES:=$(BUILDDIR)/  $(BUILDDIR)/index.html  $(BUILDDIR)/.htaccess 

# Find all JSFILES IN BOILERPLATE folder
JSFILESSRC:=$(foreach dir,$(BOILERPLATE), $(wildcard $(dir)/*.js))
JSFILES:=$(subst $(BOILERPLATE),$(BUILDDIR),$(JSFILESSRC))


$(BUILDDIR)/main.wasm: main.go $(BOILERFILES) $(JSFILES)
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o $@ .

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
	
.phony: run stop dev clean remake $(BUILDDIR)/


run:
	./startserver.zsh

stop:
	./stopserver.zsh

dev:
	./entr.zsh

remake: 
	rm -f $(BUILDDIR)/main.wasm
	make $(BUILDDIR)/main.wasm

$(BUILDDIR)/:
	mkdir $(BUILDDIR)

clean:
	rm -rf $(BUILDDIR) 

