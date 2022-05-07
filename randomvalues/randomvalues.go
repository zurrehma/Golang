// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var wait sync.WaitGroup
var seconds int = 2

func reciveChannel(channels chan interface{}) {
	defer wait.Done()
	for {
		select {
		case message := <-channels:
			switch message.(type) {
			case string:
				fmt.Println("String", message)
			case int:
				fmt.Println("Integer", message)
			default:
				fmt.Println("default", reflect.TypeOf(message), message)
			}
		default:
			fmt.Println("waiting......")
		}
		//seconds := 10
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
func main() {
	mychan := make(chan interface{})
	wait.Add(1)
	go func() {
		for {
			switch rand.Int() % 4 {
			case 0:
				mychan <- rand.Int()
			case 1:
				mychan <- "Hello"
			case 3:
				mychan <- rand.Float32()
			default:
				mychan <- rand.Float64()

			}
			//seconds := 10
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}()

	go reciveChannel(mychan)
	wait.Wait()
}
