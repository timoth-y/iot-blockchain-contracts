package main

import (
	"syscall"

	"github.com/timoth-y/chainmetric-contracts/shared"
	"github.com/ztrue/shutdown"
)

func init() {
	shared.InitCore()
}

func main() {
	go shared.BootstrapContract(NewReadingsContract())

	shutdown.Add(shared.CloseCore)
	shutdown.Listen(syscall.SIGINT, syscall.SIGTERM)
}
