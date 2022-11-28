#! /usr/bin/env bash
set -euox pipefail

mkdir -p dist
rm -f dist/crds.yaml

for f in $(ls package/crds/*.yaml); do
    cat "$f" >> dist/crds.yaml
done