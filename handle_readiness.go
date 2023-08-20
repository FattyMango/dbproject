package main

import "net/http"

func handleReadiness(writer http.ResponseWriter, request *http.Request) {
	respondWithJson(writer, http.StatusOK, map[string]bool{"ok": true})
}
