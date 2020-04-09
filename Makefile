BINARY=app.bin

install:
	. deploy/env.sh
	go mod download

local:
	air -c deploy/local/air.conf

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
