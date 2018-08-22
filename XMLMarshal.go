/*
XMLMarshal.go

本程序将将Person结构转换成xml字符串，再把xml字符串转换成go结构体并输出
*/
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName Name    `xml:"person"`
	Name    Name    `xml:"name"`
	Email   []Email `xml:"email"`
}

type Name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {
	p := Person{
		Name: Name{
			Family:   "Chu",
			Personal: "Tianle",
		},
		Email: []Email{
			Email{
				Type:    "life",
				Address: "life@gmail.com",
			},
			Email{
				Type:    "work",
				Address: "work@gmail.com",
			},
		},
	}

	// marshal操作
	bytes, err := xml.Marshal(p)
	checkError(err)

	str := string(bytes)

	fmt.Println("The XML string is: \n" + str)
	/*
		str := `<?xml version="1.0" encoding="utf-8"?>
			<person>
				<name>
					<family>Newmarch</family>
					<personal>Jan</personal>
				</name>
				<email type="personal">personal@gmail.com</email>
				<email type="work">work@gmail.com</email>
			</person>`
	*/
	var person Person
	err = xml.Unmarshal([]byte(str), &person)
	checkError(err)

	fmt.Println("Family name: \"" + person.Name.Family + "\"")
	fmt.Println("Second email address: \"" + person.Email[1].Address + "\"")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
