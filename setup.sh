#!/bin/bash

set -e

os=$(uname)
dir_name=".ntui"
config_file="config.json"
home_dir=""
config_data='{
    "Home_Dir": "",
    "Config_Path": "",
    "Log_Level": "info",
    "Refresh_Rate": 5, 
    "Nomad_Server_Base_Url": "",
    "Nomad_Http_Auth": "",
    "Nomad_Token": "",
    "Nomad_Region": "",
    "Nomad_Namespace": "",
    "Nomad_Cacert": "",
    "Nomad_Capath": "",
    "Nomad_Client_Cert": "",
    "Nomad_Client_Key": "",
    "Nomad_Tls_Server": "",
    "Nomad_Skip_Verify": true
}'

if [ "$os" = "Darwin" ]; then
    home_dir="$HOME"/"$dir_name"
    if [ ! -f "$home_dir"/"$config_file" ]; then
        mkdir -p "$home_dir"
        touch "$home_dir"/"$config_file"
        echo "$config_data" > "$home_dir"/"$config_file"
    fi
elif [ "$os" = "Linux" ]; then
    home_dir="$HOME"/"$dir_name"
    if [ ! -f "$home_dir"/"$config_file" ]; then
        mkdir -p "$home_dir"
        touch "$home_dir"/"$config_file"
        echo "$config_data" > "$home_dir"/"$config_file"
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




