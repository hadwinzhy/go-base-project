BINARY=go_build

default:
	go build -o output/${BINARY} -v

prod:
	go build -o output/${BINARY} -v


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
	rm -rf go_build

.PHONY:  default clean install dev prod test
