# op-kv [![https://concourse\.ik\.am:14161/api/v1/teams/cappyzawa/pipelines/op\-kv/jobs/test\-master/badge](https://concourse.ik.am:14161/api/v1/teams/cappyzawa/pipelines/op-kv/jobs/test-master/badge)](https://concourse.ik.am:14161/teams/cappyzawa/pipelines/op-kv)
This CLI can use op (https://support.1password.com/command-line/) like as key-value.

## Install
```bash
$ go get github.com/cappyzawa/op-kv/cmd/op-kv
```
or
```bash
# for mac
$ curl -sL https://github.com/cappyzawa/op-kv/releases/download/v1.0.0/op-kv-darwin-amd64 > op-kv
$ chmod +x op-kv
```

zplug
```zsh
zplug "cappyzawa/op-kv", \
    from:gh-r, \
    as:command, \
    rename-to:op-kv, \
    on:"stedolan/jq", \
    on:"cappyzawa/get-op"
```

## Required
* `op`: [1Password command\-line tool: Full documentation](https://support.1password.com/command-line/)
* `jq`: [stedolan/jq: Command\-line JSON processor](https://github.com/stedolan/jq)

## Usage

```bash
$ op-kv -h
use "op" like as kv

Usage:
  op-kv [flags]
  op-kv [command]

Available Commands:
  help        Help about any command
  list        Display item titles
  read        Display one password of specified item by UUID or name
  write       Generate one password by specified item and password

Flags:
  -h, --help   help for op-kv

Use "op-kv [command] --help" for more information about a command.
```

**It need to execute `op signin` in advanced.**
### Read
```bash
$ op-kv read -h
Display one password of specified item by UUID or name

Usage:
  op-kv read [<UUID>|<name>] [flags]

Flags:
  -h, --help   help for read
```
This Command is same as below.
```bash
$ op get item [<UUID>|<name>] | jq -r '.details.fields[] | select(.designation=="password").value'
```
This can adjust only _item_ subcommand.

### write
```bash
$ op-kv write -h 
Generate one password by specified item and password

Usage:
  op-kv write <item> <password> [flags]

Flags:
  -h, --help   help for write
```
This Command is same as below.
```bash
$ D=$(op get template login | jq -c '.fields[1].value = <password>' | op encode)
$ op create item login $D --title=<item>
```
This can adjust only _login_ template.

### list
```bash
$ op-kv list -h
Display item titles

Usage:
  op-kv list [flags]

Flags:
  -h, --help   help for list
```
This Command is same as below.
```bash
$ op list items | jq -r ".[].overview.title"
```
