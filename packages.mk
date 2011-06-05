
PACKAGE_OBJ_DIR = $(1)/_obj
PACKAGE_FILE = $(call PACKAGE_OBJ_DIR,$(1))/$(1).a
PACKAGE_OBJ_DIRS = $(foreach P,$(PACKAGES),$(call PACKAGE_OBJ_DIR,$(P)))
PACKAGE_FILES    = $(foreach P,$(PACKAGES),$(call PACKAGE_FILE,$(P)))
PREREQ += $(PACKAGE_FILES)

FIRST_DIR = $(firstword $(subst /, ,$(1)))
RECURSE = make --no-print-directory -C
PROPOGATE = $(foreach P,$(PACKAGES),$(RECURSE) $(P) $(1) && )true

# TODO: make this work for more than a single package
$(call PACKAGE_FILE,$(PACKAGES)):
	@$(RECURSE) $(PACKAGES)

GC += $(foreach P,$(PACKAGE_OBJ_DIRS),-I$(P))
LD += $(foreach P,$(PACKAGE_OBJ_DIRS),-L$(P))

.PHONY: fmt
fmt:
	gofmt -w $(GOFILES)
	@$(call PROPOGATE,$@)

.PHONY: clean_propogate
clean: clean_propogate
clean_propogate:
	@$(call PROPOGATE,clean)

# Don't let the targets in this file be used
# as the default make target.
.DEFAULT_GOAL:=

