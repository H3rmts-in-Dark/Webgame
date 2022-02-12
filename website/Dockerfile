FROM node:17-alpine as BUILD
COPY --from=rust:1.57.0-alpine /usr/local/cargo /usr/local/cargo
ENV PATH="$PATH:/usr/local/cargo/bin"

# add curl to get wasm-pack, add build-base for cc linker
RUN apk add curl && apk add build-base

# setup rust
RUN rustup default stable && rustup target add wasm32-unknown-unknown

# install wasm-pack
RUN curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh

WORKDIR /webgame
COPY . .

RUN npm install --no-update-notifier

CMD npm run build --no-update-notifier