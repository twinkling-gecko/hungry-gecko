FROM node:latest as builder

EXPOSE 3000

ENV HOST "0.0.0.0"
ENV PORT "3000"

WORKDIR /usr/src/app

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY . .
RUN yarn build
RUN yarn export

CMD ["yarn", "serve"]
