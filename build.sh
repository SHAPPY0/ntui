#!/bin/bash

set -e

package="github.com/shappy0/ntui"
git_rev=$(git rev-parse --short HEAD)
version="v0.1"
output_bin="bin/ntui"

rm -rf "$output_bin"

#build code
go build -ldflags "-w -s -X $package/cmd.Version=$version -X $package/cmd.Commit=$git_rev" -o $output_bin main.go