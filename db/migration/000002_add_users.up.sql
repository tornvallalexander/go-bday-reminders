DROP TABLE IF EXISTS "birthdays";

CREATE TABLE "users" (
     "username" varchar UNIQUE PRIMARY KEY,
     "hashed_password" varchar NOT NULL,
     "full_name" varchar NOT NULL,
     "email" varchar UNIQUE NOT NULL,
     "phone_number" varchar UNIQUE NOT NULL,
     "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
     "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reminders" (
     "id" bigserial PRIMARY KEY,
     "full_name" varchar NOT NULL,
     "personal_number" bigint NOT NULL,
     "user" varchar NOT NULL,
     "phone_number" varchar NOT NULL,
     "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "reminders" ADD FOREIGN KEY ("user") REFERENCES "users" ("username");