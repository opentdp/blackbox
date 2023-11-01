package api

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/blackbox/client/cmd/args"
)

func Server() {
	engine := httpd.Engine(false)

	engine.GET("/api/ps", ps)
	engine.GET("/api/meta", meta)

	engine.NoRoute(forward())
	httpd.Server(args.MetaServerListen)
}
