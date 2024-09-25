CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE "properties" (
  "id" UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  "owner_id" UUID NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  "price" DECIMAL(10,2) NOT NULL,
  "property_type" VARCHAR(50),
  "bedrooms" INT,
  "bathrooms" INT,
  "square_footage" DECIMAL(10,2),
  "listing_status" VARCHAR(20),
  "description" TEXT,
  "roommate_count" INT,
  "lease_terms" TEXT,
  "lease_duration" INT,
  "top_status" BOOL DEFAULT FALSE,
  "location" GEOGRAPHY(Point, 4326),  -- Yangi field qo'shildi
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW()),
  "deleted_at" TIMESTAMP
);
