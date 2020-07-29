# ----- production用ビルドステージ ------

FROM golang:latest as builder

WORKDIR /go/src/macksnow

COPY . .

RUN go build


# ----- production実行ステージ ------

FROM alpine:latest as production

WORKDIR /root/

COPY --from=builder /go/src/macksnow .

CMD ["./macksnow"]


# ----- development実行ステージ ------
# NOTE: volumesでソースコードを/go/src/macksnowにマウントする前提

FROM golang:latest as development

EXPOSE 4000

WORKDIR /go/src/macksnow

RUN go get github.com/pilu/fresh
RUN go get github.com/swaggo/swag/cmd/swag

CMD ["fresh"]

