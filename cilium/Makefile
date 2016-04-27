include ../Makefile.defs

TARGET=cilium

all: $(TARGET)

tests:
	go fmt ./...
	-go test ./...

runtime-tests:

run:
	./cilium

clean:
	rm -f $(TARGET)

SOURCES := $(shell find ../policy-repo ../cilium-net-daemon ../common . -name '*.go')

$(TARGET): $(SOURCES)
	go build -o $(TARGET) ./main.go

install: all
	$(INSTALL) -m 0755 $(TARGET) $(DESTDIR)$(BINDIR)