BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

EXEBUILD=main.go
EXENAME=$(BUILDPATH) | awk -F'/' '{print $(NF)}'

export GOPATH=$(CURDIR)

echo:
	@echo "buildpath $(BUILDPATH) exebuild $(EXEBUILD) exename $(EXENAME)"
makedir:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg; fi

build:
	@echo "start building..."
	$(GOBUILD) $(EXEBUILD)
	@cp $(BUILDPATH)/$(EXEBUILD) $(EXENAME)
	@echo "completed..."

clean:
	@echo "cleaning up..."
	@rm -rf $(BUILDPATH)/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

all: makedir build


