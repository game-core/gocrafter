#!/bin/bash

GAME_PROTO_BASE_DIR="./docs/proto/api/game"

find "$GAME_PROTO_BASE_DIR" -type f -name '*.proto' -exec sh -c '
  for file do
    filename=$(basename "$file")
    protoc --go_out=. --go-grpc_out=. --proto_path="${file%/*}" "$file"
  done
' sh {} +
