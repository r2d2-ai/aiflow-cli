<!--
title: AIflow
weight: 5020
pre: "<i class=\"fa fa-terminal\" aria-hidden=\"true\"></i> "
-->

# Commands

- [build](#build) - Build the AIflow application
- [create](#create) - Create a AIflow application project
- [help](#help)  - Help about any command
- [imports](#imports) - Manage project dependency imports
- [install](#install) - Install a AIflow contribution/dependency
- [list](#list) - List installed AIflow contributions
- [plugin](#plugin) - Manage CLI plugins
- [update](#update) - Update an application contribution/dependency

### Global Flags
```
  --verbose   verbose output
```

  
## build

This command is used to build the application.

```
Usage:
  AIflow-cli build [flags]

Flags:
  -e, --embed         embed configuration in binary
  -f, --file string   specify a AIflow.json to build
  -o, --optimize      optimize build
      --shim string   use shim trigger   
```
_**Note:** the optimize flag removes unused trigger, acitons and activites from the built binary._


### Examples
Build the current project application

```bash
$ AIflow build
```
Build an application directly from a AIflow.json

```bash
$ AIflow build -f AIflow.json
```
_**Note:** this command will only generate the application binary for the specified json and can be run outside of a AIflow application project_

## create

This command is used to create a AIflow application project.

```
Usage:
  AIflow create [flags] [appName]

Flags:
      --cv string     specify core library version (ex. master)
  -f, --file string   specify a AIflow.json to create project from
```

_**Note:** when using the --cv flag to specify a version, the exact version specified might not be used the project.  The application will install the version that satisfies all the dependency constraints.  Typically this flag is used when trying to use the master version of the core library._

### Examples

Create a base sample project with a specific name:

```
$ AIflow create my_app
```

Create a project from an existing AIflow application descriptor:

```
$ AIflow create -f myapp.json
```

## help

This command shows help for any AIflow commands.

```
Usage:
  AIflow help [command]
```  

### Examples
Get help for the build command:

```bash
$ AIflow help build
```
## imports

This command helps manage project imports of contributions and dependencies.

```
Usage:
  AIflow imports [command]

Available Commands:
  sync     sync Go imports to project imports
  resolve  resolve project imports to installed version
  list     list project imports
```   

## install

This command is used to install a AIflow contribution or dependency.

```
Usage:
  AIflow install [flags] <contribution|dependency>

Flags:
  -f, --file string      specify contribution bundle
  -r, --replace string   specify path to replacement contribution/dependency
```
      
### Examples
Install the basic REST trigger:

```bash
$ AIflow install github.com/r2d2-ai/aiflow/contrib/trigger/rest
```
Install a contribution that you are currently developing on your computer:

```bash
$ AIflow install -r /tmp/dev/myactivity github.com/myuser/myactivity
```

Install a contribution that is being developed by different person on their fork:

```bash
$ AIflow install -r github.com/otherusr/myactivity@master github.com/myuser/myactivity
```

## list

This command lists installed contributions in your application

```
Usage:
  AIflow list [flags]

Flags:
      --filter string   apply list filter [used, unused]
  -j, --json            print in json format (default true)
      --orphaned        list orphaned refs
```  
_**Note** orphaned refs are `ref` entries that use an import alias (ex. `"ref": "#log"`) which has no corresponding import._

### Examples
List all installed contributions:

```bash
$ AIflow list
```
List all contributions directly used by the application:

```bash
$ AIflow list --filter used
```
_**Note:** the results of this command are the only contributions that will be compiled into your application when using `AIflow build` with the optimize flag_


## plugin

This command is used to install a plugin to the AIflow CLI.

```
Usage:
  AIflow plugin [command]

Available Commands:
  install     install CLI plugin
  list        list installed plugins
  update      update plugin
```      

### Examples
List all installed plugins:

```bash
$ AIflow plugin list
```
Install the legacy support plugin:

```bash
$ AIflow plugin install github.com/r2d2-ai/legacybridge/cli`
```
_**Note:** more information on the legacy support plugin can be found [here](https://github.com/r2d2-ai/legacybridge/tree/master/cli)_

Install and use custom plugin:

```
$ AIflow plugin install github.com/myuser/myplugin

$ AIflow `your_command`
```
<br>
More information on AIflow CLI plugins can be found [here](plugins.md)

## update

This command updates a contribution or dependency in the project.

```
Usage:
  AIflow update [flags] <contribution|dependency>

```   
### Examples
Update you log activity to master:

```bash
$ AIflow update github.com/r2d2-ai/aiflow/contrib/activity/log@master
```

Update your AIflow core library to latest master:

```bash
$ AIflow update github.com/r2d2-ai/aiflow/core@master
```
