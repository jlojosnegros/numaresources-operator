#!/usr/bin/env bash

source hack/common.sh

NO_COLOR=""
if ! which tput &> /dev/null 2>&1 || [[ $(tput -T$TERM colors) -lt 8 ]]; then
  echo "Terminal does not seem to support colored output, disabling it"
  NO_COLOR="-ginkgo.noColor"
fi

# Make sure that we always properly clean the environment
trap '{ if [ "${NO_TEARDOWN}" = false ]; then
            echo "Running NROScheduler uninstall test suite";
            ${BIN_DIR}/e2e-sched-uninstall.test ${NO_COLOR} --ginkgo.v --ginkgo.reportFile=/tmp/artifacts/sched-uninstall;
        fi }' EXIT SIGINT SIGTERM SIGSTOP

# Run install test suite
echo "Running NROScheduler install test suite"
${BIN_DIR}/e2e-sched-install.test ${NO_COLOR} --ginkgo.v --ginkgo.failFast --ginkgo.reportFile=/tmp/artifacts/sched-install

# The install failed, no taste to continue
if [ $? -ne 0 ]; then
    echo "Failed to install NROScheduler"
    exit 1
fi

echo "Running Functional Tests: ${GINKGO_SUITS}"
# -v: print out the text and location for each spec before running it and flush output to stdout in realtime
# -r: run suites recursively
# --failFast: ginkgo will stop the suite right after the first spec failure
# --flakeAttempts: rerun the test if it fails
# -requireSuite: fail if tests are not executed because of missing suite
${BIN_DIR}/e2e-sched.test ${NO_COLOR} --ginkgo.v --ginkgo.failFast --ginkgo.flakeAttempts=2 --ginkgo.reportFile=/tmp/artifacts/e2e-sched
