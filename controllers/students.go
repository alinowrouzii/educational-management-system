package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type testStruct struct {
	Test string `json:"test"`
}

func TestHandler(w http.ResponseWriter, _ *http.Request) {

	respondWithJSON(w, http.StatusOK, testStruct{
		Test: "hello world",
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload testStruct) {
	fmt.Println(payload)
	response, _ := json.Marshal(payload)
	// fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
