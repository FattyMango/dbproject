package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithError(Writer http.ResponseWriter, code int, message string) {
	
	if code>=500 {
		log.Println(message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(Writer, code, errorResponse{Error: message})
}


func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}