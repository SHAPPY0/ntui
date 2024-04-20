# ntui - A Terminal Based UI To Manage Nomad Clusters.

This is an unofficial open source terminal UI tool to manage your [Hashicorp Nomad](https://www.nomadproject.io/) clusters. The purpose of this tui is to just make it easy to access Nomad services like Jobs, Allocations, Deployments, TaskGroups, Tasks, Logs, Restart/Delete resources.

`ntui` is free to use and currently at its `BETA` stage. We are continuously working on adding more features of ntui, so stay tune and keep using `ntui`.

## Screens:
#### 1. Regions/Namespaces:

![ntui Regions/Namespaces](https://github.com/SHAPPY0/ntui/blob/main/assets/images/regions_namespace.png)

#### 2. Jobs:

![ntui Jobs](https://github.com/SHAPPY0/ntui/blob/main/assets/images/jobs.png)

#### 3. Allocations:
![ntui Allocations](https://github.com/SHAPPY0/ntui/blob/main/assets/images/allocations.png)

#### 4. Versions:
![ntui Versions](https://github.com/SHAPPY0/ntui/blob/main/assets/images/versions.png)

## Installations:
`ntui` can be installed through shell script or `make` through source code. 

#### Install By Make:
1. Clone this git repository.
2. Run `make install`
3. Run `make build`.
4. Make sure to set all the configurations inside `config.json`
5. Run `ntui` to start the tui.

#### Install By shell script:
1. Clone this git repository.
2. Run setup script using `bash ./setup.sh`. It will setup a home directory(`.ntui`) with configs. Make sure to set all the configurations inside `config.json`.
3. Run build script using `bash ./build.sh`. It will build the code in local system.
4. Run it using `./bin/ntui`. More options can be viewed using `./bin/ntui --help`  

## How To Use It:

`ntui` requires some configurations to be set, the default config file should be at user's root home diretory.

Default config json looks like below - 

```json
  {
    "Home_Dir": "",
    "Config_Path": "",
    "Log_Level": "",
    "Log_Dir": "",
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
}
```

### Commands:
```shell
# Run ntui
ntui

# View Help options
ntui help

#  View current ntui version
ntui version

# View config values.
ntui config 
```
### Flags:

Below are the falgs which can be passed while running ntui - 

`-c or --config-path` to set ntui config path.

`--home-dir` to  set home directory of ntui app.

`--host` to set nomad host.

`-l or --log-level` to set the ntui log level.

`--region` to set the nomad region.

`-n or --namespace` to set the nomad namespace.

`-r or --refresh` to set refresh rate to refresh the screen data.

`--skip-verify` to set if skip cetificate verification.

`-t or --token` to set nomad token to perform actions, which requires it.

#### Keys:
```shell
# Global Keys
[1]: To view Nodes
[2]: To view regions and namespaces 
[esc]: To go back to previous screen
[enter]: To select the row

# Jobs Screen
[ctrl+q]: To stop job
[ctrl+s]: To start job

# TaskGroups Screen
[v]: To view job verions

# Versions Screen
[ctrl+v]: To revert the selected job verions

# Allocations Screen
[ctrl+t]: To restart selected task.

# Tasks Screen
[ctrl+t]: To restart selected task.
[l]: To view logs of selected task.

# Logs Screen
[e]: To view stderr logs.
[o]: To view stdout logs.
```

## Developer:
[Saurabh Sharma](http://linkedin.com/in/sharmasaurabh450)
