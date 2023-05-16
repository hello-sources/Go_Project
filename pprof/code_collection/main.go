package main

import (
	"log"
	"os"
	"pprof/data"
	"pprof/data/block"
	"pprof/data/cpu"
	"pprof/data/goroutine"
	"pprof/data/mem"
	"pprof/data/mutex"
	"runtime/pprof"
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
	cpufile, err := os.OpenFile("code_collection/out/cpu.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	memfile, err := os.OpenFile("code_collection/out/mem.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.WriteHeapProfile(memfile)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		for _, v := range cmds {
			v.Run()
		}
		time.Sleep(time.Second)
	}
}
