package api

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/psutil"

	"github.com/opentdp/blackbox/client/cmd/args"
)

func ps(c *gin.Context) {
	c.AbortWithStatusJSON(200, psutil.Detail(true))
}

func meta(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"Metadata":      args.Metadata,
		"Version":       args.Version,
		"VersionNumber": args.VersionNumber,
	})
}

func forward() gin.HandlerFunc {
	backend, _ := url.Parse("http://127.0.0.1:9115")
	proxy := httputil.NewSingleHostReverseProxy(backend)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
