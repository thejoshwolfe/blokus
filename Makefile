
# GOROOT must be set.
ifeq ($(GOROOT),)
  $(error $$GOROOT is not set; use gomake or set $$GOROOT in your environment)
endif
include $(GOROOT)/src/Make.inc

TARG=blokus.bin
GOFILES:=$(wildcard *.go)

PACKAGES = blokus
include packages.mk

include $(GOROOT)/src/Make.cmd

