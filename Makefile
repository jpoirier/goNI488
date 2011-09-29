include $(GOROOT)/src/Make.inc


ifeq ($(GOARCH),amd64)
CGO_CFLAGS=-DAMD64
endif
TARG=github.com/jpoirier/ni488

CGOFILES:=ni488.go ni.$(GOARCH).go

include $(GOROOT)/src/Make.pkg





