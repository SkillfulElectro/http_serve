install: main.go
	@go build .
	@cp http_serve /usr/bin
	@rm -f http_serve
	@echo compiled and installed !

uninstall:
	@rm -f /usr/bin/http_serve
	@echo uninstalled
