package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerHello(c *gin.Context) {
	c.String(http.StatusOK, "The server is working!! :)")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Hello World how is everything Andres maybe is complete now")
}
