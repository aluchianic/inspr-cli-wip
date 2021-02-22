### Inspr commands to operate with Cluster and Local files (init Workspace, App)
```
Inspr orchestrator

Usage:
  inspr [command]

Available Commands:
  auth        Manage Inspr authorization
  cluster     Manage Inspr cluster
  config      Manage Inspr configs
  help        Help about any command
  hub         Manage Inspr Hub
  version     Inspr cli version

Flags:
  -h, --help   help for inspr

Use "inspr [command] --help" for more information about a command.

```
### Inspr Cluster commands
```
Manage Inspr cluster

Usage:
  inspr cluster [flags]
  inspr cluster [command]

Aliases:
  cluster, clusters

Available Commands:
  deploy      Deploy application on cluster
  describe    DescribeApp dApp with it dependencies, channel types and third parties
  remove      Remove deployed dApp from cluster
  stats       Return statistics from clusters, by default returns statistics only for running clusters

Flags:
      --apps string        Directory to search applications in
  -h, --help               help for cluster
      --workspace string   Inspr workspace config path

Use "inspr cluster [command] --help" for more information about a command.

```

### Inspr Cluster commands
```
Manage Inspr cluster

Usage:
  inspr cluster [flags]
  inspr cluster [command]

Aliases:
  cluster, clusters

Available Commands:
  deploy      Deploy application on cluster
  describe    DescribeApp dApp with it dependencies, channel types and third parties
  remove      Remove deployed dApp from cluster
  stats       Return statistics from clusters, by default returns statistics only for running clusters

Flags:
      --apps string        Directory to search applications in
  -h, --help               help for cluster
      --workspace string   Inspr workspace config path

Use "inspr cluster [command] --help" for more information about a command.

```

### Inspr Hub commands
```
Manage Inspr Hub

Usage:
  inspr hub [command]

Available Commands:
  find        Find application in Inspr Hub
  get         Get applications from Inspr Hub

Flags:
  -h, --help   help for hub

Use "inspr hub [command] --help" for more information about a command.

```

### Usage :

### Adding a command

* Create a new directory under `handler/`
* Add an import to the `handler/handler.go`
* Add the `get.go`, `create.go`, or any other other handlers in this new
  directory.
* Expose a `XXXAddCommand` function to let others attach commands to your new
  command.
* Attach your command to another by calling their `XXXAddCommand` function.
* See `handler/example` for example.
