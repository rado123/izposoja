package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Imena polj se začnejo z Veliko črko zaradi json pretvorbe
type uporabnikSummary struct {
	ID      int
	Ime     string
	Priimek string
}

type uporabniki struct {
	Uporabniki []uporabnikSummary
}

func dodajUporabnikaHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("dodaj uporanika to do"))
		fmt.Println("request:", r)
		fmt.Println("r.Body:", r.Body)

	})
}

func prikaziUporabnikeHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ulist := uporabniki{} // init

		err := queryUporabniki(db, &ulist)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// pretvorim v json obliko
		out, err := json.Marshal(ulist)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// rezultat vrne klientu
		fmt.Fprintf(w, string(out))

	})
}

// fetcha uporabnike iz db
func queryUporabniki(db *sql.DB, ulist *uporabniki) error {

	rows, err := db.Query(`
		SELECT
			ID,
			Ime,
			Priimek
		FROM uporabnik
		ORDER BY Priimek`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		uporabnik := uporabnikSummary{}
		err = rows.Scan(
			&uporabnik.ID,
			&uporabnik.Ime,
			&uporabnik.Priimek,
		)
		if err != nil {
			return err
		}
		ulist.Uporabniki = append(ulist.Uporabniki, uporabnik)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
