app_name = "gimg"

.PHONY: build

build:
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	go build -v -o build/$(app_name) build/main.go

.PHONY: clean

clean:
	rm build/$(app_name)
