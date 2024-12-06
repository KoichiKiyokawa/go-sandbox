-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "email", ADD COLUMN "nickname" character varying(255) NULL;
