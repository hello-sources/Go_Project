package goroutine

import (
	"log"
	"time"
)

type Goroutine struct {
}

func (r *Goroutine) Name() string {
	return "Goroutine"
}

func (r *Goroutine) Run() {
	log.Println(r.Name(), "Run")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}
