package main

import (
	"time"
	"transaction-management/app/controller"
	"transaction-management/app/database"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

// func HTTPSRedirect(writer http.ResponseWriter,
// 	request *http.Request) {
// 	host := strings.Split(request.Host, ":")[0]
// 	target := "https://" + host + ":8080" + request.URL.Path
// 	if len(request.URL.RawQuery) > 0 {
// 		target += "?" + request.URL.RawQuery
// 	}
// 	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
// }

func main() {
	database.OpenDatabase()
	app := gin.Default()
	memoryStore := persist.NewMemoryStore(2 * time.Minute)

	{
		app.GET("/ping", controller.PingPong)
		app.GET("/transactions", cache.CacheByRequestURI(memoryStore, 1*time.Minute), controller.GetTransactions)
		app.POST("/transactions", controller.AddTransaction)
		app.GET("/transactions/:id", controller.GetSingleTransaction)
		app.GET("/accounts/:id", controller.GetSingleAccount)
	}
	app.Run()
}
