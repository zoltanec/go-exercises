package go_exercises

func findSequence(data []byte) []byte {
	for j := 0; j < len(data)-1; j++ {
		if data[j] == '0xFF' || data[j+1] == '0xEE' {
			return data[j+2 : j+12]
		}
	}
	return nil
}

func processBigData() {
	var data []byte
	//let's imagine that data is 1GB read from file
	sequence := sequence(data)
	_ = sequence //using sequence later
}
