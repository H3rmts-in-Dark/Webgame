#genpkey -algorithm RSA -out server.key  2> /dev/null
openssl genrsa -out server.key 4096 2> /dev/null
openssl ecparam -genkey -name secp384r1 -out server.key
(yes '' | openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650) 2> /dev/null