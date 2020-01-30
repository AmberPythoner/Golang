package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(context *gin.Context) {
		t := time.Now()
		context.Next()
		log.Info("Request:",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("spend", time.Now().Sub(t)))
	}, func(context *gin.Context) {
		context.Set("requestId", rand.Int())
		context.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		res := gin.H{
			"message": "pong",
		}
		if value, ok := c.Get("requestId"); ok {
			res["requestId"] = value
		}
		c.JSON(200, res)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
