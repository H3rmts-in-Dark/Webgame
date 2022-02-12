FROM golang:1.17.4-alpine

RUN apk add openssl

ENV PortHTTP=80
ENV PortHTTPS=443
ENV APIPORT=18266
EXPOSE 80
EXPOSE 443
EXPOSE 18266

WORKDIR /server
COPY . .

# create certificates TODO should be chaned to 1 certificate
#RUN ./certs/generate.sh

RUN go install && go build -o server

CMD ["./server"]