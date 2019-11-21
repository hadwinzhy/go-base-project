BINARY=app.bin

default:
	go build -o target/${BINARY} -v

prod:
	go build -o target/${BINARY} -v


install:
	go mod download

# Installs our project: copies binaries
dev:
	fresh -c deploy/fresh_runner.conf

test:
	go test

# cover:
#   go test -v -cover -coverprofile=coverage.out

# Cleans our project: deletes binaries
clean:
	rm -rf target/*

.PHONY:  default clean install dev prod test
