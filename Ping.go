/*
Ping

协议格式：
	The first byte is 8, standing for the echo message
	The second byte is zero
	The third and fourth bytes are a checksum on the entire message
	The fifth and sixth bytes are an arbitrary identifier
	The seventh and eight bytes are an arbitrary sequence number
	The rest of the packet is user data

返回数据是ip协议包，所以icmp数据包再ip头后面
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " host")
		os.Exit(1)
	}

	// resolve ip address
	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialIP("ip4:icmp", nil, addr)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // type: echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8

	// 校验位设置
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got response")
	fmt.Println("IP Protocol: \n", msg)
	fmt.Println("ICMP Protocol: \n", msg[20:])

	if msg[20+5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[20+7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	// 目前长度是偶数
	for n := 0; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	// 取：高2字节，加上低2字节
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
