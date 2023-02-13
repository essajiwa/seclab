# seclab
API With Security Vulnerabilities

`CAUTIONS! Don't use this project as a base for real application`

## Pre-Requisites
* Create a database at Postgres Server
* Change the DB connection at `main.go:15` accordingly
* Run following script on those database

```sql
BEGIN;
-- CREATE TABLE "users" ----------------------------------------
CREATE TABLE "public"."users" ( 
	"id" Serial NOT NULL,
	"username" Character Varying NOT NULL,
	"password" Character Varying NOT NULL,
	"role" Character Varying NOT NULL,
	PRIMARY KEY ( "id" ),
	CONSTRAINT "unique_users_username" UNIQUE( "username" ) );
 ;
-- -------------------------------------------------------------
COMMIT;
```
```sql
BEGIN;

-- CREATE TABLE "products" -------------------------------------
CREATE TABLE "public"."products" ( 
	"id" Serial NOT NULL,
	"quantity" SmallInt NOT NULL,
	"name" Character Varying NOT NULL,
	"description" Character Varying,
	"category" Character Varying NOT NULL,
	"price" Double Precision DEFAULT 0 NOT NULL,
	PRIMARY KEY ( "id" ) );
 ;
-- -------------------------------------------------------------

-- CREATE INDEX "index_name" -----------------------------------
CREATE INDEX "index_name" ON "public"."products" USING btree( "name" Asc NULLS Last );
-- -------------------------------------------------------------

-- CREATE INDEX "index_category" -------------------------------
CREATE INDEX "index_category" ON "public"."products" USING btree( "category" Asc NULLS Last );
-- -------------------------------------------------------------

COMMIT;

```
