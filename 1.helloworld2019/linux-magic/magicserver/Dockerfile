FROM golang:alpine as builder

RUN mkdir -p /go/src/magic
ADD . /go/src/magic
WORKDIR /go/src/magic 
RUN go build -o magic . 

FROM alpine

WORKDIR /opt/
COPY --from=builder /go/src/magic/magic /usr/bin/magic
COPY --from=builder /go/src/magic/template /opt/template
