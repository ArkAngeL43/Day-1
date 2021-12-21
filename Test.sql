CREATE TABLE Test
(
    uid serial NOT NULL,
    Username character varying(100) NOT NULL,
    Inserted date,
    CONSTRAINT Users_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);
