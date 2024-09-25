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
  "description" text,
  "roommate_count" int,
  "lease_terms" text,
  "lease_duration" int,
  "top_status" bool DEFAULT false,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "updated_at" TIMESTAMP DEFAULT (NOW())
);