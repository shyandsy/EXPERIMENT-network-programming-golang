/*
本代码演示了make分配内存自动扩容的用法
*/
package golang

import "fmt"

func main() {
	array := make([]int, 0, 10)

	// 越界错误： index out of range
	//array[0] = 1

	// ok 可以访问
	array = array[0:1]
	array[0] = 1

	// 越界错误: index out of range
	//array[1] = 2

	// 错误:  slice bounds out of range
	//array = array[0:11]
	//array[10] = 10

	// ok 不会越界 一切正常
	for i := 0; i < 100; i++ {
		array = append(array, i)
	}

	for i := 0; i < len(array); i += 1 {
		fmt.Println(i)
	}
}
