FROM preparedimage:latest AS PREPAREDIMAGE

WORKDIR /usr/app
COPY . .

RUN npm install
RUN npm run build

FROM golang:1.17.1 AS GO
COPY --from=PREPAREDIMAGE /usr/app /usr/app
WORKDIR /usr/app

EXPOSE 18265
CMD ["go", "run", "/usr/app"]
# docker build -t webgame .