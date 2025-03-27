package go_exercises

func action() {}

func changeTwoStrings() {
	var str = "go"
	newStr := str + "-go"

	strData = unsage.StringData(str)
	newStrData = unsafe.StringData(newStr)

	printLn("action: ", action)
	printLn("strData: ", strData)
	printLn("newStr: ", newStrData)

	slice := usafe.Slice(strData, len(strData))
	newSlice := unsafe.Slice(newStrData, len(newStrData))

	newSlice[0] = 'G'
	printLn("newstr: ", newStrData)

	slice[0] = 'G'
	printLn("str: ", strData)
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
