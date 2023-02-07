BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

EXENAME=main@latest

export GOPATH=$(CURDIR)

makedir:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg; fi

build:
	@echo "start building..."
	$(GOINSTALL) $(EXENAME)
	@echo "completed..."

clean:
	@echo "cleaning up..."
	@rm -rf $(BUILDPATH)/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

all: makedir build


