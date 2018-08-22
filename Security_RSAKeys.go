/*
本程序旨在演示如何生成，保存，加载RSA key
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// 生成RSA key,并保存文件
	generateKey()

	fmt.Println("============================")
	loadKey()
}
func generateKey() {
	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes ", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent ", key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modules ", publicKey.N.String())
	fmt.Println("public key exponent ", publicKey.E)

	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)
	savePEMKey("private.pem", key)
}

func saveGobKey(fileName string, key interface{}) {
	// 打开文件
	outFile, err := os.Create(fileName)
	checkError(err)

	// 写入文件
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

	// 关闭文件
	outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	pem.Encode(outFile, privateKey)

	outFile.Close()
}

func loadKey() {
	var key rsa.PrivateKey

	load("private.key", &key)
	fmt.Println("Private key primes ", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent ", key.D.String())

	var publicKey rsa.PublicKey
	load("public.key", &publicKey)
	fmt.Println("Private key modules ", publicKey.N.String())
	fmt.Println("Private key exponent ", publicKey.E)
}

func load(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)

	inFile.Close()
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
