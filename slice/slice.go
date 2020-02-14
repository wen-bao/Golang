package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlient:")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	mySlice2 := make([]int, 5, 10)

	for i, v := range mySlice2 {
		fmt.Println("mySlice2[", i, "] = ", v)
	}

	fmt.Println("len(mySlice2):", len(mySlice2)) //5
	fmt.Println("cap(mySlice2):", cap(mySlice2)) //10

}
