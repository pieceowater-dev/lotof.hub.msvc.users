-- Modify "friendships" table
ALTER TABLE "public"."friendships" DROP CONSTRAINT "friendships_pkey", ADD COLUMN "id" uuid NOT NULL, ADD COLUMN "created_at" timestamptz NULL, ADD PRIMARY KEY ("id");
