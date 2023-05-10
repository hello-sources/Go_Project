package main

import (
	"context"
	_case "goBasic/generic-T/case"
	"os"
	"os/signal"
)

func main() {
	_case.Simplecase()
	_case.CusNumTCase()
	_case.BuiltInCase()

	_case.TTypeCase()
	_case.TTypeCase1()

	_case.InterfaceCase()

	_case.ReceiverCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
