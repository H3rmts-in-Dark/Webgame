openssl genpkey -algorithm RSA -out server.key
#openssl genrsa -out server.key 4096
openssl ecparam -genkey -name secp384r1 -out server.key
yes '' | openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
mv server.pem server/server.pem
mv server.key server/server.key