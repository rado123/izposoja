package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Imena polj se začnejo z Veliko črko zaradi json pretvorbe
type knjigaSummary struct {
	ID    string
	Naziv string
}

type knjige struct {
	Knjige []knjigaSummary
}

func restKnjiga(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch m := r.Method; m {
		case http.MethodGet:
			{
				prikaziKnjige(db, w, r)
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

func prikaziKnjige(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	klist := knjige{} // init
	err := queryKnjige(db, &klist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// pretvorim v json obliko
	out, err := json.Marshal(klist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// rezultat vrne klientu
	fmt.Fprintf(w, string(out))

}

// fetcha knjige iz db
func queryKnjige(db *sql.DB, klist *knjige) error {
	rows, err := db.Query(`
		SELECT
			ID,
			Naziv
		FROM knjiga
		ORDER BY Naziv`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		knjiga := knjigaSummary{}
		err = rows.Scan(
			&knjiga.ID,
			&knjiga.Naziv,
		)
		if err != nil {
			return err
		}
		klist.Knjige = append(klist.Knjige, knjiga)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
