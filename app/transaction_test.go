package main

import (
	"encoding/json"
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

func TestGetNoExistedTransaction(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/transactions/32966FEF-283F-4F13-8BFC-638D0A099B22", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Transaction not found" {
		t.Errorf("Expected the value of 'error' key of the response is 'Transaction not found', got '%s'", m["error"])
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
