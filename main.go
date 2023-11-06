package main

import "fmt"

func main() {
	var name string
	fmt.Println("你好，你叫甚麼名字？")
	fmt.Scanln(&name)
	if name != "" {
		fmt.Println("你好，" + name + "！")
	} else {
		fmt.Println("你好，世界！")
	}
}
