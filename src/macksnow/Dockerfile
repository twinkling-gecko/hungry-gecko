FROM golang:latest as builder

WORKDIR /go/src/macksnow

COPY . .

RUN go build


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/macksnow .

CMD ["./macksnow"]
