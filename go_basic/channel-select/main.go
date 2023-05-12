package main

import (
	_case "goBasic/channel-select/case"
	"os"
	"os/signal"
)

func main() {
	_case.Communication()
	_case.ConcurrentSync()
	_case.NoticeAndMultiplexing()

	// 利用channel的阻塞机制
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
