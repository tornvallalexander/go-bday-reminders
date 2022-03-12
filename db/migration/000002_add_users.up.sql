DROP TABLE IF EXISTS "birthdays";

CREATE TABLE "users" (
     "user_name" varchar UNIQUE PRIMARY KEY,
     "hashed_password" varchar NOT NULL,
     "full_name" varchar NOT NULL,
     "email" varchar UNIQUE NOT NULL,
     "phone_number" bigint UNIQUE NOT NULL,
     "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
     "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reminders" (
     "id" bigserial PRIMARY KEY,
     "full_name" varchar NOT NULL,
     "personal_number" bigint NOT NULL,
     "user" varchar NOT NULL,
     "phone_number" bigint NOT NULL,
     "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "reminders" ADD FOREIGN KEY ("user") REFERENCES "users" ("user_name");

CREATE INDEX ON "reminders" ("user");

CREATE UNIQUE INDEX ON "reminders" ("user");