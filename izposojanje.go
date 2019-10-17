package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Imena polj se začnejo z Veliko črko zaradi json pretvorbe
// Izposodi in Vrni morata biti različna,
// ker si lahko uporabnik le izposodi ali vrne knjigo
type izposojaSummary struct {
	UporabnikID string
	IzvodID     string
	Izposodi    bool
	Vrni        bool
}

// handlanje rpc klica
func rpcIzposoja(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch m := r.Method; m {
		case http.MethodPost:
			{
				azurirajIzposojo(db, w, r)
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

func azurirajIzposojo(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	// rezultat vrne klientu
	fmt.Fprintf(w, string("ažuriraj izposojo - to do"))

	// preberem json body v izp
	var izp izposojaSummary
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &izp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//preverim podatke
	// @todo
	if izp.Izposodi {
		fmt.Fprintf(w, string("izposoja - to do"))

		return
	}

	if izp.Vrni {
		fmt.Fprintf(w, string("vračanje - to do"))

		return
	}

	http.Error(w, "Napačni parametri - Izposodi, Vrni", http.StatusBadRequest)

	// rezultat vrne klientu
	//	fmt.Fprintf(w, string(out))

}
