#! /bin/sh
go env -w GO111MODULE=on

go env -w GOPROXY=https://goproxy.io,direct
