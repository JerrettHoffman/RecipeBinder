.PHONY: httpServer
httpServer:
	go build -o "RecipeBinder.exe" RecipeBinder/cmd/httpServer

.PHONY: run
run: httpServer
	./RecipeBinder.exe

.PHONY: all
all: httpServer
