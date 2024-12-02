-- Modify "friendships" table
ALTER TABLE "public"."friendships" ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "deleted_at" timestamptz NULL, ADD COLUMN "status" bigint NULL DEFAULT 100;
-- Create index "idx_friendships_deleted_at" to table: "friendships"
CREATE INDEX "idx_friendships_deleted_at" ON "public"."friendships" ("deleted_at");
