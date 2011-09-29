include $(GOROOT)/src/Make.inc


CHEADER_VER=v2007
ifdef v2001
CGO_CFLAGS=-DV2001 
CHEADER_VER=v2001
endif
TARG=github.com/jpoirier/ni488

CGOFILES:=ni488.go $(CHEADER_VER).go

include $(GOROOT)/src/Make.pkg





