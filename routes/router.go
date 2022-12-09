package routes

import (
	"github.com/gin-gonic/gin"
	"goblog/config"
	"net/http"
)

func InitRouter() {
	gin.SetMode(config.GlobalConfig.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSONP(http.StatusOK, gin.H{
				"mgs": "ok",
			})
		})

	}
	err := r.Run(string((config.GlobalConfig.HttpPort)))
	if err != nil {
		return
	}
}
