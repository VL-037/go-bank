CREATE TABLE "users"
(
    "username"            varchar PRIMARY KEY,
    "hashed_password"     varchar     NOT NULL,
    "full_name"           varchar     NOT NULL,
    "email"               varchar UNIQUE,
    "password_updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_by"          varchar,
    "created_at"          timestamptz NOT NULL DEFAULT (now()),
    "updated_by"          varchar,
    "updated_at"          timestamptz NOT NULL DEFAULT (now()),
    "mark_for_delete"     boolean     NOT NULL DEFAULT false
);

ALTER TABLE "accounts"
    ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts"
    ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency")