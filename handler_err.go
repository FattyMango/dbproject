
package main

import "net/http"

func handleErr(writer http.ResponseWriter, request *http.Request) {
	respondWithError(writer, http.StatusBadRequest, "Something went wrong")
}