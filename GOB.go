/*
JSON
本代码旨在对JSON编码解码进行测试
*/
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}
type Name struct {
	Family   string
	Personal string
}
type Email struct {
	Kind    string
	Address string
}

// Person的toString方法
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	fileName := "person.gob"
	saveGob(fileName)

	var person Person
	loadGob(fileName, &person)

	fmt.Println("Person", person.String())
}

func saveGob(fileName string) {
	person := Person{
		Name: Name{Family: "Chu", Personal: "Tian Le"},
		Email: []Email{
			Email{Kind: "Work", Address: "sales@gmail.com"},
			Email{Kind: "Life", Address: "life@gmail.com"},
		},
	}

	// 创建文件
	outFile, err := os.Create(fileName)
	checkError(err)

	// 编码写入
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(person)
	checkError(err)

	// 关闭文件
	outFile.Close()
}

func loadGob(fileName string, key interface{}) {
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
