package power

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/go-helper/httpd"
)

func metaServer() {
	engine := httpd.Engine(false)

	engine.GET("/tdp/info", info)
	engine.NoRoute(forward())

	httpd.Server(args.MetaServerListen)
}

func forward() gin.HandlerFunc {
	backend, _ := url.Parse("http://127.0.0.1:9115")
	proxy := httputil.NewSingleHostReverseProxy(backend)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func info(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"metaInfo": args.NodeMetaInfo,
		"version":  args.Version,
	})
}
