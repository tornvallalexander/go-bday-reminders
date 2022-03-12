CREATE TABLE "birthdays" (
     "id" bigserial PRIMARY KEY,
     "full_name" varchar NOT NULL,
     "future_birthday" timestamptz NOT NULL
);
