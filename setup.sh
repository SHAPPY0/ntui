#!/bin/bash

set -e

os=$(uname)
dir_name=".ntui"
config_file="config.toml"
home_dir=""

if [ "$os" = "Darwin" ]; then
    home_dir="$HOME"/"$dir_name"
    if [ ! -f "$home_dir"/"$config_file" ]; then
        mkdir -p "$home_dir"
        touch "$home_dir"/"$config_file"
        echo $(config_file) > "$home_dir"/"$config_file"
    fi
elif [ "$os" = "Linux" ]; then
    home_dir="$HOME"/"$dir_name"
    if [ ! -f "$home_dir"/"$config_file" ]; then
        mkdir -p "$home_dir"
        touch "$home_dir"/"$config_file"
        echo $(config_file) > "$home_dir"/"$config_file"
    fi
elif [ "$os" = "MINGW64_NT-10.0" ]; then
    home_dir="$HOME"\\"$dir_name"
    if [ ! -f "$home_dir"\\"$config_file" ]; then
        mkdir -p "$home_dir"
        touch "$home_dir"\\"$config_file"
    fi
else
    echo "Unsupported operating system."
    exit 1
fi




