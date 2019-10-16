package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"net/http"
)

func izposojaKnjigeHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("izposoja knjige to do"))
		fmt.Println("request:", r)
		fmt.Println("r.Body:", r.Body)

	})
}

func vracanjeKnjigeHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("vračanje knjige to do"))
		fmt.Println("request:", r)
		fmt.Println("r.Body:", r.Body)

	})
}
