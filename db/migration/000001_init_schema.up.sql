CREATE TABLE "base_entity" (
  "id" bigserial PRIMARY KEY,
  "created_by" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_by" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "mark_for_delete" boolean NOT NULL DEFAULT false
);

CREATE TABLE "accounts" (
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL
)INHERITS ("base_entity");

CREATE TABLE "entries" (
  "account_id" bigint,
  "amount" bigint NOT NULL
)INHERITS ("base_entity");

CREATE TABLE "transfers" (
  "from_account_id" bigint,
  "to_account_id" bigint,
  "amount" bigint NOT NULL
)INHERITS ("base_entity");

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'mmust be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "base_entity" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "base_entity" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "base_entity" ("id");
