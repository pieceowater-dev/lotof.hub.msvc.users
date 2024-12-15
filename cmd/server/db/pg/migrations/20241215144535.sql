-- Modify "users" table
ALTER TABLE "public"."users" ADD CONSTRAINT "uni_users_username" UNIQUE ("username");
