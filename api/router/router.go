package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"rdr/api/handler"
	"rdr/app/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	hs          *handler.Handlers
	VersionInfo config.VersionInfo
}

func GinAuthMW(c *gin.Context) {
	if u, p, ok := c.Request.BasicAuth(); !ok || !(u == "admin" && p == "AdminPass") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unautorized"))
		return
	}
	c.Next()
}

func NewRouter(hs *handler.Handlers) *Router {
	r := gin.Default()
	ret := &Router{
		hs: hs,
		VersionInfo: config.VersionInfo{
			Version: config.Version,
			Commit:  config.Commit,
			Build:   config.Build,
		},
	}

	r.GET("/", ret.Goto)

	// Admin
	adm := r.Group("/admin", GinAuthMW)
	adm.GET("/", ret.Admin)

	ret.Engine = r
	return ret
}

func (rt *Router) Goto(c *gin.Context) {
	u := c.Query("url")
	ur, err := url.Parse(u)
	if err != nil {
		log.Println("Error parse url", err)
	}

	// if ur.Scheme == "" {
	// 	ur.Scheme = "http"
	// }

	// jj, _ := json.MarshalIndent(ur, "", " ")
	// log.Println("Host: ", ur.Host, string(jj))

	if ur.Host == "" {
		// c.JSON(http.StatusNotFound, gin.H{"error": true, "message": "Page not found"})
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("Page not found"))
	}

	log.Println(c.ClientIP(), " => ", ur.String())

	hdrs, _ := json.MarshalIndent(c.Request.Header, "", " ")

	log.Println(string(hdrs))
	c.Redirect(http.StatusFound, ur.String())
}

func (rt *Router) Admin(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"status": "ok", "version": rt.VersionInfo})
}
