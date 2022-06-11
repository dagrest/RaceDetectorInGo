// How to run
// go run atomic/main.go
// go run -race atomic/main.go

package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var atomicint uint64

var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect(i int) *DriverPg {

	// making sure you're already logged in
	if atomic.LoadUint64(&atomicint) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	// login only once
	if atomicint == 0 {
		instance = &DriverPg{conn: "DriverConnectPostgres" + ":" + strconv.Itoa(i)}
		atomic.StoreUint64(&atomicint, 1)
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
