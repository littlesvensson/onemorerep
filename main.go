package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/littlesvensson/onemorerep/internal/pr"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })
    pr.RegisterRoutes(r)

    log.Println("listening on :8080")
    log.Fatal(r.Run(":8080"))
}
