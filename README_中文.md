# protocg

[English](README.md)

## 用途

可以直接使用`远程git库`作为源生成protobuf源码，兼容`本地文件`

需要自行安装好 `git`, `protoc` 及自己需要的语言插件，并可以正常运行

## 安装

```
go install github.com/yockliu/protocg:v0.0.3
```

## 使用

1. Add `protocg.yaml`

2. run `protocg gen`

### protocg.yaml

| 参数 | 类型 | 描述 |
|----|----|----|
| version | string | 配置文件格式版本号，暂无用处 |
| name | string | 配置文件名称，暂无用处 |
| opts | string array | protoc 的 opts, 类似`--go_out=./pb` |
| src.local  | string array | 本地 proto 文件路径 |
| src.remote  | remote array | git-remote proto source |
| proto_path.local | string array | 本地 proto_path 文件路径 |
| proto_path.remote  | remote array | git-remote proto_path |

* remote:

| 参数 | 类型 | 描述 |
|----|----|----|
| repo | string | git 远程库地址 |
| refs | string | 远程库refs (branch or tag) | 
| path | string array | proto文件在git库中的文件路径 | 

参考：[example/protocg.yaml](example/protocg.yaml)

## 关键字

`protoc` `remote` `git` `version control`