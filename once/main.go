// How to run
// go run once/main.go
// go run -race once/main.go

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// call only once
var once sync.Once

var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect(i int) *DriverPg {

	once.Do(func() {
		instance = &DriverPg{conn: "DriverConnectPostgres" + ":" + strconv.Itoa(i)}
	})

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
		time.Sleep(time.Millisecond * 200)
		fmt.Println(*Connect(777))
	}()

	go func() {
		fmt.Println(*Connect(999))
	}()

	fmt.Scanln()
}
