package main

import (
	"database/sql"
	"encoding/json"
	//"fmt"
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
type izvodSummary struct {
	ID           string
	Uporabnik_id string
	Knjiga_id    string
	Signatura    string
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

		//preverim, če še ni izposojena
		var iout izvodSummary

		row := db.QueryRow(`SELECT
			ID,
			Uporabnik_id
		FROM izvod
		WHERE ID=$1`, izp.IzvodID)
		row.Scan(
			&iout.ID,
			&iout.Uporabnik_id,
		)
		if iout.Uporabnik_id != "" {
			http.Error(w, "Izvod že izposojen", http.StatusBadRequest)
			return
		}

		// zapišem u v bazo
		ukaz := `UPDATE izvod
			SET 	uporabnik_id=$1
			WHERE	id=$2`
		_, err = db.Exec(ukaz, izp.UporabnikID, izp.IzvodID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if izp.Vrni {
		//preverim, če še ni izposojena
		var iout izvodSummary

		row := db.QueryRow(`SELECT
			ID,
			Uporabnik_id
		FROM izvod
		WHERE ID=$1`, izp.IzvodID)
		row.Scan(
			&iout.ID,
			&iout.Uporabnik_id,
		)
		if iout.Uporabnik_id != izp.UporabnikID {
			http.Error(w, "Uporabnik nima izposojenega tega izvoda", http.StatusBadRequest)
			return
		}

		//popravim v bazi
		ukaz := `UPDATE izvod
			SET 	uporabnik_id=null
			WHERE	id=$1`
		_, err = db.Exec(ukaz, izp.IzvodID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Napačni parametri - Izposodi, Vrni", http.StatusBadRequest)

	// rezultat vrne klientu
	//	fmt.Fprintf(w, string(out))

}
