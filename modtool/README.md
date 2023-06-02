# Modtool

## Install

```shell
go install github.com/hilaily/cmds/modtool@latest
```

## tag subcommand

### [show] show remote tags and local tags

```shell
modtool tag show
```

output

```shell
remote tags:
config/v1.1.4-pre01     config/v1.1.4-ops01     config/v1.1.3                   config/v1.1.2           config/v1.1.1
config/v1.1.0           config/v1.0.3           config/v1.0.2                   config/v1.0.2-pre6      config/v1.0.2-pre5
           
local tags:
config/v1.1.3                   config/v1.1.2           config/v1.1.1           config/v1.1.0           config/v1.0.3
config/v1.0.2                   config/v1.0.2-pre6      config/v1.0.2-pre5      config/v1.0.2-pre4      config/v1.0.2-pre3 
```

### [new] create a new tag and push to remote

```shell
modtool tag new major
# v1.2.3 => v2.0.0

modtool tag new minor
# v1.2.3 => v1.3.0

modtool tag new patch
# v1.2.3 => v1.2.4

modtool tag new beta
# v1.2.3 => v1.2.4-beta01
# v1.2.3-beta02 => v1.2.4-beta03

modtool tag new pre
# v1.2.3 => v1.2.4-pre01

modtool tag new -p=false patch
# just print the tag name will create, don't create it really.
```

## rename subcommand

used to rename a mod.

```bash
modtool rename <old_name> <new_name>
```

and then you should modify go.mod manually(this command just modify go source file).

