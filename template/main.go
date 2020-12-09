package main

func main() {
	recordGenerator := constructRecordGenerator("input", "\n")
	for {
		record, done := recordGenerator()
		if done {
			break
		}
	}
}
