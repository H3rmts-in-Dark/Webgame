openssl req -newkey rsa:8192 \
  -new -nodes -x509 \
  -days 7300 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=DE/ST=Bayern/L=Donauwörth/O=H3rmts in Dark/OU=Unit/CN=localhost"