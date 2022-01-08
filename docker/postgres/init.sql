DROP TABLE IF EXISTS "cryptocurrencies";
DROP SEQUENCE IF EXISTS cryptocurrencies_id_seq;
CREATE SEQUENCE cryptocurrencies_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."cryptocurrencies" (
    "id" integer DEFAULT nextval('cryptocurrencies_id_seq') NOT NULL,
    "name" varchar(100) NOT NULL UNIQUE,
    "token" varchar(10) NOT NULL UNIQUE,
    "votes" numeric DEFAULT '0' NOT NULL,
    CONSTRAINT "cryptocurrencies_pkey" PRIMARY KEY ("id")
) WITH (oids = false);