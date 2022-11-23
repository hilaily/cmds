# TS

a command line tool to transfer timestamp and date time string.

## Install

```shell
go install github.com/hilaily/cmds/ts@latest
```

## Usage

```shell
$ ts 
1669191184 # show current timestamp

$ ts -q 1669191184
2022-11-23 16:13:04 # show date time string

$ ts -q "2022-11-23 16:13:04"
1669191184 # show timestamp
```

