package main

import (
	"errors"
	"fmt"
)

// panic and error
func main() {
	if err := checkDbConnection(); err != nil {
		// fmt.Println("failed when sql connect")
		panic("failed when mysql connect in panic")
	} else {
		fmt.Println("my sql connect success fully")
	}
	fmt.Println("co chay duoi nay khong")

}

func checkMysql() bool {
	return false
}

func checkDbConnection() error {
	if checkMysql() {
		return nil
	} else {
		return errors.New("MY SQL CAN NOT CONNECT .....")
	}

}
