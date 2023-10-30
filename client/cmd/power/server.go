package power

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/opentdp/blackbox/client/cmd/args"
	"github.com/opentdp/go-helper/httpd"
	"github.com/opentdp/go-helper/psutil"
)

func metaServer() {
	engine := httpd.Engine(false)

	engine.GET("/tdp/ps", ps)
	engine.GET("/tdp/meta", meta)

	engine.NoRoute(forward())

	httpd.Server(args.MetaServerListen)
}

func ps(c *gin.Context) {
	c.AbortWithStatusJSON(200, psutil.Detail(true))
}

func meta(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"Version":       args.Version,
		"VersionNumber": args.VersionNumber,
		"NodeMetaInfo":  args.NodeMetaInfo,
	})
}

func forward() gin.HandlerFunc {
	backend, _ := url.Parse("http://127.0.0.1:9115")
	proxy := httputil.NewSingleHostReverseProxy(backend)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
