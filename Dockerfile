#!/bin/bash

FROM golang:1.17.1

ADD . /app

WORKDIR /app

RUN make

ENTRYPOINT ./go-mongo-debug