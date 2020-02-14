# sorter

## 编译

```shell
export GOPATH=/Users/wenbao/code/go/study/sorter

go build sorter/algorithms/qsort
go test sorter/algorithms/qsort
go build sorter/algorithms/bubblesort
go test sorter/algorithms/bubblesort

go install sorter/algorithms/qsort
go install sorter/algorithms/bubblesort

go build sorter
go install sorter
```

## 运行

```shell
cd bin

./sorter -i unsorted.dat -o qsorted.dat -a qsort
./sorter -i unsorted.dat -o bsorted.dat -a bubblesort
```