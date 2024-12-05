-- Create "users" table
CREATE TABLE "public"."users" ("id" character(26) NOT NULL, "name" character varying(255) NOT NULL, "email" character varying(255) NOT NULL, "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL, PRIMARY KEY ("id"));
-- Create "user_secrets" table
CREATE TABLE "public"."user_secrets" ("user_id" character(26) NOT NULL, "password_hash" character varying(255) NOT NULL, CONSTRAINT "user_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
