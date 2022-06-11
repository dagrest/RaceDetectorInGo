// How to run
// go run mutex/main.go
// go run -race mutex/main.go

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect(i int) *DriverPg {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
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
		time.Sleep(time.Millisecond * 200)
		fmt.Println(*Connect(777))
	}()

	go func() {
		fmt.Println(*Connect(999))
	}()

	fmt.Scanln()
}
