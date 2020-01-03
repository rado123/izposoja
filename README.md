# Izposoja

Aplikacija v golang-u, ki omogoča nekaj API endpointov za namišljeno aplikacijo izposoje knjig:
- dodajanje uporabnika
- seznam uporabnikov
- izposoja knjige
- vračanje knjige
- seznam knjig s številom prostih izvodov

## Priprava okolja

Primeri za Linux OS

Kreiranje baze, npr.

`createdb -l sl_SI.UTF-8 -E UTF8 -T template0 izposoja`

Postavitev inicialne baze z nekaj testnimi podatki:

`psql -d izposoja -a -f sql/napolniDb.sql`

Postavimo se v vrhnji direktorij aplikacije.
Priprava parametrov konekcije v bazo:

`cp config/pgConn.sh.dist config/pgConn.sh`

in popravimo geslo in ostale konekcijske parametre v datoteki config/pgConn.sh

## Zagon strežniške aplikacije

Po potrebi se koda prevede z

`go install`

zažene pa s

`source config/pgConn.sh; izposoja`

### Primeri uporabe 

Za CRUD operacije se sledi de facto stanardu za REST:
- GET metoda za branja
- PUT metoda za popravek zapisov
- POST metoda za nove zapise

Za ne CRUD operacije se uporabljajo RPC-ji s POST metodo.

V obeh primerih (rpc in rest) se uporabljajo JSON datoteke tako za vhod in izhod

#### Dodajanje uporabnika

Uporabljen je REST s POST metodo in JSON vhodno in izhodno datoteko, npr:



	curl -d '{"Ime":"Monika","Priimek": "Žagar"}' -H "Content-pe: application/json" -X POST http://localhost:8000/rest/uporabnik

in dobimo rezultat v json obliki:
    
	{"ID":"91495f17-99c4-4c98-7342-d76e7d95dbf5"
    ,"Ime":"Monika"
    ,"Priimek":"Žagar"}
    
#### Pregled uporabnikov

Uporabljen je REST z GET metodo na istem URL-ju kot dodajanje uporabnika

	curl  -X GET http://localhost:8000/rest/uporabnik
    
in dobimo seznam uporabnikov v json obliki:

	{"Uporabniki":[
    	{"ID":"fa906169-e296-4a4d-aaa2-72a63e37278d","Ime":"Janko","Priimek":"Bezjak"},
    	{"ID":"88144d9b-9db4-4d6a-89f7-713d27e2a2a2","Ime":"Miha","Priimek":"Novak"},
   	 	{"ID":"f95260c9-921e-46bd-80c6-3ded23437db0","Ime":"Anica","Priimek":"Veber"},
    	{"ID":"91495f17-99c4-4c98-7342-d76e7d95dbf5","Ime":"Monika","Priimek":"Žagar"}
    ]}
    
#### Izposoja knjige

Uporabljen je RPC s POST metodo in vhodno json datoteko. Za izposojo uporabimo polje "Izposodi" :true

	curl -d '{  "UporabnikID":"88144d9b-9db4-4d6a-89f7-713d27e2a2a2", "IzvodID":"3cf1d23e-faeb-4dd9-ba08-e69b6b395269", "Izposodi" :true}' -H "Content-pe: application/json" -X POST  http://localhost:8000/rpc/izposoja

#### Vračanje knjige

Podobno kot izposoja knjige, le da uporabimo polje  "Vrni" :true

	curl -d '{  "UporabnikID":"88144d9b-9db4-4d6a-89f7-713d27e2a2a2", "IzvodID":"3cf1d23e-faeb-4dd9-ba08-e69b6b395269", "Vrni":true}' -H "Content-pe: application/json" -X POST  http://localhost:8000/rpc/izposoja

#### Seznam knjig s številom prostih izvodov

Uporabljen je REST z GET metodo.

	curl  -X GET http://localhost:8000/rest/knjiga

in dobimo rezultat v json obliki:

	{"KnjigeCount":[
    	{"Naziv":"Pod snegom","Prostih":3},
    	{"Naziv":"Metulj","Prostih":1},
    	{"Naziv":"Vojna in mir","Prostih":1}
    ]}



