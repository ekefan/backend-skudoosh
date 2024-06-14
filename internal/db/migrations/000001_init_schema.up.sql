CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "fullname" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "trip_state" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "emergency_contacts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "owner" bigint NOT NULL,
  "email" varchar NOT NULL,
  "phone_number" varchar NOT NULL
);

CREATE TABLE "itineraries" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "owner" bigint NOT NULL,
  "take_off_date" timestamptz NOT NULL,
  "return_date" timestamptz NOT NULL,
  "destination" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "activity_lists" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "trip_owner" bigint NOT NULL,
  "activity" varchar NOT NULL,
  "date_time" timestamptz NOT NULL,
  "checked" boolean NOT NULL DEFAULT false
);

CREATE TABLE "travel_checklists" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "trip_owner" bigint NOT NULL,
  "item_task" varchar NOT NULL,
  "checked" boolean NOT NULL
);

CREATE TABLE "trip_bookings" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "trip_owner" bigint NOT NULL,
  "booking_type" varchar NOT NULL,
  "booking_details" varchar NOT NULL
);

CREATE TABLE "trip_logs" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "trip_owner" bigint NOT NULL,
  "logs" jsonb NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "emergency_contacts" ("owner");

CREATE INDEX ON "itineraries" ("owner");

CREATE INDEX ON "activity_lists" ("trip_owner");

CREATE INDEX ON "travel_checklists" ("trip_owner");

CREATE INDEX ON "trip_bookings" ("trip_owner");

CREATE INDEX ON "trip_logs" ("trip_owner");

ALTER TABLE "emergency_contacts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "itineraries" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "activity_lists" ADD FOREIGN KEY ("trip_owner") REFERENCES "itineraries" ("id");

ALTER TABLE "travel_checklists" ADD FOREIGN KEY ("trip_owner") REFERENCES "itineraries" ("id");

ALTER TABLE "trip_bookings" ADD FOREIGN KEY ("trip_owner") REFERENCES "itineraries" ("id");

ALTER TABLE "trip_logs" ADD FOREIGN KEY ("trip_owner") REFERENCES "itineraries" ("id");
