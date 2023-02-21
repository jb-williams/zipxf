BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

EXEBUILD=main.go
EXEFILE=main
EXENAME=$(notdir $(CURDIR))
MANNAME=$(EXENAME).1.gz

export GOPATH=$(CURDIR)

echo:
	@echo "buildpath $(BUILDPATH)"
	@echo "go $(GO)"
	@echo "gobuild $(GOBUILD)"
	@echo "goinstall $(GOINSTALL)"
	@echo "goclean $(GOCLEAN)"
	@echo "exec build file $(EXEBUILD)"
	@echo "exec file $(EXEFILE)"
	@echo "new exec name $(EXENAME)"
	@echo "man name $(MANNAME)"

makedir:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg; fi

build:
	@echo "start building executable..."
	GOARCH=amd64 GOOS=linux go build 
	GOARCH=amd64 GOOS=darwin go build
	GOARCH=amd64 GOOS=windows go build
	@echo "completed..."

clean:
	@echo "cleaning up..."
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)

super:
	@echo "super clean"
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

man:
	@if [ ! -f /usr/share/man/man1/$(MANNAME) ] ; then sudo cp $(BUILDPATH)/$(MANNAME) /usr/share/man/man1/$(MANNAME)

all: makedir build
