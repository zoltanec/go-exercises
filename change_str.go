package main

import (
	"fmt"
	"unsafe"
)

func action() {}

func changeTwoStrings() {
	var str = "go"
	newStr := str + "-go"

	strData = unsafe.StringData(str)
	newStrData = unsafe.StringData(newStr)

	fmt.println("action: ", action)
	fmt.println("strData: ", strData)
	fmt.println("newStr: ", newStrData)

	slice := unsafe.Slice(strData, len(strData))
	newSlice := unsafe.Slice(newStrData, len(newStrData))

	newSlice[0] = 'G'
	fmt.println("newstr: ", newStrData)

	slice[0] = 'G'
	fmt.println("str: ", strData)
}

//action:  0x1001596f0
//strData: 0x100126e90
//newStr:  0x14000056708
//newstr: Go-go
//unexpected fault address
//bus error

func changeString() {
	data := []byte("Hello world")
	strData := unsafe.String(unsafe.SliceData(data), len(data))

	fmt.Println(strData)
	data[0] = 'W'
	fmt.Println(strData)
}
