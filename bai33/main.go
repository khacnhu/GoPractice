package main

import (
	"fmt"
	"os"
)

// timer
func main() {
	createFile()
	openFile()
	openFile2()
}

func createFile() {
	file, err := os.Create("info.txt")

	if err != nil {
		panic(err)
	}

	data := []byte("this is my information for the text file")
	file.Write(data)
	file.WriteString("\nthis is another way to write information in text file")
}

func openFile() {
	file, err := os.Open("info.txt")
	if err != nil {
		panic(err)
	}
	data := make([]byte, 100)

	file.Read(data)
	fmt.Println(string(data))

}

func openFile2() {
	data, err := os.ReadFile("info.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}
