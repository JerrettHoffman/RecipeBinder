.PHONY: build
build:
	go build -o "RecipeBinder.exe" RecipeBinder

.PHONY: run
run: build
	./RecipeBinder.exe
