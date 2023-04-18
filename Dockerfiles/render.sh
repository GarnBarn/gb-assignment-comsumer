#!/bin/bash

render() {
sedStr="
  s!%%BUILD_APP%%!$BUILD_APP!g;
"

sed -r "$sedStr" $1
}

render ./Dockerfiles/Dockerfile.template > Dockerfile