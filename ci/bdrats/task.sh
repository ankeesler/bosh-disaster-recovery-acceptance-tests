#!/usr/bin/env bash

set -e
set -u

pushd $BBR_RELEASE_PATH
  tar xvf *.tar
  export BBR_BINARY_PATH="$PWD/releases/bbr"
popd

BOSH_SSH_PRIVATE_KEY_PATH="$(mktemp)"
echo "$BOSH_SSH_PRIVATE_KEY" > $BOSH_SSH_PRIVATE_KEY_PATH
chmod 400 $BOSH_SSH_PRIVATE_KEY_PATH

export GOPATH="$PWD"
export PATH="$PATH:$GOPATH/bin"

pushd src/github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests
  scripts/_run_acceptance_tests.sh
popd
