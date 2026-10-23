#!/bin/sh
set -eu
version="${RELEASE_VERSION:-0.1.0}"
version="${version#v}"
mkdir -p dist
go build -ldflags="-X main.version=$version" -o dist/pulumi-resource-provider .
pulumi package get-schema ./dist/pulumi-resource-provider > schema.json
pulumi package gen-sdk schema.json --language go --out sdk
pulumi package gen-sdk schema.json --language nodejs --out sdk
python3 scripts/pin-sdk.py
(cd sdk/nodejs && npm install --ignore-scripts --no-audit --no-fund)
