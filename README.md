# Izposoja

Aplikacija v golang-u, ki omogoča nekaj API endpointov za namišljeno aplikacijo izposoje knjig:
- vnos uporabnika
- seznam uporabnikov
- izposoja knjige
- vračanje knjige
- seznam knjig s številom prostih izvodov

## Priprava okolja

Kreiranje baze, npr.

`createdb -l sl_SI.UTF-8 -E UTF8 -T template0 izposoja`

Postavitev inicialne baze z nekaj testnimi podatki:

`psql -d izposoja -a -f sql/napolniDb.sql`

