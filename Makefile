install: main.go
	@go build .
	@cp http_serve $(HOME)/http_serve
	@rm -f http_serve
	@echo 'export PATH=$PATH:$(HOME)' >> ~/.bashrc;\
	@echo compiled and installed !

uninstall:
	@rm -f $(HOME)/http_serve
	@echo uninstalled
