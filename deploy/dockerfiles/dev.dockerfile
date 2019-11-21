########## build image  ##########
FROM hadwinzhy/golang:latest as builder

WORKDIR /go/src/app

EXPOSE 8081

COPY go.mod go.sum Makefile ./

RUN make install

COPY . .

RUN make

########## runtime image ##########

FROM ubuntu:latest

COPY --from=builder /go/src/app/target/app.bin /app/

COPY --from=builder /go/src/app/configs /app/

RUN chmod 777 /app/app.bin

ENTRYPOINT [ "sh",  "-c", "/app/app.bin serve" ]
