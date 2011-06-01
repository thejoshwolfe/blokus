
# GOROOT must be set.
ifeq ($(GOROOT),)
  $(error $$GOROOT is not set; use gomake or set $$GOROOT in your environment)
endif
include $(GOROOT)/src/Make.inc

TARG=blokus
GOFILES:=$(wildcard *.go)

include $(GOROOT)/src/Make.cmd

.PHONY: fmt
fmt:
	gofmt -w $(GOFILES)

