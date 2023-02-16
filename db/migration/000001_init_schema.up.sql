CREATE TABLE "accounts"
(
    "id"              bigserial PRIMARY KEY,
    "owner"           varchar     NOT NULL,
    "balance"         bigint      NOT NULL,
    "currency"        varchar     NOT NULL,
    "created_by"      varchar,
    "created_at"      timestamptz NOT NULL DEFAULT (now()),
    "updated_by"      varchar,
    "updated_at"      timestamptz NOT NULL DEFAULT (now()),
    "mark_for_delete" boolean     NOT NULL DEFAULT false
);

CREATE TABLE "entries"
(
    "id"              bigserial PRIMARY KEY,
    "account_id"      bigint      NOT NULL,
    "amount"          bigint      NOT NULL,
    "created_by"      varchar,
    "created_at"      timestamptz NOT NULL DEFAULT (now()),
    "updated_by"      varchar,
    "updated_at"      timestamptz NOT NULL DEFAULT (now()),
    "mark_for_delete" boolean     NOT NULL DEFAULT false
);

CREATE TABLE "transfers"
(
    "id"              bigserial PRIMARY KEY,
    "from_account_id" bigint      NOT NULL,
    "to_account_id"   bigint      NOT NULL,
    "amount"          bigint      NOT NULL,
    "created_by"      varchar,
    "created_at"      timestamptz NOT NULL DEFAULT (now()),
    "updated_by"      varchar,
    "updated_at"      timestamptz NOT NULL DEFAULT (now()),
    "mark_for_delete" boolean     NOT NULL DEFAULT false
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT
ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT
ON COLUMN "transfers"."amount" IS 'mmust be positive';

ALTER TABLE "entries"
    ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
