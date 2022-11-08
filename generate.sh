#!/bin/bash

set -eux

pushd kc/swagger
java -jar ../../swagger-codegen-cli.jar generate -i ./api/swagger.yaml -l go --disable-examples
# go format the generated files
gofmt -w ./
rm .travis.yml | true
rm git_push.sh | true
