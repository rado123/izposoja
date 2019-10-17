-- createdb -l sl_SI.UTF-8 -E UTF8 -T template0 izposoja

drop table if exists uporabnik ;
CREATE TABLE uporabnik (
   id uuid not null PRIMARY KEY,
   ime      varchar(255), 
   priimek  varchar(255)
);

drop table if exists knjiga ;
CREATE TABLE knjiga (
   id uuid not null PRIMARY KEY,
   naziv  varchar(255)
);


drop table if exists izvod ;
CREATE TABLE izvod (
   id uuid not null PRIMARY KEY,
   uporabnik_id uuid, 
   knjiga_id uuid, 
   signatura  varchar(255)
);
--
ALTER TABLE ONLY IZVOD
    ADD CONSTRAINT fk_izvod_uporabnik_id FOREIGN KEY (uporabnik_id) REFERENCES uporabnik(id);
ALTER TABLE ONLY IZVOD
    ADD CONSTRAINT fk_izvod_knjiga_id FOREIGN KEY (knjiga_id) REFERENCES knjiga(id);
--
create index "idx_uporabnik_id_on_izvod"
on izvod using btree (uporabnik_id);
create index "idx_knjiga_id_on_izvod"
on izvod using btree (knjiga_id);

create extension "pgcrypto";

-- insert into uporabnik  (id,ime,priimek) values (gen_random_uuid(),'Janko','Banko') ;
insert into uporabnik  (id,ime,priimek) values ('88144d9b-9db4-4d6a-89f7-713d27e2a2a2','Miha','Novak') ;
insert into uporabnik  (id,ime,priimek) values ('f95260c9-921e-46bd-80c6-3ded23437db0','Anica','Veber') ;
insert into uporabnik  (id,ime,priimek) values ('fa906169-e296-4a4d-aaa2-72a63e37278d','Janko','Bezjak') ;
-- select * from uporabnik;

insert into knjiga  (ID, Naziv) values ('af1aef60-468c-4ed5-afff-8e2060edeb1c','Vojna in mir') ;
insert into knjiga  (ID, Naziv) values ('8325b051-8f97-437c-b960-db5e5b3499cb','Metulj') ;
insert into knjiga  (ID, Naziv) values ('49c152fc-bea8-4eb4-b67d-26807ec79a22','Pod snegom') ;
-- select * from knjiga ;

insert into izvod  (ID, knjiga_id, uporabnik_id, Signatura) 
	values ('e2b3fe0d-3869-40ee-afa2-f5433be75f92', 'af1aef60-468c-4ed5-afff-8e2060edeb1c',null, 'Vojna in Mir 332') ;
insert into izvod (ID, knjiga_id, uporabnik_id, Signatura)  
	values ('40fa44cb-9d61-4b19-b301-7b99e9afe9dc', 'af1aef60-468c-4ed5-afff-8e2060edeb1c', '88144d9b-9db4-4d6a-89f7-713d27e2a2a2' , 'Vojna in Mir 654') ;
insert into izvod  (ID, knjiga_id, uporabnik_id, Signatura) 
	values ('d4e03cfe-5f9c-4668-90c4-8903002cbbca', 'af1aef60-468c-4ed5-afff-8e2060edeb1c','f95260c9-921e-46bd-80c6-3ded23437db0', 'Vojna in Mir 192') ;
insert into izvod  (ID, knjiga_id, uporabnik_id, Signatura) 
	values ('f7d85cc4-aa0c-4946-bb8c-8dd04136304e', '8325b051-8f97-437c-b960-db5e5b3499cb',null, 'Metulj 56') ;
insert into izvod  (ID, knjiga_id, uporabnik_id, Signatura) 
	values ('af3459ef-3c13-43a1-b897-a924ba9b217e', '8325b051-8f97-437c-b960-db5e5b3499cb','88144d9b-9db4-4d6a-89f7-713d27e2a2a2', 'Metulj 11') ;
insert into izvod (ID, knjiga_id, uporabnik_id, Signatura)  
	values ('3cf1d23e-faeb-4dd9-ba08-e69b6b395269', '49c152fc-bea8-4eb4-b67d-26807ec79a22',null, 'Pod snegom 6') ;
insert into izvod  (ID, knjiga_id, uporabnik_id, Signatura) 
	values ('8baddabb-2b4c-4720-9001-2a9ccce11792', '49c152fc-bea8-4eb4-b67d-26807ec79a22',null, 'Pod snegom 9') ;
insert into izvod (ID, knjiga_id, uporabnik_id, Signatura) 
	values ('85943d38-a68d-4575-8849-a2dbf59371d2', '49c152fc-bea8-4eb4-b67d-26807ec79a22',null, 'Pod snegom 11') ;
-- select * from izvod ;




