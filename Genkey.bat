rm -rf ./nginx/keys
mkdir ./nginx/keys
openssl genrsa -out ./nginx/keys/server.key 4096
openssl req -out ./nginx/keys/server.csr -key ./nginx/keys/server.key -new
openssl x509 -req -days 3650 -signkey ./nginx/keys/server.key -in ./nginx/keys/server.csr -out ./nginx/keys/server.crt

cd ./auth/src/keys
openssl genrsa -out server.key 4096
openssl req -out server.csr -key server.key -new
openssl x509 -req -days 3650 -signkey server.key -in server.csr -out server.crt