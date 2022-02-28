openssl req -newkey rsa:8192 -x509 -nodes -days 7300 \
	-subj "/C=DE/ST=Bayern/L=Donauwörth/O=H3rmts in Dark/OU=Unit/CN=localhost" \
	-keyout localhost.key \
	-out localhost.pem

openssl verify localhost.pem
openssl verify -CAfile localhost.pem localhost.pem

(yes securePassword) | openssl pkcs12 -export \
	-inkey localhost.key \
	-in localhost.pem \
	-out localhost.pfx

# (yes securePassword) |
