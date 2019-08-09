FROM golang:latest as builder
RUN mkdir /go/src/go-leboncoin
WORKDIR /go/src/go-leboncoin
ADD . ./
RUN go get -d -v && go build -o goleboncoin

FROM frolvlad/alpine-glibc:latest
LABEL maintainer="francois.allais@hotmail.com"

ENV RUN_ON_STARTUP=true

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/goleboncoin /app/

CMD /app/goleboncoin --filters_file /app/filters.yaml  --run_on_startup=$RUN_ON_STARTUP
