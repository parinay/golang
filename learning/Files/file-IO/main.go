package main

func main() {
	// read file in memory
	readFileInMemory()
	// read file path from cmd line
	readCmdLine()
	// read in small chunks
	readSmall()
	// read lines
	readLine()

	// write string
	writeString()

	// write bytes
	writeBytes()

	//writeLine
	writeLine()

	//append file
	appendFile()

	//concurrent write
	concurrent()
}
