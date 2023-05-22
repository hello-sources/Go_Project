package main

import (
	"os"
	"os/signal"
	many_many "producer-consumer/many-many"
	"producer-consumer/out"
	"syscall"
)

func main() {
	o := out.NewOut()
	go o.OutPut()

	//out.Println("AAAAAA")
	//out.Println("BBBBBB")
	//out.Println("CCCCCC")
	//out.Println("DDDDDD")
	//out.Println("EEEEEE")

	many_many.Exec()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
