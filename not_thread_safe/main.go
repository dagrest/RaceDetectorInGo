// How to run
// go run not_thread_safe/main.go
// go run -race not_thread_safe/main.gopackage main

package main

import (
	"fmt"
	"strconv"
)

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect(i int) *DriverPg {

	if instance == nil {
		// <--- NOT THREAD SAFE / when to use Goroutine
		instance = &DriverPg{conn: "DriverConnectPostgres" + ":" + strconv.Itoa(i)}
	}

	return instance
}

func main() {
	fmt.Println("Hello RaceDetectorInGo")

	go func() {
		for i := 0; i < 5; i++ {
			//time.Sleep(time.Millisecond * 600)
			fmt.Println(*Connect(i), " - ", i)
		}
	}()

	go func() {
		fmt.Println(*Connect(777))
	}()

	go func() {
		fmt.Println(*Connect(999))
	}()

	fmt.Scanln()
}
