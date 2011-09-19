include $(GOROOT)/src/Make.inc

ifneq ($(GOARCH),386)
$(error Invalid $$GOARCH '$(GOARCH)'; must be a 32-bit package)
endif
TARG=github.com/jpoirier/ni488

CGOFILES=\
	ni488.go\


include $(GOROOT)/src/Make.pkg


