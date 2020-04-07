BINARY=app.bin

install:
	. deploy/env.sh
	go mod download

local:
	fresh -c deploy/dev/fresh_runner.conf

dev:
	go build -o target/${BINARY} -v

prod:
	go build -o target/${BINARY} -v

test:
	go test

clean:
	rm -rf target/*

deepclean:
	rm -rf target/*
	go clean -modcache

.PHONY: install local dev prod test clean deepclean
