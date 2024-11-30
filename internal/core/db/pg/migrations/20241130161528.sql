-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "username" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  "state" bigint NULL DEFAULT 100,
  "deleted" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email")
);
-- Create "friendships" table
CREATE TABLE "public"."friendships" (
  "user_id" uuid NOT NULL,
  "friend_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "friend_id"),
  CONSTRAINT "fk_friendships_friends" FOREIGN KEY ("friend_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_friendships_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
