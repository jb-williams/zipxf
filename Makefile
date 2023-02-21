BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

EXEBUILD=main.go
EXEFILE=main
EXENAME=$(notdir $(CURDIR))
MANNAME=$(EXENAME).1.gz


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
	export GOPATH=$(CURDIR)
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg; fi

build:
	export GOPATH=$(CURDIR)
	@echo "start building executables..."
	GOARCH=amd64 GOOS=linux go build -o ./bin/$(EXENAME)-linux
	GOARCH=amd64 GOOS=darwin go build -o ./bin/$(EXENAME)-darwin
	GOARCH=amd64 GOOS=windows go build -o ./bin/$(EXENAME)-windows
	@echo "completed..."

clean:
	@echo "cleaning up..."
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)-linux
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)-darwin
	@rm -rf $(BUILDPATH)/bin/$(EXENAME).exe
	@rm -rf $(BUILDPATH)/pkg

install:
	@echo "installing zipxf..either /usr/local/go/bin or /usr/local/bin...."
	@if [ ! -d /usr/local/go/bin ] ; then sudo cp bin/zipxf-linux /usr/local/bin/zipxf ; else sudo cp bin/zipxf-linux /usr/local/bin/zipxf ; fi

man:
	@if [ ! -f /usr/share/man/man1/$(MANNAME) ] ; then sudo cp $(BUILDPATH)/$(MANNAME) /usr/share/man/man1/$(MANNAME)

