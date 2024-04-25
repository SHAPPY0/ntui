#!/bin/bash

set -e

package="github.com/shappy0/ntui"
git_rev=$(git rev-parse --short HEAD)
version="v1.0.0"
output_bin="bin/ntui"

rm -rf "$output_bin"

echo "===> Bulding ntui"

# build code
go build -ldflags "-w -s -X $package/cmd.Version=$version -X $package/cmd.Commit=$git_rev" -o $output_bin main.go

# move binary in targeted folder
set -- "/usr/local/bin" "/usr/bin" "/opt/bin"

while [ -n "$1" ]; do
    # Check if destination is in path.
    if echo "$PATH"|grep "$1" -> /dev/null ; then
        if cp $output_bin "$1" ; then
            echo ""
            echo "Done!"
            exit 0
        else
            echo ""
            echo "We'd like to move the ntui executable in $1. Please enter your password."
            if sudo cp $output_bin "$1" ; then
                echo ""
                echo "Done!"
                exit 0
            fi
        fi
    fi
    shift
done

echo "could not find supported destination path in \$PATH"
exit 1