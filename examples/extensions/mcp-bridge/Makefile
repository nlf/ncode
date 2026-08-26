.PHONY: default build install uninstall

BIN := $(notdir $(CURDIR))
# Must match "name" in extension.json (slash command/logical extension name).
EXT := mcp
NCODE_HOME ?= $(shell \
	if [ -n "$$XDG_STATE_HOME" ]; then \
		printf '%s/ncode' "$$XDG_STATE_HOME"; \
	elif [ "$$(uname -s)" = "Darwin" ]; then \
		printf '%s/Library/Application Support/ncode' "$$HOME"; \
	else \
		printf '%s/.local/state/ncode' "$$HOME"; \
	fi)

default: install

build:
	go build -o $(BIN) .

install:
	-ncode ext remove $(EXT) --yes
	rm -rf "$(NCODE_HOME)/extensions/$(BIN)"
	mkdir -p "$(NCODE_HOME)/extensions/$(EXT)"
	cp extension.json "$(NCODE_HOME)/extensions/$(EXT)/extension.json"
	go build -o "$(NCODE_HOME)/extensions/$(EXT)/$(BIN)" .

uninstall:
	-ncode ext remove $(EXT) --yes
