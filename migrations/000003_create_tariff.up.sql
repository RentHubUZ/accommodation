CREATE TABLE IF NOT EXISTS "top_properties" (
  "id" UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  "property_id" UUID NOT NULL,
  "user_id" UUID NOT NULL,
  "start_date" DATE NOT NULL,
  "end_date" DATE NOT NULL,
  "tariff_name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  FOREIGN KEY ("property_id") REFERENCES "properties" ("id")
);

CREATE TABLE IF NOT EXISTS "tariffs" (
  "id" UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" VARCHAR(50),
  "days" INT,
  "price" FLOAT,
  "offers" TEXT,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);

CREATE TABLE IF NOT EXISTS "payments" (
  "id" UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  "amount" DECIMAL(10,2) NOT NULL,
  "status" VARCHAR(50) DEFAULT 'pending',
  "transaction_date" TIMESTAMP DEFAULT (NOW())
);
