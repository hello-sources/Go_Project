package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"pprof/data"
	"pprof/data/block"
	"pprof/data/cpu"
	"pprof/data/goroutine"
	"pprof/data/mem"
	"pprof/data/mutex"
	"runtime"
	"time"
)

var cmds = []data.Cmd{
	&cpu.Cpu{},
	&mem.Mem{},
	&block.Block{},
	&goroutine.Goroutine{},
	&mutex.Mutex{},
}

func main() {
	//开启对阻塞操作的追踪
	runtime.SetBlockProfileRate(1)
	// 开启对锁调用的追踪
	runtime.SetMutexProfileFraction(1)

	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range cmds {
			v.Run()
		}
		time.Sleep(time.Second)
	}
}
