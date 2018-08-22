/*
Security_MD5Hash

源代码在计算32位元组的时候，使用加法，那是不对的。。
*/
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	hash := md5.New()
	// := []byte("hello\n")
	//hash.Write(bytes)
	io.WriteString(hash, "hello\n")
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n]) << 24
		val |= uint32(hashValue[n+1]) << 16
		val |= uint32(hashValue[n+2]) << 8
		val |= uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}

	fmt.Println()
}
