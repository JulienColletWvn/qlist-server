CREATE TYPE "pricing_source_type" AS ENUM (
  'web',
  'onsite'
);

CREATE TYPE "transaction_status" AS ENUM (
  'pending',
  'validated',
  'cancelled'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "username" varchar(30),
  "email" varchar(30) NOT NULL,
  "password" varchar(128) NOT NULL,
  "firstname" varchar(30) NOT NULL,
  "lastname" varchar(30) NOT NULL,
  "phone" varchar(30)
);

CREATE TABLE "guests" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "creator_id" int NOT NULL,
  "email" varchar(30) NOT NULL,
  "firstname" varchar(30),
  "lastname" varchar(30),
  "phone" varchar(30)
);

CREATE TABLE "guests_groups" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "guests_id" int NOT NULL,
  "guests_groups_types_id" int NOT NULL
);

CREATE TABLE "guests_groups_types" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "creator_id" int NOT NULL,
  "group_name" varchar(30) NOT NULL,
  "group_color" varchar(7) DEFAULT NULL
);

CREATE TABLE "cashiers" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "users_id" int NOT NULL,
  "events_id" int NOT NULL
);

CREATE TABLE "sellers" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "users_id" int NOT NULL,
  "events_id" int NOT NULL
);

CREATE TABLE "events_administrators" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "users_id" int NOT NULL,
  "events_id" int NOT NULL
);

CREATE TABLE "events" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "name" varchar(50) NOT NULL,
  "description" varchar(200) NOT NULL,
  "start_date" TIMESTAMPTZ NOT NULL,
  "end_date" TIMESTAMPTZ NOT NULL,
  "location" varchar(50) NOT NULL,
  "free_wifi" boolean NOT NULL,
  "public" boolean NOT NULL,
  "tickets_amount" int NOT NULL
);

CREATE TABLE "events_guests" (
  "id" SERIAL PRIMARY KEY,
  "guests_id" int NOT NULL,
  "events_id" int NOT NULL
);

CREATE TABLE "events_guests_groups" (
  "id" SERIAL PRIMARY KEY,
  "guests_groups_types" int NOT NULL,
  "events_id" int NOT NULL
);

CREATE TABLE "events_photos" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "events_id" int NOT NULL,
  "url" varchar(50) NOT NULL
);

CREATE TABLE "tokens_transactions" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "transaction_date" TIMESTAMPTZ NOT NULL,
  "amount" int NOT NULL,
  "online_sell" boolean NOT NULL,
  "cashiers_id" int NOT NULL,
  "sellers_id" int NOT NULL,
  "tokens_id" int NOT NULL,
  "events_products_id" int NOT NULL,
  "status" transaction_status NOT NULL
);

CREATE TABLE "tokens" (
  "id" SERIAL PRIMARY KEY,
  "uuid" varchar(50) NOT NULL,
  "wallets_id" int NOT NULL
);

CREATE TABLE "wallets" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "guests_id" int NOT NULL,
  "wallets_type_id" int NOT NULL
);

CREATE TABLE "wallets_type" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "events_id" int NOT NULL,
  "name" varchar(50) NOT NULL,
  "start_validity_date" TIMESTAMPTZ NOT NULL,
  "end_validity_date" TIMESTAMPTZ NOT NULL,
  "max_amount" int NOT NULL,
  "online_reload" boolean NOT NULL
);

CREATE TABLE "wallets_pricing" (
  "id" SERIAL PRIMARY KEY,
  "type" pricing_source_type NOT NULL,
  "quantity" int NOT NULL,
  "unit_price" int NOT NULL,
  "wallets_type_id" int NOT NULL
);

CREATE TABLE "wallets_transactions" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "cashiers_id" int NOT NULL,
  "wallets_id" int NOT NULL,
  "wallets_pricing_id" int NOT NULL,
  "units_sold" int NOT NULL,
  "status" transaction_status NOT NULL
);

