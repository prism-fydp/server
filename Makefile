BUILDDIR 	= $(CURDIR)/build
BUILDNAME 	= main

# Go options
TAGS       :=
LDFLAGS    := -w -s
GOFLAGS    :=

.PHONY: build
build:
	GO111MODULE=on go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BUILDDIR)'/$(BUILDNAME) ./cmd/api

.PHONY: run
run: build
	$(BUILDDIR)/$(BUILDNAME)