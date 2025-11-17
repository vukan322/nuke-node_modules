BINARY_NAME=nukenm
VERSION?=dev
INSTALL_PATH=/usr/local/bin
MAN_PATH=/usr/local/share/man/man1

.PHONY: all build man install clean release

all: build

build:
	go build -ldflags="-X 'github.com/vukan322/nuke-node_modules/cmd.version=$(VERSION)'" -o $(BINARY_NAME)

man:
	@mkdir -p docs/man
	@go run cmd/gendocs.go

install: build man
	@echo "Installing binary to $(INSTALL_PATH)..."
	@sudo install -m 755 $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Installing man page to $(MAN_PATH)..."
	@sudo mkdir -p $(MAN_PATH)
	@sudo install -m 644 docs/man/$(BINARY_NAME).1 $(MAN_PATH)/$(BINARY_NAME).1
	@sudo mandb -q 2>/dev/null || true
	@echo "Installation complete. Try: man $(BINARY_NAME)"

uninstall:
	@sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@sudo rm -f $(MAN_PATH)/$(BINARY_NAME).1
	@echo "Uninstalled."

clean:
	@rm -f $(BINARY_NAME)
	@rm -rf docs/

release:
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/vukan322/nuke-node_modules/cmd.version=$(VERSION)'" -o dist/$(BINARY_NAME)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'github.com/vukan322/nuke-node_modules/cmd.version=$(VERSION)'" -o dist/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags="-X 'github.com/vukan322/nuke-node_modules/cmd.version=$(VERSION)'" -o dist/$(BINARY_NAME)-darwin-arm64
	@echo "Release binaries built in dist/"
