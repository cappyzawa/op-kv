# op-kv 
[![Actions Status](https://github.com/cappyzawa/op-kv/workflows/CI/badge.svg)](https://github.com/cappyzawa/op-kv/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/op-kv)](https://goreportcard.com/report/github.com/cappyzawa/op-kv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This CLI can use op (https://support.1password.com/command-line/) like as key-value.

## Install
go
```bash
$ go get github.com/cappyzawa/op-kv/cmd/op-kv
```

[zdharma/zplugin](https://github.com/zdharma/zplugin)
```zsh
zplugin ice wait'2' lucid as"program" from"gh-r" \
  has"op"
zplugin light cappyzawa/op-kv
```

## Required
* `op`: [1Password command\-line tool: Full documentation](https://support.1password.com/command-line/)

## Usage
This cli required `$XDG_CONFIG_HOME/.op/config` (or `$HOME/.op/config`) file. This file is created by executing [`op signin`](https://support.1password.com/command-line/#sign-in-or-out).
```bash
$ op signin -h
usage: op signin <signinaddress> <emailaddress>

Example account details:

   <signinaddress>   example.1password.com
   <emailaddress>    user@example.com
```

If you set `OP_PASSWORD` as master password to ENV var, `op-kv` command is very easy.

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
  -h, --help                 help for op-kv
      --op-password string   password for 1password 
  -d, --subdomain string     subdomain of 1password

Use "op-kv [command] --help" for more information about a command.
```

if `-p` is not set, `$OP_PASSWORD` is used as password.

And if `-d` is not set, this cli access to latest signin subdomain.

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
  op-kv write <key> <value> [flags]

Flags:
  -h, --help              help for write
  -p, --password string   register password to item(key)
  -u, --username string   register username to item(key)

Global Flags:
      --op-password string   password for 1password ()
  -d, --subdomain string     subdomain of 1password ()
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
