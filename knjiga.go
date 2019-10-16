package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func prikaziKnjigeHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("izposoja knjige to do"))
		fmt.Println("request:", r)
		fmt.Println("r.Body:", r.Body)

	})
}
