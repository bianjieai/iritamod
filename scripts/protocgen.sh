#!/usr/bin/env bash

set -eo pipefail

cd proto

buf generate --template buf.gen.gogo.yaml

cd ..

# move proto files to the right places
cp -r github.com/bianjieai/iritamod/* ./
rm -rf github.com