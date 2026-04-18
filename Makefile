BINARY_NAME=nekobotic

build:
	go build -o $(BINARY_NAME) ./cmd/nekobotic/main.go

install: build
	mkdir -p $(HOME)/.local/bin
	cp $(BINARY_NAME) $(HOME)/.local/bin/
	@echo "Add 'export PATH=\$$HOME/.local/bin:\$$PATH' to your .zshrc"
	@echo "Then add 'nekobotic' to the end of the file."

clean:
	rm -f $(BINARY_NAME)