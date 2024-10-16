package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type student struct {
	Name   string `json:"name"`
	Course string `json:"course"`
}

// timer
func main() {
	db := []student{
		{"nhutran", "b-TECH"},
		{"LETRANG", "pythongo"},
		{"SanjiKu", "C#C++"},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(db)
	file, err := os.Create("student.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, buf)

	fmt.Println("--------------")
	dbnew := []student{}
	filenew, errnew := os.Open("student.json")
	if errnew != nil {
		panic(errnew)
	}
	defer filenew.Close()

	json.NewDecoder(filenew).Decode(&dbnew)
	fmt.Println("dbnew = ", dbnew)

}
