FROM node:18-alpine as BUILD
COPY --from=rust:1.57.0-alpine /usr/local/cargo /usr/local/cargo
ENV PATH="$PATH:/usr/local/cargo/bin"

# add curl to get wasm-pack, add build-base for cc linker
RUN apk add curl && apk add build-base

# setup rust
RUN rustup default stable && rustup target add wasm32-unknown-unknown

# install wasm-pack
RUN cargo install wasm-pack

WORKDIR /webgame

COPY ./package.json ./package-lock.json ./
RUN npm install --no-update-notifier

COPY ./src/css/_smui-theme.scss ./src/css/vars.scss ./src/css/
RUN npm run theme --no-update-notifier

COPY ./webassembly ./webassembly
RUN npm run wasm --no-update-notifier

COPY . .
CMD npm run dev --no-update-notifier