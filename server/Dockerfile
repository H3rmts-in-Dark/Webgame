FROM golang:1.17.4-alpine

WORKDIR /server
COPY . .

ENV PortHTTP=80
ENV PortHTTPS=443
ENV APIPORT=18266
EXPOSE 80
EXPOSE 443
EXPOSE 18266

# create certificates
RUN apk add openssl && ./createCerts.sh

RUN go install

CMD ["go", "run", "."]