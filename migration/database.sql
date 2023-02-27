CREATE TABLE IF NOT EXISTS HOUSE
(
	id_cadastral TEXT NOT NULL COLLATE NOCASE PRIMARY KEY,
	address TEXT NOT NULL COLLATE NOCASE,
	id_owner INTEGER NULL,
	id_municipality INTEGER NULL,
	CONSTRAINT FK_HOUSE_MUNICIPALITY FOREIGN KEY (id_municipality) REFERENCES MUNICIPALITY (id_municipality) ON DELETE No Action ON UPDATE No Action,
	CONSTRAINT FK_HOUSE_PERSON FOREIGN KEY (id_owner) REFERENCES PERSON (id_number) ON DELETE No Action ON UPDATE No Action
)
;

CREATE TABLE IF NOT EXISTS MUNICIPALITY
(
	id_municipality INTEGER NOT NULL PRIMARY KEY,
	municipality_name TEXT NOT NULL COLLATE NOCASE
)
;

CREATE TABLE IF NOT EXISTS PERSON
(
	id_number INTEGER NOT NULL PRIMARY KEY,
	forename TEXT NOT NULL COLLATE NOCASE,
	surname TEXT NOT NULL COLLATE NOCASE,
	birth_date TEXT NOT NULL COLLATE NOCASE,
	sex TEXT NOT NULL COLLATE NOCASE,
	id_home TEXT NULL COLLATE NOCASE,
	CONSTRAINT FK_PERSON_HOUSE FOREIGN KEY (id_home) REFERENCES HOUSE (id_cadastral) ON DELETE No Action ON UPDATE No Action,
	CONSTRAINT CS_sex CHECK (sexo in ('M','F','N'))
)
;

CREATE TABLE IF NOT EXISTS RESPONSIBLE_PERSON
(
	id_responsible INTEGER NOT NULL,
	id_person INTEGER NOT NULL,
	CONSTRAINT PK_RESPONSIBLE PRIMARY KEY (id_responsible,id_person),
	CONSTRAINT FK_RESPONSIBLE_PERSON_PERSON FOREIGN KEY (id_responsible) REFERENCES PERSON (id_number) ON DELETE No Action ON UPDATE No Action,
	CONSTRAINT FK_RESPONSIBLE_PERSON_PERSON_02 FOREIGN KEY (id_person) REFERENCES PERSON (id_number) ON DELETE No Action ON UPDATE No Action
)
;


CREATE INDEX IXFK_HOUSE_MUNICIPALITY
 ON HOUSE (id_municipality ASC)
;

CREATE INDEX IXFK_HOUSE_PERSON
 ON HOUSE (id_owner ASC)
;

CREATE INDEX IXFK_PERSON_HOUSE
 ON PERSON (id_home ASC)
;

CREATE INDEX IXFK_RESPONSIBLE_PERSON_PERSON
 ON RESPONSIBLE_PERSON (id_responsible ASC)
;

CREATE INDEX IXFK_RESPONSIBLE_PERSON_PERSON_02
 ON RESPONSIBLE_PERSON (id_person ASC)
;