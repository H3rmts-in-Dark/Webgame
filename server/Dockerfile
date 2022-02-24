FROM golang:1.17.4-alpine as build

RUN apk update && apk add openssl ca-certificates tzdata
RUN update-ca-certificates
RUN adduser  --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "10001" "appuser"

ENV PortHTTP=80
ENV PortHTTPS=443
ENV APIPORT=18266
EXPOSE 80
EXPOSE 443
EXPOSE 18266

WORKDIR /server
COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server/server .

FROM scratch
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

COPY --from=build /server/server /server
COPY server.yml /

USER appuser:appuser

ENTRYPOINT ["/server"]