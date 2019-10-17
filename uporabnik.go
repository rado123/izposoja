package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
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

func restUporabnik(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch m := r.Method; m {
		case http.MethodPost:
			{
				dodajUporabnika(db, w, r)
			}
		case http.MethodGet:
			{
				prikaziUporabnike(db, w, r)
			}
		default:
			{
				e := "Metoda" + m + " ni implementirana"
				http.Error(w, e, http.StatusMethodNotAllowed)
			}
		}
		return
	})
}

func dodajUporabnika(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("metoda JE POST"))
	fmt.Fprintf(w, string("dodaj uporanika to do"))
	fmt.Println("request:", r)
	fmt.Println("request.Method:", r.Method)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("body=", string(body))

	var t uporabnikSummary
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("t.,ime=", t.Ime)
	fmt.Println("t=", t)
}

func prikaziUporabnike(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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
