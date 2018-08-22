package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	test_string_len()

	utf8_to_utf16()
}

func test_string_len() {
	str := "请先google下"
	fmt.Println("字符长度", len([]rune(str))) // 9个字符
	fmt.Println("字节长度", len([]byte(str))) // 15字节，一个中文3字节，一个英文1字节

	fmt.Println("")

	str = "请先谷歌下"
	fmt.Println("字符长度", len([]rune(str))) // 5个字符
	fmt.Println("字节长度", len([]byte(str))) // 15字节，一个中文3字节，一个英文1字节

	fmt.Println("================================")
}

func utf8_to_utf16() {
	str := "google一下，你就知道了"

	renes := utf16.Encode([]rune(str))
	ints := utf16.Decode(renes)

	str = string(ints)
	fmt.Println("解码：", str)
}
