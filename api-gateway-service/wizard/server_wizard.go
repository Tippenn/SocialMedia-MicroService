package wizard

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Tippenn/Social-Media-Project/shared/middleware"
	"github.com/gin-gonic/gin"
)

func proxyTo(target string) gin.HandlerFunc {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatalf("bad service url %q: %v", target, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error":"service unavailable"}`))
	}
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterServer(router *gin.Engine) {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// public: nobody has a token yet
	r.Any("/v1/auth/*path", proxyTo(Config.AccountServer.BaseURL))

	// protected
	api := r.Group("/v1", middleware.AuthMiddleware(SessionUsecase))
	api.Any("/users/*path", proxyTo(Config.AccountServer.BaseURL))
	api.Any("/posts/*path", proxyTo(Config.PostServer.BaseURL))
	api.Any("/timeline/*path", proxyTo(Config.TimelineServer.BaseURL))
}
