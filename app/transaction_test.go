package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction-management/app/database"
)

var db = database.OpenDatabase()
var router = setupRouter()

func clearTable() {
	db.Exec("DELETE FROM transactions")
	db.Exec("DELETE FROM accounts")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/transactions", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "{\"error\":\"Transactions not found\"}" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
