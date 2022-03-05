CREATE TABLE "birthdays" (
     "id" bigserial PRIMARY KEY,
     "full_name" varchar UNIQUE NOT NULL,
     "future_birthday" timestamptz NOT NULL
);
