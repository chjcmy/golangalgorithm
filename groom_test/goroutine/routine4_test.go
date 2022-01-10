package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func TestRoutine4(t *testing.T) {
	tirech := make(chan *Car)
	paintch := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tirech)
	go InstallTire(tirech, paintch)
	go PaintCar(paintch)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}

	}
}

func InstallTire(tirech, paintch chan *Car) {
	for car := range tirech {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintch <- car
	}
	wg.Done()
	close(paintch)
}

func PaintCar(paintch chan *Car) {
	for car := range paintch {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.4f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Body)
	}
	wg.Done()
}
