#! /usr/bin/env bash

# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail


function errcho {
  echo "$@" 1>&2
}

# usage: error_exit <error message> <exit code>
function errxit {
    errcho "$1"
    exit "${2:-1}"  ## Return a code specified by $2 or 1 by default.
}

function _zip {
    zip --recurse-paths --verbose --test $@
}

function create_lambda_package {
  if [ -z "${1:-}" ] || [ "${2:-}" ]; then
    errxit 'Usage: create_lambda_package <name>';
  fi

  _lambda_package_name="$1";
  echo "Zipping '${_lambda_package_name}' lambda package."
  _zip "${_lambda_package_name}".zip "${_lambda_package_name}"

  #_zip --grow "${_lambda_package_name}".zip main.go
}

GITSHA=$(git rev-parse HEAD)
if [ -z "$GITSHA" ]  || [ ${#GITSHA} -ne 40 ]; then
  errxit 'Could not determine current git sha.'
fi
echo "Build for version: ${GITSHA}"

REPOSITORY='github.com/samstav/loafers'

function do_go_build_for_lambda {
  _lambda_package_name="$1"
  # The -i flag installs the packages that are dependencies of the target.
  # In the following line, should "main" be ${_lambda_package_name} ?
  GOOS=linux GOARCH=amd64 go build \
    -i \
    -ldflags "-X main.Version=${GITSHA:-HELLO}" \
    -o "${_lambda_package_name}" \
    main.go
  echo "wrote go package '${_lambda_package_name}'"
}

do_go_build_for_lambda main
create_lambda_package main
