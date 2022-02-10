openssl req -newkey rsa:8192 \
  -new -nodes -x509 \
  -days 7300 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"