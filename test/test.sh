#!/usr/bin/env bash

expect_example_to_exit_with_code() {
    here=$(cd "$(dirname $0)" && pwd -P)
    filename="${here}/../demo/${1}"
    expected_code=$2
    go run ${here}/../mdx/main.go $filename > /dev/null 2>&1
    result=$?
    if [ $result -ne $expected_code ]; then
        echo "Failed on assertion for $1. Wanted $2 got $result"
        exit 1
    else
        echo "$1 PASS"
    fi
}

expect_example_to_exit_with_code sample.md 0
expect_example_to_exit_with_code fail.md 1
expect_example_to_exit_with_code linked_code_blocks.md 0
expect_example_to_exit_with_code assertions.md 0
