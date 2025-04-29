#!/bin/bash
set -e

# move to the root dir of the package
rd=$(git rev-parse --show-toplevel)
cd $rd
  
protoc \
--proto_path=. "aed.proto" \
--proto_path=$(dirname $(dirname "$rd")) \
"--go_out=." --go_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
"--go-grpc_out=." --go-grpc_opt=paths=source_relative

protoc \
--proto_path=. "aed-grpc.proto" \
--proto_path=$(dirname $(dirname "$rd")) \
"--go_out=." --go_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
"--go-grpc_out=." --go-grpc_opt=paths=source_relative
  
git add *.go

# if there's TS project in the package, generate a protobuf file for TS
rm -rf node_modules
npm i

protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
  --proto_path=. \
  --proto_path=$(dirname $(dirname "$rd")) \
  --ts_proto_out=. \
  --ts_proto_opt=esModuleInterop=true \
  --ts_proto_opt=outputServices=grpc-js \
  aed.proto

protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
  --proto_path=. \
  --proto_path=$(dirname $(dirname "$rd")) \
  --ts_proto_out=. \
  --ts_proto_opt=esModuleInterop=true \
  --ts_proto_opt=outputServices=grpc-js \
  aed-grpc.proto

npm run build
git add build/

git add *.ts
rm -rf node_modules