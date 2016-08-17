package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	port := envString("PORT", "80")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/test", handleTest)

	fmt.Printf("Server starting...\n")
	endless.ListenAndServe(":"+port, router)
	fmt.Printf("Server finished.\n")
}

func handleTest(c *gin.Context) {
	if time.Now().Second()%2 == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusText(http.StatusOK),
			"info":   "Current second is even!",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"info":   "Current second is odd!",
	})
}

func envString(env string, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
