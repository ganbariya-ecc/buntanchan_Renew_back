#!/bin/bash

# スクリプトのディレクトリを取得
CURRENT=$(cd $(dirname $0);pwd)

echo $CURRENT

# スクリプトのディレクトリに移動
cd $CURRENT

cd ./sdks/proto

apt update
apt install protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. authsdk.proto

# スクリプトのディレクトリに移動
cd $CURRENT

# GRPC の鍵ディレクトリに移動
cd ./sdks/cert

# generate ca.key 
openssl genrsa -out ca.key 4096
# generate certificate
openssl req -new -x509 -key ca.key -sha512 -subj "/C=SE/ST=HL/O=Auth, INC." -days 3650 -out ca.cert
# generate the server key
openssl genrsa -out server.key 4096
# Generate the csr
openssl req -new -key server.key -out server.csr -config certificate.conf
# 
openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.key -CAcreateserial -out server.crt -days 3650 -sha512 -extfile certificate.conf -extensions req_ext

cd $CURRENT
# ファイルをコピー
cp ./sdks/cert/server.crt ./server.crt
cp ./sdks/cert/server.csr ./server.csr
cp ./sdks/cert/server.key ./server.key