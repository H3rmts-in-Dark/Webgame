FROM node:17-alpine as BUILD
COPY --from=rust:1.57.0-alpine /usr/local/cargo /usr/local/cargo
ENV PATH="$PATH:/usr/local/cargo/bin"

WORKDIR /webgame
COPY . .

# add curl to get wasm-pack, add build-base for cc linker
RUN apk add curl
RUN apk add build-base

# setup rust
RUN rustup default stable
RUN rustup target add wasm32-unknown-unknown

# install wasm-pack
RUN curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh


RUN npm install
RUN npm run build


# switch to go for serving
FROM golang:1.17.4-alpine

WORKDIR /webgame
COPY --from=BUILD /webgame/site ./site
COPY server ./server
COPY scripts ./scripts
COPY ["config.json", "go.mod", "go.sum", "main.go", "./"]

ENV PORT=443
ENV APIPORT=18040
EXPOSE 18040

RUN go install

# create certificates
RUN apk add openssl
RUN scripts/createCerts.sh

CMD ["go", "run", "."]