CREATE TABLE "tickets" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "tickets_type_id" int NOT NULL,
  "sellers_id" int NOT NULL,
  "guests_id" int NOT NULL
);

CREATE TABLE "tickets_type" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "events_id" int NOT NULL,
  "name" varchar(50) NOT NULL,
  "start_validity_date" TIMESTAMPTZ,
  "end_validity_date" TIMESTAMPTZ,
  "usage_limitation" int NOT NULL,
  "usage_unlimited" boolean NOT NULL
);

CREATE TABLE "tickets_transactions" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "tickets_id" int NOT NULL,
  "amount" int NOT NULL,
  "status" transaction_status NOT NULL
);

CREATE TABLE "events_products" (
  "id" SERIAL PRIMARY KEY,
  "created_at" TIMESTAMPTZ DEFAULT (now()),
  "updated_at" TIMESTAMPTZ DEFAULT (now()),
  "events_id" int NOT NULL,
  "sellers_id" int NOT NULL,
  "name" varchar(30) NOT NULL,
  "tokens_amount_pricing" int NOT NULL
);

ALTER TABLE "guests" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "guests_groups" ADD FOREIGN KEY ("guests_id") REFERENCES "guests" ("id");

ALTER TABLE "guests_groups" ADD FOREIGN KEY ("guests_groups_types_id") REFERENCES "guests_groups_types" ("id");

ALTER TABLE "guests_groups_types" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "cashiers" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "cashiers" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "sellers" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "sellers" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "events_administrators" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "events_administrators" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "events_guests" ADD FOREIGN KEY ("guests_id") REFERENCES "guests" ("id");

ALTER TABLE "events_guests" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "events_guests_groups" ADD FOREIGN KEY ("guests_groups_types") REFERENCES "guests_groups_types" ("id");

ALTER TABLE "events_guests_groups" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "events_photos" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "tokens_transactions" ADD FOREIGN KEY ("cashiers_id") REFERENCES "cashiers" ("id");

ALTER TABLE "tokens_transactions" ADD FOREIGN KEY ("sellers_id") REFERENCES "sellers" ("id");

ALTER TABLE "tokens_transactions" ADD FOREIGN KEY ("tokens_id") REFERENCES "tokens" ("id");

ALTER TABLE "tokens_transactions" ADD FOREIGN KEY ("events_products_id") REFERENCES "events_products" ("id");

ALTER TABLE "tokens" ADD FOREIGN KEY ("wallets_id") REFERENCES "wallets" ("id");

ALTER TABLE "wallets" ADD FOREIGN KEY ("guests_id") REFERENCES "guests" ("id");

ALTER TABLE "wallets" ADD FOREIGN KEY ("wallets_type_id") REFERENCES "wallets_type" ("id");

ALTER TABLE "wallets_type" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "wallets_pricing" ADD FOREIGN KEY ("wallets_type_id") REFERENCES "wallets_type" ("id");

ALTER TABLE "wallets_transactions" ADD FOREIGN KEY ("cashiers_id") REFERENCES "cashiers" ("id");

ALTER TABLE "wallets_transactions" ADD FOREIGN KEY ("wallets_id") REFERENCES "wallets" ("id");

ALTER TABLE "wallets_transactions" ADD FOREIGN KEY ("wallets_pricing_id") REFERENCES "wallets_pricing" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("tickets_type_id") REFERENCES "tickets_type" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("sellers_id") REFERENCES "sellers" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("guests_id") REFERENCES "guests" ("id");

ALTER TABLE "tickets_type" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "tickets_transactions" ADD FOREIGN KEY ("tickets_id") REFERENCES "tickets" ("id");

ALTER TABLE "events_products" ADD FOREIGN KEY ("events_id") REFERENCES "events" ("id");

ALTER TABLE "events_products" ADD FOREIGN KEY ("sellers_id") REFERENCES "sellers" ("id");
