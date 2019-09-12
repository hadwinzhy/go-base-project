########## build image  ##########
FROM golang:1.13.0 as builder

WORKDIR /go/src/app
EXPOSE 8080

ENV GIN_MODE release
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

##### install vendor
COPY go.mod go.sum Makefile ./
RUN make install


##### build binary
COPY . .
RUN make prod


FROM ubuntu:latest AS base
COPY --from=builder /go/src/app/build/ /bin/
ENTRYPOINT [ "sh",  "-c", "/bin/go_build" ]
