FROM node:17 AS NODE
FROM rust:1.56.1 AS RUST
COPY --from=NODE . .

COPY . /webgame
WORKDIR /webgame

RUN npm install
RUN npm run build

FROM golang:1.17.2 AS GO
COPY --from=RUST /webgame /webgame
WORKDIR /webgame

EXPOSE 18265
CMD ["go", "run", "/webgame"]