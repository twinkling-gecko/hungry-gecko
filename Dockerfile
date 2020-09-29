FROM alpine:latest

ENV DEBIAN_FRONTEND=noninteractive
ENV PATH=$PATH:/usr/lib/go-1.14/bin
ARG AXIOS_BASE_URL=http://localhost/api/v1/
ENV AXIOS_BASE_URL=$AXIOS_BASE_URL

RUN apk add \
    go \
    nodejs \
    npm \
    ruby \
    curl \
    git \
    nginx

# https://github.com/gliderlabs/docker-alpine/issues/185
RUN mkdir -p /run/nginx

COPY src/macksnow /usr/src/macksnow
WORKDIR /usr/src/macksnow
RUN go build

COPY src/tangerine /usr/src/tangerine
WORKDIR /usr/src/tangerine
RUN npm install

RUN go get -u github.com/pressly/goose/cmd/goose

COPY src/eclipse /usr/src/eclipse
CMD ["sh", "/usr/src/eclipse/start.sh"]
