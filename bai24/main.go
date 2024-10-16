package main

import (
	"errors"
	"fmt"
)

// panic and error
func main() {
	if err := checkDbConnection(); err != nil {
		// fmt.Println("failed when sql connect")
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic error: %s\n ", r)
				fmt.Println("backup server is running...")
				fmt.Println("my working in going on backup server...")
			}
		}()
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
