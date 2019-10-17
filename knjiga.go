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

// za drug View
type knjigaCountSummary struct {
	Naziv            string
	CountIzposojenih int
}

type knjigeCount struct {
	KnjigeCount []knjigaCountSummary
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
	klist := knjigeCount{} // init
	err := queryKnjigeCount(db, &klist)
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
func queryKnjigeCount(db *sql.DB, klist *knjigeCount) error {
	rows, err := db.Query(`
		SELECT  k.naziv
  			,count(i.uporabnik_id)
       		FROM izvod i
		LEFT JOIN knjiga k  ON i.knjiga_id = k.id
		GROUP BY k.id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		knjiga := knjigaCountSummary{}
		err = rows.Scan(
			&knjiga.Naziv,
			&knjiga.CountIzposojenih,
		)
		if err != nil {
			return err
		}
		klist.KnjigeCount = append(klist.KnjigeCount, knjiga)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
