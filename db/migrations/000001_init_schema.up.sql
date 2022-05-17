CREATE TABLE "statuses" (
  "code" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "members" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT (now()),
  "created_at" timestamp DEFAULT (now()),
  "status_code" varchar NOT NULL
);

CREATE TABLE "deposits" (
  "id" SERIAL PRIMARY KEY,
  "reference_num" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "currency" varchar DEFAULT 'KES',
  "member_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "status_code" varchar NOT NULL
);

CREATE TABLE "contributions" (
  "id" SERIAL PRIMARY KEY,
  "deposit_id" int,
  "amount" bigint NOT NULL,
  "currency" varchar DEFAULT 'KES',
  "member_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "status_code" varchar NOT NULL
);

CREATE INDEX ON "deposits" ("reference_num");

CREATE INDEX ON "deposits" ("member_id");

CREATE INDEX ON "contributions" ("deposit_id");

CREATE INDEX ON "contributions" ("member_id");

ALTER TABLE "members" ADD FOREIGN KEY ("status_code") REFERENCES "statuses" ("code");

ALTER TABLE "deposits" ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id");

ALTER TABLE "deposits" ADD FOREIGN KEY ("status_code") REFERENCES "statuses" ("code");

ALTER TABLE "contributions" ADD FOREIGN KEY ("deposit_id") REFERENCES "deposits" ("id");

ALTER TABLE "contributions" ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id");

ALTER TABLE "contributions" ADD FOREIGN KEY ("status_code") REFERENCES "statuses" ("code");