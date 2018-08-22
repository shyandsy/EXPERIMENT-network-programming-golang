/*
ASN.1

本代码测试ASN1的marshal和unmarsshal方法。
分别对int，string，time进行marshal/unmarshal操作。
*/
package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"time"
)

func main() {
	test_int()
	test_string()
	test_time()
	test_time2()
}

func test_int() {
	a := 13
	fmt.Println("Input: ", a)
	mdata, err := asn1.Marshal(a)
	checkError(err)

	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", n)
	fmt.Println("=================================================")
}

func test_string() {
	a := "hello world"
	fmt.Println("Input: ", a)
	mdata, err := asn1.Marshal(a)
	checkError(err)

	var newstr string
	_, err1 := asn1.Unmarshal(mdata, &newstr)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newstr)
	fmt.Println("=================================================")
}

func test_time() {
	a := time.Now()
	fmt.Println("Input: ", a)
	mdata, err := asn1.Marshal(a)
	checkError(err)

	var newtime = new(time.Time)
	_, err1 := asn1.Unmarshal(mdata, newtime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime)
	fmt.Println("=================================================")
}

func test_time2() {
	a := time.Now()
	fmt.Println("Input: ", a)
	mdata, err := asn1.Marshal(a)
	checkError(err)

	var newtime time.Time                      // 定义变量，不创建instance
	_, err1 := asn1.Unmarshal(mdata, &newtime) // 需要 &
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime)
	fmt.Println("=================================================")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
