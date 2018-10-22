package main

import (
	"fmt"
	"os"
	"github.com/shiyunjin/go-telnet-cisco"
)

func main() {
	client := new(telnet.Client)

	err := client.Connect("192.168.63.250:23")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	err = client.Login("YOU USERNAME","YOU PASS","YOU ENA PASS")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}


	text,err := client.Cmd("show port")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	fmt.Println(text)

	//text,err = client.Cmd("show access-lists")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	//	return
	//}
	//
	//fmt.Println(text)

	text,err = client.Cmd("show interface")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	fmt.Println(text)
}