package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func transactionsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
	case http.MethodPost:
	default:
		fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported")
	}

}
func singleTransactionsHandler(writer http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/transaction/")
	
}
func getAllTransactionsHandler(writer http.ResponseWriter, request *http.Request) {

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am root"))
}

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
	http.HandleFunc("/ping", pingPongHandler)
	http.HandleFunc("/transaction", transactionsHandler)
	http.HandleFunc("/transaction/{id}", singleTransactionsHandler)
	http.HandleFunc("/", rootHandler)
}

func main() {
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
