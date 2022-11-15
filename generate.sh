#!/bin/bash

set -eux

# remove the old generated files
rm ./kc/swagger/model_*.go | true
rm ./kc/swagger/docs/*.md | true

# if swagger-codegen-cli.jar is not available, exit
if [ ! -f swagger-codegen-cli.jar ]; then
    echo "swagger-codegen-cli.jar is not available. Please download with: wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.35/swagger-codegen-cli-3.0.35.jar -O ./swagger-codegen-cli.jar"
    exit 0
fi

pushd kc/swagger
java -jar ../../swagger-codegen-cli.jar generate -l go -i ./api/swagger.yaml --output ./
# go format the generated files

gofmt -w ./
rm .travis.yml || true
rm git_push.sh || true
