#!/bin/bash

find "./docs/proto/api/game" -type f -name '*.proto' -exec sh -c '
  for file do
    filename=$(basename "$file")
    protoc --go_out=. --go-grpc_out=. --proto_path="${file%/*}" "$file"
  done
' sh {} +

find "./docs/proto/api/multi" -type f -name '*.proto' -exec sh -c '
  for file do
    filename=$(basename "$file")
    protoc --go_out=. --go-grpc_out=. --proto_path="${file%/*}" "$file"
  done
' sh {} +
