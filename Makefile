
OUTDIR=$(CURDIR)/out

all:
	mkdir -p $(OUTDIR)
	go build -buildmode=c-archive -o $(OUTDIR)/libsample.a
