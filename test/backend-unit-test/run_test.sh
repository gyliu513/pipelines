#!/bin/bash

# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -xe

usage()
{
    echo "usage: run_test.sh
    --results-gcs-dir GCS directory for the test results. Usually gs://<project-id>/<commit-sha>/api_integration_test
    [-h help]"
}

while [ "$1" != "" ]; do
    case $1 in
             --results-gcs-dir )shift
                                RESULTS_GCS_DIR=$1
                                ;;
             -h | --help )      usage
                                exit
                                ;;
             * )                usage
                                exit 1
    esac
    shift
done

if [ -z "$RESULTS_GCS_DIR" ]; then
    usage
    exit 1
fi

GITHUB_REPO=kubeflow/pipelines
BASE_DIR=/go/src/github.com/${GITHUB_REPO}
JUNIT_TEST_RESULT=junit_BackendUnitTestOutput.xml
TEST_DIR=backend/src

cd "${BASE_DIR}/${TEST_DIR}"

# Run test and store the exit code.
echo "Run unit test..."
TEST_RESULT=`go test -v ./... 2>&1`
TEST_EXIT_CODE=$?

# Log the test result
printf '%s\n' "$TEST_RESULT"
# Convert test result to junit.xml
printf '%s\n' "$TEST_RESULT" | go-junit-report > ${JUNIT_TEST_RESULT}

echo "Copy test result to GCS ${RESULTS_GCS_DIR}/${JUNIT_TEST_RESULT}"
gsutil cp ${JUNIT_TEST_RESULT} ${RESULTS_GCS_DIR}/${JUNIT_TEST_RESULT}

exit $TEST_EXIT_CODE