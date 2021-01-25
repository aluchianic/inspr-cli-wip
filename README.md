### Inspr commands to operate with Cluster and Local files (init Workspace, App)
```
Usage:
  inspr [command]

Available Commands:
  deploy      [Cluster] Deploy Workspace on cluster if no arguments passed assuming that Workspace is current directory.
  describe    Describe dApp with it dependencies, channel types and third parties
  help        Help about any command
  hub         Commands to operate with Inpsr HUB
  init        [Workspace] Initialize Inspr application
  register    Register account to get token for Inspr
  remove      [Cluster] Remove deployed dApp from cluster
  stats       [Cluster] Return statistics from clusters, by default returns statistics only for running clusters
  version     Print the version number of Inspr cli

Flags:
  -h, --help   help for inspr

Use "inspr [command] --help" for more information about a command.
```
### Inspr HUB commands
```
Usage:
  inspr hub [command]

Available Commands:
  describe    DescribeApp dApp with it dependencies, channel types and third parties
  find        Find dApp in Inpsr Hub
  get         Get dApp from Inspr Hub

Flags:
  -h, --help   help for hub

Use "inspr hub [command] --help" for more information about a command.
```