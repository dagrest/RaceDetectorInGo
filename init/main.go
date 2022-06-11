// How to run
// go run init/main.go
// go run -race init/main.go

package main

import (
	"fmt"
	"time"
)

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {
	instance = &DriverPg{conn: "DriverConnectPostgres"}
	return instance
}

func init() {
	Connect()
}

func main() {
	fmt.Println("Hello RaceDetectorInGo")

	go func() {
		for i := 0; i < 5; i++ {
			//time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(), " - ", i)
		}
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		fmt.Println(*Connect())
	}()

	go func() {
		fmt.Println(*Connect())
	}()

	fmt.Scanln()
}
