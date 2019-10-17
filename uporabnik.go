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
	ID      string
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
	fmt.Println("request:", r)
	fmt.Println("request.Method:", r.Method)

	// preberem json body v u
	var u uporabnikSummary
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("body=", string(body))
	fmt.Println("...1")

	err = json.Unmarshal(body, &u)
	fmt.Println("...2")
	if err != nil {
		fmt.Println("...2.5")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("...3")
	fmt.Println("u.ime=", u.Ime)
	fmt.Println("u=", u)
	//generiram guid
	id, err := generateGuid()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("uuid=", stringGuid(id))
	u.ID = stringGuid(id)

	// zapišem u v bazo
	ukaz := "INSERT INTO uporabnik (ID,Ime,Priimek) VALUES ($1, $2, $3);"
	_, err = db.Exec(ukaz, u.ID, u.Ime, u.Priimek)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//preberem iz baze

}

func prikaziUporabnike(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ulist := uporabniki{} // init84

	err := queryUporabniki(db, &ulist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
