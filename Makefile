bin: bin/ganesha_exporter_darwin_amd64 bin/ganesha_exporter_linux_amd64
bin: bin/ganesha_exporter_darwin_arm64 bin/ganesha_exporter_linux_arm64

bin/ganesha_exporter_darwin_amd64:
	@mkdir -p bin
	@echo "Compiling ganesha_exporter..."
	GOOS=darwin GOARCH=amd64 go build -o $@ *.go

bin/ganesha_exporter_linux_amd64:
	@mkdir -p bin
	@echo "Compiling ganesha_exporter..."
	GOOS=linux GOARCH=amd64 go build -o $@ *.go

bin/ganesha_exporter_darwin_arm64:
	@mkdir -p bin
	@echo "Compiling ganesha_exporter..."
	GOOS=darwin GOARCH=arm64 go build -o $@ *.go

bin/ganesha_exporter_linux_arm64:
	@mkdir -p bin
	@echo "Compiling ganesha_exporter..."
	GOOS=linux GOARCH=arm64 go build -o $@ *.go