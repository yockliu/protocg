# protocg

[中文](README_中文.md)

## Use

Use both local file and git-remote as source for protoc.

Need install `git` `protoc` and language-plugin yourself, and make them work good.

## Install

```
go install github.com/yockliu/protocg:v0.0.3
```

## How

1. Add `protocg.yaml`

2. run `protocg gen`


### protocg.yaml

| param | type | desc |
|----|----|----|
| version | string | version of the config file, no use now |
| name | string | name of the config file, no use now |
| opts | string array | opts of protoc, like `--go_out=./pb` |
| src.local  | string array | local proto source path |
| src.remote  | remote array | git-remote proto source |
| proto_path.local | string array | local proto_path |
| proto_path.remote  | remote array | git-remote proto_path |

* remote:

| param | type | desc |
|----|----|----|
| repo | string | git remote repo url |
| refs | string | refs of git (branch or tag) | 
| path | string array | proto file path in git | 

see：[example/protocg.yaml](example/protocg.yaml)

## Keywords

`protoc` `remote` `git` `version control`