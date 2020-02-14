# music

## 编译

```shell
export GOPATH=/Users/wenbao/code/go/study/music

go build smp/mlib
go test smp/mlib
go build smp/mp
go test smp/mp

go install smp/mlib
go install smp/mp

go build player
go install player
```

## 运行

```shell
cd bin

./player
```