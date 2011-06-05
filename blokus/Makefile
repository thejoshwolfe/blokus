
include $(GOROOT)/src/Make.inc

TARG=blokus
GOFILES:=$(wildcard *.go)

include $(GOROOT)/src/Make.pkg

PACKAGES = ../blokus/blokus.a
$(PREREQ): $(PACKAGES)
$(PACKAGES):
	@make --no-print-directory -C $(dir $@)

.PHONY: fmt
fmt:
	gofmt -w $(GOFILES)
