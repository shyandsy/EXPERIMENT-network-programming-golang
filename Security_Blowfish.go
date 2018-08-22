/*
Security_Blowfish

go支持两只对称加密算法Blowfish和DES

需要执行
	golang.org/x/crypto/blowfish
*/
package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

func main() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	src := []byte("hello world\n\n\n")

	// 8字节分组加密
	var enc [512]byte
	for i := 0; i < len(src); i += 8 {
		cipher.Encrypt(enc[i:i+8], src[i:i+8])
	}

	// 8字节分组解密
	var decrypt [512]byte
	for i := 0; i < len(enc); i += 8 {
		cipher.Decrypt(decrypt[i:i+8], enc[i:i+8])
	}

	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:len(src)])
	fmt.Println(string(result.Bytes()))
}
