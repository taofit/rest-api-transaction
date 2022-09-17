package main

import (
	"time"
	"transaction-management/app/controller"
	"transaction-management/app/database"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	app := gin.Default()
	memoryStore := persist.NewMemoryStore(2 * time.Minute)
	cacheHandler := cache.CacheByRequestURI(memoryStore, 1*time.Minute)

	{
		app.GET("/ping", controller.PingPong)
		app.GET("/transactions", cacheHandler, controller.GetTransactions)
		app.POST("/transactions", controller.AddTransaction)
		app.GET("/transactions/:id", controller.GetSingleTransaction)
		app.GET("/accounts/:id", controller.GetSingleAccount)
		app.HandleMethodNotAllowed = true
	}
	return app
}

func main() {
	database.OpenDatabase()
	setupRouter().Run()
}
