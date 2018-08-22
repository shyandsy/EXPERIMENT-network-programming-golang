/*
Security_GenX509Cert

PKI(public key infrastructure)是一个集成公钥和附属信息，比如所有人名字，位置，以及名字和位置的关联。
现在使用的最主要的PKI是基于x509证书的,比如浏览器使用它来验证网站身份

本程序用于为网站生成一个自签名的x509证书，并保存为.cer文件
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	generateX509Cert()

	loadX509Cert()
}

func loadX509Cert() {
	certCerFile, err := os.Open("shyandsy.company.cer")
	checkError(err)

	derBytes := make([]byte, 500)
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()

	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())
}
func generateX509Cert() {
	random := rand.Reader

	// 加载private key
	var key rsa.PrivateKey
	loadKey("private.key", &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) // 一年
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "shyandsy.company",
			Organization: []string{"shyandsy"},
		},
		NotBefore:             now,
		NotAfter:              then,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:     true,
		DNSNames: []string{"shyandsy.company", "localhost"},
	}
	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	checkError(err)

	// 写入.cert文件
	certCerFile, err := os.Create("shyandsy.company.cer")
	checkError(err)
	certCerFile.Write(derBytes)
	certCerFile.Close()

	// 写入.pem文件
	certPEMFile, err := os.Create("shyandsy.company.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	keyPEMFile.Close()
}

func loadKey(fileName string, key interface{}) {
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
