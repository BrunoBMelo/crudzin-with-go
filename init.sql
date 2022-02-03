CREATE DATABASE PG_DBGO;

\c PG_DBGO;

CREATE TABLE public.customer (
	id serial primary KEY,
	"name" varchar NOT NULL,
	birthdate date NOT NULL,
	phonenumber varchar NOT NULL,
	email varchar NOT NULL,
	zipcode varchar NOT NULL,
	city varchar NOT NULL,
	street varchar NOT NULL,
	country varchar NOT NULL
);
