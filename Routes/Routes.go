package Routes

import (
	"net/http"

	"bitcoin-checker/Controller"

	"github.com/gin-gonic/gin"
)

func Init() (c *gin.Engine) {
	r := gin.Default()

	// checking the health of application
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Health check, OK",
		})
	})

	// grouping apis by group name api
	api := r.Group("/api")
	{
		api.GET("prices/btc", Controller.GetPrices)
	}
	r.Run()
	return
}
