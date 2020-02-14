# cgss

## 编译

```shell
export GOPATH=/Users/wenbao/code/go/study/cgss

go install cgss
```

## 运行

```shell
cd bin

./cgss
```

command 测试
```shell
login Tom 1 101
login Jerry 2 321
listplayer
send Hello everybody
logout Tom
listplayer
send Hello the people online
logout Jerry
listplayer
send anybody?
q
```