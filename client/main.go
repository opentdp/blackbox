package main

import (
	"github.com/opentdp/blackbox/client/cmd/fork"
	"github.com/opentdp/blackbox/client/module/api"
)

func main() {
	fork.Checkin()
	api.Server()
}
