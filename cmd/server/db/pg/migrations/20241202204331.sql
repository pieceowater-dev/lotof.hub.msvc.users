-- Modify "friendships" table
ALTER TABLE "public"."friendships" DROP CONSTRAINT "fk_friendships_user", ALTER COLUMN "status" TYPE smallint, ADD
 CONSTRAINT "fk_friendships_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD
 CONSTRAINT "fk_friendships_friend" FOREIGN KEY ("friend_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
