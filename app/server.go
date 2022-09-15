package main

import (
	"transaction-management/app/controller"
	"transaction-management/app/database"

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

func init() {
	database.OpenDatabase()
	r := gin.Default()
	{
		r.GET("/ping", controller.PingPong)
		r.GET("/transactions", controller.GetTransactions)
		r.POST("/transactions", controller.AddTransaction)
		r.GET("/transactions/:id", controller.GetSingleTransaction)
		r.GET("/accounts/:id", controller.GetSingleAccount)
	}
	r.Run()
}

func main() {

}
