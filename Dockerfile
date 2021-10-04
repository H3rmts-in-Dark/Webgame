FROM node:16.10.0 AS NODE
FROM rust:1.55.0-slim AS RUST
COPY --from=NODE . .

WORKDIR /usr/app
COPY . .

RUN npm install
RUN npm run build

FROM golang:1.17.1 AS GO
COPY --from=RUST /usr/app /usr/app
WORKDIR /usr/app

EXPOSE 18265
CMD ["go", "run", "/usr/app"]
# docker build -t webgame .